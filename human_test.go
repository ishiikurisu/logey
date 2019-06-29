package logey

import (
    "testing"
    "time"
)

func TestLogeyCanConvertSimpleStringIntoEntries(t *testing.T) {
    testCase := "Besoldung $1000"
    expectedEntry := NewEntry("Besoldung", 1000, make([]string, 0), time.Now())
    resultEntry, oops := Understand(testCase)
    if oops != nil {
        t.Error("Failed valid test case")
    }
    if expectedEntry.How != resultEntry.How {
        t.Error("Descriptions don't match")
    }
    if expectedEntry.HowMuch != resultEntry.HowMuch {
        t.Error("Values don't match")
    }

    testCase = "This test case contains no money description"
    _, oops = Understand(testCase)
    if oops == nil {
        t.Error("Succeeded invalid test case")
    }

    testCase = "-€500 Steuer"
    expectedEntry = NewEntry("Steuer", -500, make([]string, 0), time.Now())
    resultEntry, oops = Understand(testCase)
    if oops != nil {
        t.Error("Failed valid test case")
    }
    if expectedEntry.How != resultEntry.How {
        t.Error("Descriptions don't match")
    }
    if expectedEntry.HowMuch != resultEntry.HowMuch {
        t.Error("Values don't match")
    }
}
