package util

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(year int, month int, day int) Date {
	date := Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
	date.correct()
	return date
}

func (date *Date) correct() {
	if date.Year < 1970 {
		date.Year = 1970
	}

	if date.Month < 1 {
		date.Month = 1
	} else if date.Month > 12 {
		date.Month = 12
	}

	maxDay := time.Date(
		date.Year,
		time.Month(date.Month)+1,
		0,
		0, 0, 0, 0,
		time.UTC,
	).Day()

	if date.Day < 1 {
		date.Day = 1
	} else if date.Day > maxDay {
		date.Day = maxDay
	}
}

func (date *Date) ToString() string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day)
}
