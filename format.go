package datelib

import (
	"fmt"
	"time"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 28/08/19 05.47
 */

const YMD = "2006-01-02"
const YMD_SHORT = "06-01-02"
const YMD_HMS = "2006-01-02 15:04:05"

var YMD_HMS_WS = "20060102150405"

const YMD_WS = "20060102"
const YMD_SHORT_WS = "060102"
const YMD_HM = "2006-01-02 15:04"
const YMD_HMS_TZ = "2006-01-02 15:04:05-0700"
const DMY = "02 January 2006"
const DMY_HMS = "02 January 2006 15:04:05"
const DMY_HM = "02 January 2006 15:04"
const HMS = "15:04:05"
const HM = "15:04"
const ISO1 = time.RFC3339

const DD_MM_YYYY = "02-Jan-2006"
const D_M_Y = "02-01-2006"

type RomanNumber struct {
	Year  string
	Month string
	Date  string
}

func RomanNumberFromTime(t time.Time) RomanNumber {
	return RomanNumber{
		Year:  romanNumberParser(t.Year()),
		Month: romanNumberParser(int(t.Month())),
		Date:  romanNumberParser(t.Day()),
	}
}

func romanNumberParser(num int) string {
	if num == 0 {
		return ""
	}
	number := []int{1000, 500, 100, 50, 10, 5, 1}
	romanNumber := []string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	var result string

	for num > 0 {
		for _, k := range number {
			// fmt.Println("IN NUM", k)
			if num == 0 {
				break
			}
			km1 := k - 1
			if num == km1 {
				// fmt.Println(num, "=", km1)
				result += "I" + romanNumber[k]
				num -= km1
				break
			}
			// if num/k >= 4 {
			// 	num -= 4 * num
			// 	break
			// }
			if num/k >= 1 {
				// fmt.Println(num, "/", k, num/k)
				count := num / k
				for i := 0; i < count; i++ {
					result += romanNumber[k]
				}
				// fmt.Println("MOD", num, "mod", k, "=", num%k)
				num %= k
				break
			}
		}
	}
	return result
}

func RomanNumberParser2(num int) string {
	if num == 0 {
		return ""
	}
	number := []int{1000, 500, 100, 50, 10, 5, 1}
	romanNumber := []string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	var result string
	for _, k := range number {
		fmt.Println("NUM", num, num+1, k)
		km1 := k - 1
		if num == km1 {
			fmt.Println("----", num)
			result += "I" + romanNumber[k]
			num = num - km1
			continue
		}
		if num/k >= 1 {
			count := num / k
			for i := 0; i < count; i++ {
				result += romanNumber[k]
			}
			num %= k
		}
		continue
	}
	return result
}
