package datelib

import "time"

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
