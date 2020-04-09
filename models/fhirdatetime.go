package fhir

import (
	"encoding/json"
	"time"
)

// Precision ... // TODO Write proper comment
type Precision string

const (
	// Date ... // TODO Write proper comment
	Date = "date"
	// Timestamp ... // TODO Write proper comment
	Timestamp = "timestamp"
)

// FHIRDateTime ... // TODO Write proper comment
type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}

// UnmarshalJSON ... // TODO Write proper comment
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

// MarshalJSON ... // TODO Write proper comment
func (f FHIRDateTime) MarshalJSON() ([]byte, error) {
	if f.Precision == Timestamp {
		return json.Marshal(f.Time.Format(time.RFC3339))
	} else {
		return json.Marshal(f.Time.Format("2006-01-02"))
	}
}
