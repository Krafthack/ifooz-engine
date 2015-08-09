package engine

import (
    "errors"
    "github.com/krafthack/ifooz-engine/liveview"
    "github.com/krafthack/ifooz-engine/events"
)

func CreateEngine() *Engine {

    tables := make(map[string] *Table)
    engine := &Engine{ tables }

    return engine
}

type Engine struct {
    tables map[string] *Table
}

func (e *Engine) AddTable(tablename string) {
    table := CreateTable(tablename)
    e.tables[tablename] = table
}

func (e *Engine) GetTableUpdateChannel(tablename string) chan *liveview.MatchStatus {
    return e.tables[tablename].GetUpdatesChannel()
}

func (e *Engine) NewMatchEvent(tablename string, newmatch *events.NewMatchEvent) error {

    if tablename == "" {
        return errors.New("Tablename cannot be empty string")
    }

    table := e.tables[tablename]

    if table == nil {
        return errors.New("No table registered with name: " + tablename)
    }

    matchstatus, err := MapNewMatchToMatchStatus(newmatch)

    if err != nil {
        return errors.New("Could not map NewMatchEvent.")
    }

    return table.NewMatch(matchstatus)
}

func (e *Engine) GoalEvent(tablename string, goal events.Goal) error {
    match := e.tables[tablename].GetCurrentMatchStatus()
    matchstatus, err := MapGoalToMatchStatus(match, goal)

    if err != nil {
        return err
    }

    return e.tables[tablename].UpdateMatch(matchstatus)
}
