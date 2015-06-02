package eventrecv

import (
    "net/http"

    "github.com/krafthack/ifooz-engine/rest"
)

func Init(mux *http.ServeMux, redis string) {
    newmatchhandler := setupNewMatchHandler(redis)
    mux.HandleFunc("/newmatch", rest.PostOnly(newmatchhandler))
}

type eventReciever struct {
    mux *http.ServeMux
}
