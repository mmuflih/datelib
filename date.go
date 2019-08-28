package datelib

import (
	"fmt"
	"time"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 28/08/19 05.57
 */

func NewTZ(tz string) time.Time{
	now := time.Now()
	t, err := time.Parse(YMD_HMS_TZ, now.Format(YMD_HMS) + tz)
	if err != nil {
		fmt.Println("Error create date with timezone")
		return now
	}
	return t
}
