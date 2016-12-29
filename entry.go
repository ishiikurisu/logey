package moneylog

type Entry struct {
    Description string
    Value float64
}

func NewEntry(description string, value float64) Entry {
    return Entry{Description: description, Value: value}
}
