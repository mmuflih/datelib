package datelib

import (
	"database/sql/driver"
	"encoding/json"
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

	t, err := time.Parse(YMD, s)
	if err != nil {
		print(err)
		t, err = time.Parse(ISO1, s)
		if err != nil {
			return err
		}
	}

	nt.Time = t
	nt.Valid = true
	return nil
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Time  string `json:"time"`
		Valid bool   `json:"valid"`
	}{
		Time:  nt.Time.Format(time.RFC3339),
		Valid: nt.Valid,
	})
}
