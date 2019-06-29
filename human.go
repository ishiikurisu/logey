package logey

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MONEY_REGEX    = "(-?[$€円¥][0-9]*)"
	CURRENCY_REGEX = "[$€円¥]"
	DATE_REGEX     = "([0-9]{4}-0?[0-9]{1,2}-0?[0-9]{1,2})"
)

func Understand(human string) (Entry, error) {
	inlet := []byte(human)
	outlet := NewEntry("", 0, make([]string, 0), time.Now())

	moneyRegex, _ := regexp.Compile(MONEY_REGEX)
	currencyRegex, _ := regexp.Compile(CURRENCY_REGEX)
	dateRegex, oops := regexp.Compile(DATE_REGEX)
	if oops != nil {
		return outlet, oops
	}

	rawMoney := moneyRegex.Find(inlet)
	if rawMoney == nil {
		return outlet, errors.New("Unable to find monetary value in description")
	}
	valueDescription := currencyRegex.ReplaceAll(rawMoney, []byte(""))
	howMuch, oops := strconv.ParseFloat(string(valueDescription), 64)
	if oops != nil {
		return outlet, oops
	}

	rawDate := string(dateRegex.Find(inlet))
	when := time.Now()
	if len(rawDate) > 0 {
		pieces := strings.Split(rawDate, "-")
		year, _ := strconv.Atoi(pieces[0])
		month, _ := strconv.Atoi(pieces[1])
		day, _ := strconv.Atoi(pieces[2])
		when = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	how := strings.ReplaceAll(human, string(rawMoney), "")
	how = strings.ReplaceAll(how, rawDate, "")

	outlet.How = strings.TrimSpace(how)
	outlet.HowMuch = howMuch
	outlet.When = when

	return outlet, nil
}
