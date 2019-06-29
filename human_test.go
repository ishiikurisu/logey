package logey

import (
    "testing"
    "time"
    "sort"
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

func TestLogeyCanLoadDates(t *testing.T) {
    testCase := "電子レンジ 円500 2019-01-01"
    expecteDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
    expectedEntry := NewEntry("電子レンジ", 500, make([]string, 0), expecteDate)
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
    resultDate := resultEntry.When
    if expecteDate.Year() != resultDate.Year() {
        t.Error("Loaded wrong year")
    }
    if expecteDate.Month() != resultDate.Month() {
        t.Error("Loaded wrong month")
    }
    if expecteDate.Day() != resultDate.Day() {
        t.Error("Loaded wrong day")
    }
}

func TestLogeyCanLoadTags(t *testing.T) {
    testCase := "-€30 #nebenkosten Rasieren, bitte!"
    expectedTags := []string{
        "nebenkosten",
    }
    resultEntry, oops := Understand(testCase)
    if oops != nil {
        t.Error("Failed valid test case")
    }
    if !compareStringArrays(resultEntry.Where, expectedTags) {
        t.Error("Loaded wrong tags")
    }

    testCase = "-$890 #fixkosten #miete Miete"
    expectedTags = []string{
        "miete",
        "fixkosten",
    }
    resultEntry, oops = Understand(testCase)
    if oops != nil {
        t.Error("Failed valid test case")
    }
    if !compareStringArrays(resultEntry.Where, expectedTags) {
        t.Error("Loaded wrong tags")
    }
}

func compareStringArrays(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }

    sort.Strings(a)
    sort.Strings(b)
    for i, x := range a {
        y := b[i]
        if x != y {
            return false
        }
    }
    
    return true
}
