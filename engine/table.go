package engine

import (
    "github.com/krafthack/ifooz-engine/liveview"
)

func CreateTable(tablename string) *Table {
    // TODO configurable buffer size
    updatesbuffer := 10
    updateschannel := make(chan *liveview.MatchStatus, updatesbuffer)

    table := &Table { updateschannel, nil }

    return table
}

type Table struct {
    updates chan *liveview.MatchStatus
    current *liveview.MatchStatus
}

func (t *Table) GetUpdatesChannel() chan *liveview.MatchStatus {
    return t.updates
}

func (t *Table) NewMatch(newmatch *liveview.MatchStatus) error {
    t.current = newmatch
    t.updates <- newmatch
    return nil
}

func (t *Table) UpdateMatch(match *liveview.MatchStatus) error {
    t.current = match
    t.updates <- match
    return nil
}

func (t *Table) GetCurrentMatchStatus() *liveview.MatchStatus {
    return t.current
}
