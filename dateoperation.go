package datelib

import (
	"fmt"
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 9/2/19 10:31
**/

func AddDate(date time.Time, count int) time.Time {
	return date.AddDate(0, 0, count)
}

func AddMonth(date time.Time, count int) time.Time {
	return date.AddDate(0, count, 0)
}

func AddYear(date time.Time, count int) time.Time {
	return date.AddDate(count, 0, 0)
}

func SubDate(date time.Time, count int) time.Time {
	return date.AddDate(0, 0, -count)
}

func SubMonth(date time.Time, count int) time.Time {
	return date.AddDate(0, -count, 0)
}

func SubYear(date time.Time, count int) time.Time {
	return date.AddDate(-count, 0, 0)
}

func LastDate(date time.Time) time.Time {
	d := date.AddDate(0, 1, 0)
	return d.AddDate(0, 0, -1)
}

func ParseTZ(format, date, tz string) time.Time {
	t, err := time.Parse(format+"-0700", date+tz)
	if err != nil {
		fmt.Println("Error create date with timezone")
		return time.Now()
	}
	return t
}

func Diff(t1, t2 time.Time) time.Duration {
	diff := t1.Sub(t2)
	if diff.Seconds() < 0 {
		return diff * -1
	}
	return diff
}
