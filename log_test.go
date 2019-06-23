package logey

import (
    "testing"
    "time"
    "fmt"
)

func TestCanCalculateBalanceFromLog(t *testing.T) {
    var log Log

    log.DescribeEntry("letzte monat", 315, []string{"anlung"}, time.Now())
    log.DescribeEntry("hausprodukten", -80, []string{"nebenkosten"}, time.Now())
    log.DescribeEntry("besoldung", 500, []string{"input"}, time.Now())

    if log.Balance != 735 {
        t.Error(fmt.Sprintf("Can't calculate balance. %.2f != %.2f", 735.00, log.Balance))
    }
}
