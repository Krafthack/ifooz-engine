package engine

import (
    "testing"
    "time"

    "code.google.com/p/go-uuid/uuid"

    "github.com/krafthack/ifooz-engine/events"
)

func TestBlueGoalOneNil(t *testing.T) {
    tablename := "skuret"
    engine := matchWithEngine(tablename)

    bluegoalevent, err := events.NewGoal("blue", 1, 1, time.Now())
    if err != nil {
        t.Errorf("did not expect error when creating goal")
    }
    engine.GoalEvent(tablename, bluegoalevent)
    updates := engine.GetTableUpdateChannel(tablename)

    <- updates
    blueGoalUpdate := <- updates

    if blueGoalUpdate.Blue.Score != 1 {
        t.Errorf("Blue goal opens match. Expect blue score to be 1 got %d", blueGoalUpdate.Blue.Score)
    }

    if blueGoalUpdate.White.Score != 0 {
        t.Errorf("Blue goal opens match. Expect white score to be 0 got %d", blueGoalUpdate.White.Score)
    }
}

func TestBlueDonut(t *testing.T) {
    tablename := "skuret"
    engine := matchWithEngine(tablename)


    updates := engine.GetTableUpdateChannel(tablename)
    <- updates

    for i := 0; i < 10; i++ {
        bluegoalevent, err := events.NewGoal("blue", 1+i, 1+i, time.Now())
        if err != nil {
            t.Errorf("did not expect error when creating goal")
        }
        engine.GoalEvent(tablename, bluegoalevent)

        blueGoalUpdate := <- updates

        if blueGoalUpdate.Blue.Score != i+1 {
            t.Errorf("Blue goal. Expect blue score to be %d got %d", i+1, blueGoalUpdate.Blue.Score)
        }

        if blueGoalUpdate.White.Score != 0 {
            t.Errorf("Blue goal. Expect white score to be 0 got %d", blueGoalUpdate.White.Score)
        }
    }
}

func matchWithEngine(tablename string) *Engine {
    newmatchevent := &events.NewMatchEvent{ uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(), time.Now() }
    engine := CreateEngine()
    engine.AddTable(tablename)
    engine.NewMatchEvent(tablename, newmatchevent)

    return engine
}
