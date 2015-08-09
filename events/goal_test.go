package events

import (
    "testing"
    "time"
)

func TestNewGoalBlueTeamExpectBlueGoal(t *testing.T) {

    goal, err := NewGoal("blue", 1, 1, time.Now())

    if err != nil {
        t.Errorf("did not expect error creating goal")
    }

    if goal.Team != "blue" {
        t.Errorf("NewGoal('blue') = %s. Want blue", goal.Team)
    }
}

func TestNewGoalRedTeamExpectError(t *testing.T) {
    _, err := NewGoal("red", 1, 1, time.Now())

    if err == nil {
        t.Errorf("expected red goal to cause error")
    }
}
