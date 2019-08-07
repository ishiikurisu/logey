package logey

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MONEY_REGEX    = "-?[$€円¥][0-9]*"
	CURRENCY_REGEX = "[$€円¥]"
	DATE_REGEX     = "([0-9]{4}-0?[0-9]{1,2}-0?[0-9]{1,2})"
	TAG_REGEX      = "#([^\\x00-\\x7F]|[\\w])+"
)

func Understand(human string) (Entry, error) {
	inlet := []byte(human)
	outlet := NewEntry("", 0, make([]string, 0), time.Now())

	moneyRegex := regexp.MustCompile(MONEY_REGEX)
	currencyRegex := regexp.MustCompile(CURRENCY_REGEX)
	dateRegex := regexp.MustCompile(DATE_REGEX)
	tagRegex := regexp.MustCompile(TAG_REGEX)

	rawMoney := moneyRegex.Find(inlet)
	if rawMoney == nil {
		return outlet, errors.New("Unable to find monetary value in description")
	}
	valueDescription := currencyRegex.ReplaceAll(rawMoney, []byte(""))
	howMuch, oops := strconv.ParseFloat(string(valueDescription), 64)
	if oops != nil {
		return outlet, oops
	}
	how := strings.ReplaceAll(human, string(rawMoney), "")

	when := time.Now()
	rawDate := string(dateRegex.Find(inlet))
	if len(rawDate) > 0 {
		pieces := strings.Split(rawDate, "-")
		year, _ := strconv.Atoi(pieces[0])
		month, _ := strconv.Atoi(pieces[1])
		day, _ := strconv.Atoi(pieces[2])
		when = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}
	how = strings.ReplaceAll(how, rawDate, "")

	where := make([]string, 0)
	rawTags := tagRegex.FindAll(inlet, -1)
	if rawTags != nil {
		for _, rawTag := range rawTags {
			tag := string(rawTag)
			where = append(where, tag[1:])
			how = strings.ReplaceAll(how, tag, "")
		}
	}

	outlet.How = strings.TrimSpace(how)
	outlet.HowMuch = howMuch
	outlet.Where = where
	outlet.When = when

	return outlet, nil
}
