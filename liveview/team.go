package liveview

import (
    "github.com/krafthack/ifooz-engine/players"
)

type Team struct {
    Players []*players.Player
    Goals []*Goal
}

func CreateTeam(defFirst *players.Player, offFirst *players.Player) *Team {

    players := []*players.Player{ defFirst, offFirst }
    goals := []*Goal{}
    team := &Team{ players, goals }

    return team
}
