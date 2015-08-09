package events

import (
    "time"
    "errors"
    "fmt"
)

var teams = []string{ "blue", "white" }

type Goal struct {
    Team string
    MatchGoalNum int
    Score int
    Timestamp time.Time
}

func NewGoal(team string, matchGoalNum int, score int, timestamp time.Time) (Goal, error) {

    if !validTeam(team) {
        return Goal{}, errors.New(fmt.Sprintf("%s is not a valid team", team))
    }

    return Goal{ team, matchGoalNum, score, timestamp }, nil
}

func validTeam(team string) bool {
    for _, t := range teams {
        if t == team {
            return true
        }
    }
    return false
}
