package fhir

// Money ... // TODO Write proper comment
type Money struct {
	Value    *float64 `bson:"value,omitempty" json:"value,omitempty"`
	Currency string   `bson:"currency,omitempty" json:"currency,omitempty"`
}
