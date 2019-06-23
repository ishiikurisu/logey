package logey

import (
    "testing"
    "fmt"
    "time"
)

func TestIfEntryCanBeParsedAsString(t *testing.T) {
    now := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
    entry := NewEntry("Hausprodukten", -80, []string{"nebenkosten"}, now)
    expected := `["Hausprodukten",-80.00,["nebenkosten"],"2020-01-01T00:00:00Z"]`
    result := entry.ToString()
    if expected != result {
        t.Error(fmt.Sprintf("Entry can't become a string.\nExpected:\t %s\nGotten:\t%s\n", expected, result))
    }
}
