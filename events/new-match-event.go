package events

import (
    "code.google.com/p/go-uuid/uuid"
    "time"
)

type NewMatchEvent struct {
    WDefFirst uuid.UUID `json:"w_def_first"`
    WOffFirst uuid.UUID `json:"w_off_first"`
    BDefFirst uuid.UUID `json:"b_def_first"`
    BOffFirst uuid.UUID `json:"b_off_first"`

    Timestamp time.Time `json:"timestamp"`
}
