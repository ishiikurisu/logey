package moneylog

import "fmt"

type Log struct {
    Entries []Entry
}

func NewLog(first Entry) Log {
    log := Log{Entries: make([]Entry, 1)}
    log.Entries[0] = first
    return log
}

func StartLog(description string, value float64) Log {
    return NewLog(NewEntry(description, value))
}

func (log *Log) Insert(entry Entry) {
    limit := 1 + len(log.Entries)
    newEntries := make([]Entry, limit)

    for i := 0; i < limit - 1; i++ {
        newEntries[i] = log.Entries[i]
    }
    newEntries[limit-1] = entry

    log.Entries = newEntries
}

func (log *Log) Add(description string, value float64) {
    log.Insert(NewEntry(description, value))
}

func (log *Log) ToString() string {
    outlet := "---\n"

    for _, entry := range log.Entries {
        outlet += fmt.Sprintf("%s: %.2f\n", entry.Description, entry.Value)
    }

    return fmt.Sprintf("%s...\n", outlet)
}
