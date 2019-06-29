package logey

import (
	"time"
)

func Understand(human string) Entry {
	outlet := NewEntry("Besoldung", 1000, make([]string, 0), time.Now())
	return outlet
}
