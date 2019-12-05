package datelib_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mmuflih/datelib"
)

type NTtest struct {
	BirthDate datelib.NullTime `json:"birth_date"`
}

func TestUnmarshalYMD(t *testing.T) {
	// PREPARATION
	dummy := `{
      "birth_date":"1997-02-19"
    }`
	expectedNT := new(datelib.NullTime)
	nTime, err := time.Parse(datelib.YMD, "1997-02-19")
	if err != nil {
		panic(err)
	}
	expectedNT.Time = nTime
	expectedNT.Valid = true

	// ACTION
	var ntt NTtest
	json.Unmarshal([]byte(dummy), &ntt)

	// ASSERTION
	if ntt.BirthDate.Time != expectedNT.Time {
		t.Errorf("Expected time %s, but got %s\n", expectedNT.Time.String(), ntt.BirthDate.Time.String())
	}
	if ntt.BirthDate.Valid != expectedNT.Valid {
		t.Errorf("Expected %t bug got %t\n", expectedNT.Valid, ntt.BirthDate.Valid)
	}
}

func TestUnmarshalISO1(t *testing.T) {
	// PREPARATION
	dummy := `{
      "birth_date":"1997-02-19T00:00:00+07:00"
    }`
	expectedNT := new(datelib.NullTime)
	nTime, err := time.Parse(time.RFC3339, "1997-02-19T00:00:00+07:00")
	if err != nil {
		panic(err)
	}
	expectedNT.Time = nTime
	expectedNT.Valid = true

	// ACTION
	var ntt NTtest
	json.Unmarshal([]byte(dummy), &ntt)

	// ASSERTION
	if ntt.BirthDate.Time != expectedNT.Time {
		t.Errorf("Expected time %s, but got %s\n", expectedNT.Time.String(), ntt.BirthDate.Time.String())
	}
	if ntt.BirthDate.Valid != expectedNT.Valid {
		t.Errorf("Expected %t bug got %t\n", expectedNT.Valid, ntt.BirthDate.Valid)
	}
}
