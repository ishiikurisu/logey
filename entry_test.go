package logey

import (
    "testing"
    "fmt"
    "time"
)

func TestIfEntryCanBeParsedAsString(t *testing.T) {
    expectedHow := "Hausprodukten"
    expectedHowMuch := -80.00
    expectedWhere := []string{"nebenkosten"}
    expectedWhen := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
    expectedPayload := `["Hausprodukten",-80.00,["nebenkosten"],"2020-01-01T00:00:00Z"]`

    entry := NewEntry(expectedHow, expectedHowMuch, expectedWhere, expectedWhen)
    result := entry.ToString()
    if expectedPayload != result {
        t.Error(fmt.Sprintf("Entry can't become a string.\nExpected:\t%s\nGotten:   \t%s\n", expectedPayload, result))
    }

    entry = LoadEntryFromString(result)
    if entry.How != expectedHow {
        t.Error(fmt.Sprintf("Wrong description.\nExpected:\t%s\nGotten:\t\t%s\n", expectedHow, entry.How))
    }
    if entry.HowMuch != expectedHowMuch {
        t.Error(fmt.Sprintf("Wrong value was parsed.\nExpected:\t%f\nGotten:\t\t%f\n", expectedHowMuch, entry.HowMuch))
    }
    if !compareLists(entry.Where, expectedWhere) {
        t.Error("wrong tags\n")
    }
    if entry.When != expectedWhen {
        t.Error(fmt.Sprintf("Wrong dates.\nExpected:\t%#v\nGotten:\t\t%#v\n", expectedWhen, entry.When))
    }
}

func compareLists(x, y []string) bool {
    limitX := len(x)
    limitY := len(y)

    if limitX != limitY {
        return false
    }

    for i := 0; i < limitX; i++ {
        if x[i] != y[i] {
            return false
        }
    }

    return true
}
