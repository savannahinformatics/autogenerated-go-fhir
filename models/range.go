package fhir

// Range ... // TODO Write proper comment
type Range struct {
	Low  *Quantity `bson:"low,omitempty" json:"low,omitempty"`
	High *Quantity `bson:"high,omitempty" json:"high,omitempty"`
}
