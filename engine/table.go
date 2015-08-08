package engine

import (
    "github.com/krafthack/ifooz-engine/liveview"
)

func CreateTable(tablename string) *Table {
    // TODO configurable buffer size
    updatesbuffer := 10
    updateschannel := make(chan *liveview.MatchStatus, updatesbuffer)

    table := &Table { updateschannel }

    return table
}

type Table struct {
    updates chan *liveview.MatchStatus
}

func (t *Table) GetUpdatesChannel() chan *liveview.MatchStatus {
    return t.updates
}

func (t *Table) NewMatch(newmatch *liveview.MatchStatus) error {

    t.updates <- newmatch
    return nil
}
