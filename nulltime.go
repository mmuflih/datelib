package datelib

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"strings"
	"time"
)

/**
 * Created by M. Muflih Kholidin
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {

	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}
	log.Println(s)

	t, err := time.Parse(YMD, s)
	log.Println(t, err)

	if err != nil {
		print(err)
		t, err = time.Parse(ISO1, s)
		if err != nil {
			return err
		}
	}

	newNT := new(NullTime)
	newNT.Time = t
	newNT.Valid = true
	*nt = *newNT
	return nil
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(nt)
}
