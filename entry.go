package logey

import (
	"fmt"
	"strings"
	"time"
)

// Defines an entry, the fundamental element in the money log.
type Entry struct {
	// How the money was spent or earned.
	How string

	// The money amount of this entry.
	HowMuch float64

	// Which categories this entry belong to
	Where []string

	// When this entry happened
	When time.Time
}

// Creates an entry from a description and a value.
func NewEntry(description string, value float64, tags []string, when time.Time) Entry {
	return Entry{
		How:     description,
		HowMuch: value,
		Where:   tags,
		When:    when,
	}
}

// Gets the current header format for an entry
func GetEntryFormat() string {
	return `["how", "how much", "where", "when"]`
}

// Creates a JSONL entry
func (entry Entry) ToString() string {
	where := fmt.Sprintf("[%#v]", strings.Join(entry.Where, ","))

	whenBytes, oops := entry.When.MarshalJSON()
	if oops != nil {
		panic(oops)
	}
	when := string(whenBytes)

	return fmt.Sprintf(
		"[\"%s\",%.2f,%s,%s]",
		entry.How,
		entry.HowMuch,
		where,
		when,
	)
}
