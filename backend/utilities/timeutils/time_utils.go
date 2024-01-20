package timeutils

import (
	"fmt"
	"time"
)

var LongDayNames = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var ShortDayNames = [...]string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var ShortMonthNames = [...]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var LongMonthNames = [...]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}
var localOffsetInMins int

func Init() {
	_, l := time.Now().Zone()
	localOffsetInMins = l / 60
}

func EndOfDayTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 1e9-1, t.Location())
}

func StartOfDayTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func CreatedLastUpdDateTimeForCenter(t time.Time, utcOffsetInMinutes int) time.Time {
	return t.Add(time.Minute * time.Duration(utcOffsetInMinutes-localOffsetInMins))
}

func EmailSMSScheduleDateTimeForCenter(t time.Time, utcOffsetInMinutes int) time.Time {
	return t.Add(time.Minute * time.Duration(localOffsetInMins-utcOffsetInMinutes))
}

func StartOfMonthTime(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonthTime(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)
}

func FormatDisplayDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%s %d %d", ShortMonthNames[m-1], d, y)
}

func FormatDisplayTime(t time.Time) string {
	h, m, _ := t.Clock()
	suffix := "am"
	if h >= 12 {
		suffix = "pm"
		if h > 12 {
			h -= 12
		}
	}
	return fmt.Sprintf("%d:%02d %s", h, m, suffix)
}

func FormatDisplayDurationCompact(d time.Duration) string {
	mins := d / time.Minute
	if mins < 60 {
		return fmt.Sprintf("%dm", mins)
	}
	hours := 0
	for mins >= 60 {
		hours++
		mins -= 60
	}
	if mins == 0 {
		return fmt.Sprintf("%dh", hours)
	}
	return fmt.Sprintf("%dh %dm", hours, mins)
}
