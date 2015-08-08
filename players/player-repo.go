package players

import (
    "github.com/docker/docker/pkg/namesgenerator"
    "code.google.com/p/go-uuid/uuid"
)

func GetPlayer(playerId uuid.UUID) *Player {
    name := namesgenerator.GetRandomName(0)
    return &Player{ playerId, name }
}
