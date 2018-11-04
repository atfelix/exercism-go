package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

var weekDayConversion = map[time.Weekday]int {
	time.Monday: 0, 
    time.Tuesday: 1, 
    time.Wednesday: 2,
    time.Thursday: 3,
    time.Friday: 4,
    time.Saturday:5,
    time.Sunday: 6,
}

func Day(week WeekSchedule, day time.Weekday, month time.Month, year int) int {
	switch week {
	case Teenth: return teenth(year, month, day)
	case Last: return last(year, month, day)
	default: return date(year, month, day, week)
	}
}

func teenth(year int, month time.Month, day time.Weekday) int {
	return edgeCase(year, month, 13, day)
}

func last(year int, month time.Month, day time.Weekday) int {
	return edgeCase(year, month, lastDay(year, month) - 6, day)
}

func date(year int, month time.Month, day time.Weekday, week WeekSchedule) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for date.Weekday() != day {
		date = date.AddDate(0, 0, 1)
	}

	for index := 0; index < int(week); index++ {
		date = date.AddDate(0, 0, 7)
	}

	return date.Day()
}

func edgeCase(year int, month time.Month, startDate int, day time.Weekday) int {
	date := time.Date(year, month, startDate, 0, 0, 0, 0, time.UTC)
	date = date.AddDate(0, 0, timeDelta(day, date.Weekday()))
	return date.Day()
}

func timeDelta(d1, d2 time.Weekday) int {
	return (weekDayConversion[d1] - weekDayConversion[d2] + 7) % 7
}

func lastDay(year int, month time.Month) int {
	monthLengths := map[time.Month]int{
		time.January: 31, 
		time.February: 28,
		time.March: 31,
		time.April: 30,
		time.May: 31,
		time.June: 30,
		time.July: 31,
		time.August: 31,
		time.September: 30,
		time.October: 31,
		time.November: 30,
		time.December: 31,
	}
	var leapDay int
	if isLeapMonth(year, month) {
		leapDay = 1
	} else {
		leapDay = 0
	}
	return monthLengths[month] + leapDay
}

func isLeapMonth(year int, month time.Month) bool {
	return month == time.February && isLeapYear(year)
}

func isLeapYear(year int) bool {
	return year % 400 == 0 ||
			(year % 100 != 0 &&
			 year % 4 == 0)
}

