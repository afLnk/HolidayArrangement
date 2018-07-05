package HolidayArrangement

import (
	"time"
)

const (
	Workday          = 1
	Holiday_Unknown  = 2
	Holiday_Weekend  = 3
	Holiday_Festival = 4
)

type DayType int64

type Calendar struct {
	holidays_ map[int64]DayType
}

var (
	Weekend    = *NewCalendarOnWeekend(nil)
	ChinaState = *NewCalendarOnWeekend(map[int64]DayType{
		// 元旦
		20180101: Holiday_Festival,

		// 春节
		20180211: Workday,
		20180215: Holiday_Festival,
		20180216: Holiday_Festival,
		20180217: Holiday_Festival,
		20180218: Holiday_Festival,
		20180219: Holiday_Festival,
		20180220: Holiday_Festival,
		20180221: Holiday_Festival,
		20180224: Workday,
		// 清明节
		20180405: Holiday_Festival,
		20180406: Holiday_Festival,
		20180407: Holiday_Festival,
		20180408: Workday,
		// 劳动节
		20180428: Workday,
		20180429: Holiday_Festival,
		20180430: Holiday_Festival,
		20180501: Holiday_Festival,
		// 端午节
		20180618: Holiday_Festival,
		// 中秋节
		20180924: Holiday_Festival,
		// 国庆节
		20180929: Workday,
		20180930: Workday,
		20181001: Holiday_Festival,
		20181002: Holiday_Festival,
		20181003: Holiday_Festival,
		20181004: Holiday_Festival,
		20181005: Holiday_Festival,
		20181006: Holiday_Festival,
		20181007: Holiday_Festival,
	})
)

func NewCalendar(holidays map[int64]DayType) *Calendar {
	calendar := &Calendar{holidays_: holidays}
	if calendar.holidays_ == nil {
		calendar.holidays_ = make(map[int64]DayType)
	}

	return calendar
}

func NewCalendarOnWeekend(exceptionDays map[int64]DayType) *Calendar {
	calendar := NewCalendar(nil)

	curDay, _ := time.ParseInLocation("2006-01-02", "2018-01-01", time.Local)
	for i := 0; i < 366; i++ {
		day := 10000*int64(curDay.Year()) + 100*int64(curDay.Month()) + int64(curDay.Day())
		var dayType DayType
		if nil != exceptionDays {
			var ok bool
			dayType, ok = exceptionDays[day]
			if !ok {
				dayType = 0
			}
		}

		if dayType != 0 && dayType != Workday { // today is holiday
			calendar.holidays_[day] = dayType
		}

		weekday := curDay.Weekday()
		if weekday != 0 && weekday != 6 { // today is weekend
			if dayType == 0 {
				calendar.holidays_[day] = Holiday_Weekend
			}
		}

		curDay = curDay.AddDate(0, 0, 1)
	}

	return calendar
}

// create new calendar base on this
func (base *Calendar) Create(exceptionDays map[int64]DayType) *Calendar {
	calendar := NewCalendar(nil)

	for holiday, holidayType := range base.holidays_ {
		dayType, ok := exceptionDays[holiday]
		if !ok {
			calendar.holidays_[holiday] = holidayType
		} else if dayType != Workday {
			calendar.holidays_[holiday] = dayType
		}
	}

	return calendar
}

func (base *Calendar) GetDayType(day int64) DayType {
	dayType, ok := base.holidays_[day]
	if !ok {
		return Workday // if not found in holidays, is workday
	}

	return dayType
}

func (base *Calendar) IsHoliday(day int64) bool {
	dayType := base.GetDayType(day)
	return dayType != Workday && dayType != 0
}

func (base *Calendar) GetHolidays() map[int64]DayType {
	return base.holidays_
}
