package eventrecv

import (
    "net/http"

    "github.com/krafthack/ifooz-engine/rest"
)

func setupNewMatchHandler() rest.Handler {
    newmatch := &newMatch{}
    return newmatch.handler
}

type newMatch struct {

}

func (nm *newMatch) handler(w http.ResponseWriter, r *http.Request) {
    rest.Response(w, 200, "hello")
}
