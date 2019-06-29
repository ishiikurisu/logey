package logey

import (
    "testing"
    "time"
)

func TestLogeyCanConvertSimpleStringIntoEntries(t *testing.T) {
    testCase := "Besoldung $1000"
    expectedEntry := NewEntry("Besoldung", 1000, make([]string, 0), time.Now())
    resultEntry := Understand(testCase)

    if expectedEntry.How != resultEntry.How {
        t.Error("Descriptions don't match")
    }
    if expectedEntry.HowMuch != resultEntry.HowMuch {
        t.Error("Values don't match")
    }
}
