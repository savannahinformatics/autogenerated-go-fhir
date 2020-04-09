package fhir

// Quantity ... // TODO Write proper comment
type Quantity struct {
	Value      *float64 `bson:"value,omitempty" json:"value,omitempty"`
	Comparator string   `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Unit       string   `bson:"unit,omitempty" json:"unit,omitempty"`
	System     string   `bson:"system,omitempty" json:"system,omitempty"`
	Code       string   `bson:"code,omitempty" json:"code,omitempty"`
}
