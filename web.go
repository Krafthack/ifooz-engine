package main

import (
    "os"
    "net/http"

    "github.com/krafthack/ifooz-engine/eventrecv"
    "github.com/krafthack/ifooz-engine/servicebase"
)

var (
    port = "PORT"
    defaultport = "8066"
)

func main() {

    redisIp := os.Getenv("REDIS_IP")
    redisPort := os.Getenv("REDIS_PORT")
    redisUrl := redisIp + ":" + redisPort
    if redisUrl == ":" {
        redisUrl = "localhost:6379"
    }

    service, mux := servicebase.NewServer()

    eventrecv.Init(mux, redisUrl)

    port := os.Getenv(port)
    if port == "" {
        port = defaultport
    }

    http.ListenAndServe(":"+ port, service)
}
