package logey

import (
    "testing"
    "time"
    "fmt"
)

func giveExampleLog() Log {
    var log Log
    now := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

    log.DescribeEntry("letzte monat", 315, []string{"anlung"}, now)
    log.DescribeEntry("hausprodukten", -80, []string{"nebenkosten"}, now)
    log.DescribeEntry("besoldung", 500, []string{"input"}, now)

    return log
}

func TestCanCalculateBalanceFromLog(t *testing.T) {
    log := giveExampleLog()

    if log.Balance != 735 {
        t.Error(fmt.Sprintf("Can't calculate balance. %.2f != %.2f", 735.00, log.Balance))
    }
}

func TestCanImportAndExportLogs(t *testing.T) {
    log := giveExampleLog()

    expected := `["how", "how much", "where", "when"]
["letzte monat",315.00,["anlung"],"2020-01-01T00:00:00Z"]
["hausprodukten",-80.00,["nebenkosten"],"2020-01-01T00:00:00Z"]
["besoldung",500.00,["input"],"2020-01-01T00:00:00Z"]
`
    result := log.Export()
    if expected != result {
        t.Error("Could not export log to string")
    }
}
