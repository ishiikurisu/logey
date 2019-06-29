package logey

import (
	"encoding/json"
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

func LoadEntryFromString(payload string) (Entry, error) {
	var outlet Entry
	var rawEntry []interface{}

	oops := json.Unmarshal([]byte(payload), &rawEntry)
	if oops != nil {
		return outlet, oops
	}

	how := rawEntry[0].(string)

	howMuch := rawEntry[1].(float64)

	rawWhere := rawEntry[2].([]interface{})
	where := make([]string, len(rawWhere))
	for i, tag := range rawWhere {
		where[i] = tag.(string)
	}

	var when time.Time
	rawWhen := fmt.Sprintf("\"%s\"", rawEntry[3].(string))
	oops = when.UnmarshalJSON([]byte(rawWhen))
	if oops != nil {
		return outlet, oops
	}
	when = when.In(time.UTC)

	outlet = NewEntry(how, howMuch, where, when)
	return outlet, nil
}
