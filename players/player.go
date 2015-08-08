package players

import (
    "code.google.com/p/go-uuid/uuid"
)

type Player struct {
    Id uuid.UUID
    Name string
}
