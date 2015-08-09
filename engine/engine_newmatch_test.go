package engine

import (
    "testing"
    "time"

    "code.google.com/p/go-uuid/uuid"

    "github.com/krafthack/ifooz-engine/events"
)

func TestNewMatchExpectTableUpdate(t *testing.T) {

    tablename := "skuret"
    newmatchevent := &events.NewMatchEvent{ uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(),  time.Now() }
    engine := CreateEngine()

    engine.AddTable(tablename)
    tableUpdates := engine.GetTableUpdateChannel(tablename)

    engine.NewMatchEvent(tablename, newmatchevent)

    update := <- tableUpdates

    if update == nil {
        t.Error("Expected tableupdate to be a non nil value.")
    }
}

func TestNewMatchExpectPlayersSet(t *testing.T) {
    tablename := "skuret"
    newmatchevent := &events.NewMatchEvent{ uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(),  time.Now() }
    engine := CreateEngine()

    engine.AddTable(tablename)
    tableUpdates := engine.GetTableUpdateChannel(tablename)

    engine.NewMatchEvent(tablename, newmatchevent)

    update := <- tableUpdates

    bluePlayers := update.Blue.Players
    whitePlayers := update.White.Players

    if len(bluePlayers) != 2 {
        t.Error("Expected blue team to have two players. Got:", len(update.Blue.Players))
    }

    if bluePlayers[1].Id.String() != newmatchevent.BOffFirst.String() {
        t.Error("Expected Blue second player to be BOffFirst from  newmatch event.")
    }

    if len(whitePlayers) != 2 {
        t.Error("Expected white team to have two players. Got:", len(update.White.Players))
    }

    if whitePlayers[0].Id.String() != newmatchevent.WDefFirst.String() {
        t.Error("Expected White first player to be WDefFirst from newmatch event.")
    }
}

func TestNewMatchExpectNoGoals(t *testing.T) {
    tablename := "skuret"
    newmatchevent := &events.NewMatchEvent{ uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(), uuid.NewRandom(),  time.Now() }
    engine := CreateEngine()

    engine.AddTable(tablename)
    tableUpdates := engine.GetTableUpdateChannel(tablename)

    engine.NewMatchEvent(tablename, newmatchevent)

    update := <- tableUpdates

    if len(update.Blue.Goals) > 0 {
        t.Errorf("Expected blue to have 0 goals got %d", len(update.Blue.Goals))
    }

    if len(update.White.Goals) > 0 {
        t.Errorf("Expected white to have 0 goals got %d", len(update.White.Goals))
    }
}
