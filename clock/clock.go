package clock

import (
	"fmt"
)

const (
	numberOfMinutesAnHour = 60
	numberOfHoursPerDay = 24
	numberOfMinutesPerDay = numberOfMinutesAnHour * numberOfHoursPerDay
)

type Clock struct {
	hour, minute int
}

func (aClock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", aClock.hour, aClock.minute)
}

func New(hour, minute int) Clock {
	minutes := numberOfMinutes(hour, minute) 
	return Clock{
		hour: minutes / numberOfMinutesAnHour % numberOfHoursPerDay,
		minute: minutes % numberOfMinutesAnHour,
	}
}

func numberOfMinutes(hour, minute int) int {
	minutes := (hour * 60 + minute) % numberOfMinutesPerDay
	minutes = (minutes + numberOfMinutesPerDay) % numberOfMinutesPerDay
	return minutes
}

func (aClock Clock) Add(minutes int) Clock {
	return New(aClock.hour, aClock.minute + minutes)
}

func (aClock Clock) Subtract(minutes int) Clock {
	return aClock.Add(-minutes)
}