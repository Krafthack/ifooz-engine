package eventrecv

import (
    "net/http"
    "strconv"
    "encoding/json"


    "github.com/krafthack/ifooz-engine/rest"
    "github.com/krafthack/ifooz-engine/redisclient"
    "github.com/krafthack/ifooz-engine/newmatch"
    "github.com/krafthack/ifooz-engine/servicebase"

    "github.com/fzzy/radix/redis"
)

var (
    commitskey = "newmatch"
    streamkey = "matches"
)

func setupNewMatchHandler(redis string) rest.Handler {
    loggers := servicebase.InitLoggers("Register New Match")
    client, err := redisclient.Connection(redis)
    if err != nil {
        loggers.E.Println(err)
        panic(err)
    }
    newmatch := &newMatch{client, loggers}
    return newmatch.emitNewMatchEvent
}

type newMatch struct {
    client *redis.Client
    loggers *servicebase.Loggers
}

func (nm *newMatch) emitNewMatchEvent(w http.ResponseWriter, r *http.Request) {
    event, err := newmatch.FromBody(r.Body)
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 400, err)
        return
    }

    storeRevision, err := nm.client.Cmd("hincrby", "revisions", commitskey, 1).Int()
    if err != nil {
        nm.loggers.E.Println("Could not get store revision.", err)
        rest.Response(w, 500, err)
        return
    }

    err = startTransaction(nm.client)
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 500, err)
        return
    }

    commitId := commitskey + strconv.Itoa(storeRevision + 1)
    commitData, err := json.Marshal(event)
    if err != nil {
        nm.loggers.E.Println("Could not marshal event.", err)
        rest.Response(w, 500, err)
        return
    }

    err = nm.client.Cmd("hset", commitskey, commitId, commitData).Err
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 500, err)
        return
    }

    err = nm.client.Cmd("rpush", streamkey, commitId).Err
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 500, err)
        return
    }

    err = nm.client.Cmd("publish", commitskey, commitData).Err
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 500, err)
        return
    }

    err = nm.client.Cmd("EXEC").Err
    if err != nil {
        nm.loggers.E.Println(err)
        rest.Response(w, 500, err)
        return
    }

    rest.Response(w, 200, commitId)
}

func startTransaction(client *redis.Client) error {
    return client.Cmd("MULTI").Err
}
