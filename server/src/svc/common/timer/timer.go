package timer

import (
	"fmt"
	"log"
	"time"
)

type DurationType int32

const (
	DuraMin DurationType = 1 + iota
	Minute
	Hour
	Day
	Week
	Month
	Year
	DuraMax
)

const durationDay int64 = 3600 * 24
const durationWeek int64 = durationDay * 7
const timeFormat string = "2006-01-02 15:04:05"

type Duration struct {
	Date     string
	DuraType DurationType
	Number   int64
}

type Timer struct {
	Duration
	trigger  int64
	handlers []func()
}

func (self *Timer) getNextTime() {

	switch self.DuraType {

	case Day:
		self.trigger += durationDay * self.Number
	case Week:
		self.trigger += durationWeek * self.Number

	}

}

func (self *Timer) getFirstTime() bool {

	tm, _ := time.ParseInLocation(timeFormat, self.Date, time.Local)

	switch self.DuraType {

	case Day:
		dtm := tm.Unix()
		now := time.Now().Unix()
		if dtm > now {
			self.trigger = dtm
		} else {
			dura := durationDay * self.Number
			self.trigger = dtm + ((now-dtm)/dura+1)*dura
		}
		return true

	case Week:
		dtm := tm.Unix()
		now := time.Now().Unix()
		if dtm > now {
			self.trigger = dtm
		} else {
			dura := durationWeek * self.Number
			self.trigger = dtm + ((now-dtm)/dura+1)*dura
		}
		return true
	}

	return false

}

func (self *Timer) check() {

	t := time.Now().Unix()

	if t < self.trigger {
		return
	}

	for _, handler := range self.handlers {
		handler()
	}

	self.getNextTime()

}

func SetTimer(duration Duration, handlers ...func()) interface{} {

	tmr := Timer{
		Duration: duration,
	}

	_, err := time.ParseInLocation(timeFormat, duration.Date, time.Local)
	if err != nil {
		log.Println("set timer error:", err)
		return nil
	}
	if duration.Number <= 0 {
		log.Println("set timer error: negative duration number")
		return nil
	}
	if duration.DuraType <= DuraMin || duration.DuraType >= DuraMax {
		log.Println("set timer error: unregister duration type")
		return nil
	}

	if !tmr.getFirstTime() {
		log.Println("set timer error: duration type has no complent")
		return nil
	}

	for _, handler := range handlers {
		tmr.handlers = append(tmr.handlers, handler)
	}

	log.Println("timer:", time.Unix(tmr.trigger, 0).Format(timeFormat))

	return tmr.check

}

func DayScheduleTask(hour int32, handlers ...func()) interface{} {

	date := fmt.Sprintf("1970-01-01 %02d:00:00", hour)

	return SetTimer(Duration{
		Date:     date,
		DuraType: Day,
		Number:   1,
	}, handlers...)

}

func WeekScheduleTask(hour int32, week int32, handlers ...func()) interface{} {

	date := fmt.Sprintf("1970-01-08 %02d:00:00", hour)

	tm, _ := time.ParseInLocation(timeFormat, date, time.Local)

	tweek := int32(tm.Weekday())
	if tweek == 0 {
		tweek = 7
	}

	tmSec := tm.Unix()
	tmSec += int64(week-tweek) * durationDay

	return SetTimer(Duration{
		Date:     time.Unix(tmSec, 0).Format(timeFormat),
		DuraType: Week,
		Number:   1,
	}, handlers...)

}
