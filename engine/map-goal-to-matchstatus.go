package engine

import (
    "github.com/krafthack/ifooz-engine/liveview"
    "github.com/krafthack/ifooz-engine/events"
)

func MapGoalToMatchStatus(lastStatus *liveview.MatchStatus, goal events.Goal) (*liveview.MatchStatus, error) {

    matchstatus := &liveview.MatchStatus{}
    *matchstatus = *lastStatus

    if goal.Team == "blue" && len(matchstatus.Blue.Goals) == goal.Score - 1 {
        matchstatus.Blue.Goals = append(matchstatus.Blue.Goals, liveview.Goal{"hei"})
        matchstatus.Blue.Score = goal.Score
    } else if goal.Team == "white" && len(matchstatus.White.Goals) == goal.Score - 1 {
        matchstatus.White.Goals = append(matchstatus.White.Goals, liveview.Goal{"Hei"})
        matchstatus.White.Score = goal.Score
    }

    return matchstatus, nil
}
