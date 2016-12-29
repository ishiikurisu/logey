package moneylog

import "fmt"
import "strings"

type Entry struct {
    Description string
    Value float64
}

func NewEntry(description string, value float64) Entry {
    return Entry{Description: description, Value: value}
}

func EntryFromString(raw string) Entry {
    var contents []string = strings.Split(raw, ":")
    var description string = contents[0]
    var value float64
    fmt.Sscanf(contents[1], "%f", &value)
    return NewEntry(description, value)
}
