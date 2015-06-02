package eventrecv

import (
    "net/http"

    "github.com/krafthack/ifooz-engine/rest"
)

func Init(mux *http.ServeMux) {
    newmatchhandler := setupNewMatchHandler()
    mux.HandleFunc("/newmatch", rest.PostOnly(newmatchhandler))
}

type eventReciever struct {
    mux *http.ServeMux
}
