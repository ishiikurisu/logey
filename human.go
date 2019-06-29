package logey

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Understand(human string) (Entry, error) {
	outlet := NewEntry("", 0, make([]string, 0), time.Now())

	moneyRegex, oops := regexp.Compile("(-?[$€円¥][0-9]*)")
	if oops != nil {
		return outlet, oops
	}
	currencyRegex, oops := regexp.Compile("[$€円¥]")
	if oops != nil {
		return outlet, oops
	}

	rawMoney := moneyRegex.Find([]byte(human))
	if rawMoney == nil {
		return outlet, errors.New("Unable to find monetary value in description")
	}
	valueDescription := currencyRegex.ReplaceAll(rawMoney, []byte(""))
	howMuch, oops := strconv.ParseFloat(string(valueDescription), 64)
	if oops != nil {
		return outlet, oops
	}

	how := strings.ReplaceAll(human, string(rawMoney), "")

	outlet.How = strings.TrimSpace(how)
	outlet.HowMuch = howMuch

	return outlet, nil
}
