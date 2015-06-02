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
    service, mux := servicebase.NewServer()

    eventrecv.Init(mux)

    port := os.Getenv(port)
    if port == "" {
        port = defaultport
    }

    http.ListenAndServe(":"+ port, service)
}
