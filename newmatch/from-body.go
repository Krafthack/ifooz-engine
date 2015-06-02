package newmatch

import (
    "io"
    "encoding/json"
    "code.google.com/p/go-uuid/uuid"
)

type NewMatchEvent struct {
    WDefFirst uuid.UUID `json:"w_def_first"`
    WOffFirst uuid.UUID `json:"w_off_first"`
    BDefFirst uuid.UUID `json:"b_def_first"`
    BOffFirst uuid.UUID `json:"b_off_first"`

    Timestamp int64 `json:"timestamp"`
}

func FromBody(body io.ReadCloser) (NewMatchEvent, error) {
    var newMatchEvent NewMatchEvent
    decoder := json.NewDecoder(body)
    err := decoder.Decode(&newMatchEvent)

    return newMatchEvent, err
}
