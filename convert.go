package datelib

import "strings"

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 10/1/19 20:30
**/

func ToYMD(date string) string {
	if strings.Contains(date[:4], "-") {
		dates := strings.Split(date, "-")
		return dates[2] + "-" + dates[1] + "-" + dates[0]
	}
	return date
}

func ToDMY(date string) string {
	if !strings.Contains(date[:4], "-") {
		dates := strings.Split(date, "-")
		return dates[2] + "-" + dates[1] + "-" + dates[0]
	}
	return date
}
