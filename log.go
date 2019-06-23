package logey

import (
	"fmt"
	"time"
)

// This is the log, which will store and deal with entries
type Log struct {
	// The list of entries
	Entries []Entry

	// The total balance, should be updated after each addition
	Balance float64
}

// Creates a new log
func NewLog() *Log {
	var log Log
	log.Balance = 0
	return &log
}

// Adds a new entry in the log
func (log *Log) AddEntry(entry Entry) {
	log.Entries = append(log.Entries, entry)
	log.Balance += entry.HowMuch
}

// Adds a new entry in the log while describing the entry
func (log *Log) DescribeEntry(how string, howMuch float64, where []string, when time.Time) {
	entry := NewEntry(how, howMuch, where, when)
	log.AddEntry(entry)
}

// Turns a log into a JSON lines table
func (log Log) Export() string {
	outlet := fmt.Sprintf("%s\n", GetEntryFormat())

	for _, entry := range log.Entries {
		outlet = fmt.Sprintf("%s%s\n", outlet, entry.ToString())
	}

	return outlet
}
