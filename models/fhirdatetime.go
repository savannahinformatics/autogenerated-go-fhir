package fhir

import (
	"encoding/json"
	"time"
)

// Precision ...
type Precision string

const (
	// Date ...
	Date = "date"
	// Timestamp ...
	Timestamp = "timestamp"
)

// FHIRDateTime ...
type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}

// UnmarshalJSON ...
func (f *FHIRDateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) <= 12 {
		f.Precision = Precision("date")
		f.Time, err = time.ParseInLocation("\"2006-01-02\"", string(data), time.Local)
	} else {
		f.Precision = Precision("timestamp")
		f.Time = time.Time{}
		f.Time.UnmarshalJSON(data)
	}
	return err
}

// MarshalJSON ...
func (f FHIRDateTime) MarshalJSON() ([]byte, error) {
	if f.Precision == Timestamp {
		return json.Marshal(f.Time.Format(time.RFC3339))
	} else {
		return json.Marshal(f.Time.Format("2006-01-02"))
	}
}
