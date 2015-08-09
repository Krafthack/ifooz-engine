package liveview

import (
    "github.com/krafthack/ifooz-engine/players"
)

type Team struct {
    Players []*players.Player
    Goals []Goal
    Score int
}

func CreateTeam(defFirst *players.Player, offFirst *players.Player) *Team {

    players := []*players.Player{ defFirst, offFirst }
    goals := []Goal{}
    score := 0
    team := &Team{ players, goals, score }

    return team
}
