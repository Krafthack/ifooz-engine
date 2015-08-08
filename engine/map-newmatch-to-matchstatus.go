package engine

import (
    "github.com/krafthack/ifooz-engine/liveview"
    "github.com/krafthack/ifooz-engine/events"
    "github.com/krafthack/ifooz-engine/players"
)

func MapNewMatchToMatchStatus(newmatch *events.NewMatchEvent) (*liveview.MatchStatus, error) {

    blueDefFirst := players.GetPlayer(newmatch.BDefFirst)
    blueOffFirst := players.GetPlayer(newmatch.BOffFirst)
    blue := liveview.CreateTeam(blueDefFirst, blueOffFirst)

    whiteDefFirst := players.GetPlayer(newmatch.WDefFirst)
    whiteOffFirst := players.GetPlayer(newmatch.WOffFirst)
    white := liveview.CreateTeam(whiteDefFirst, whiteOffFirst)

    matchstatus := &liveview.MatchStatus{
        blue,
        white,
        "no-yolo",
    }

    return matchstatus, nil
}
