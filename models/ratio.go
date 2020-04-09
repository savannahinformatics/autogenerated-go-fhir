package fhir

// Ratio ...
type Ratio struct {
	Numerator   *Quantity `bson:"numerator,omitempty" json:"numerator,omitempty"`
	Denominator *Quantity `bson:"denominator,omitempty" json:"denominator,omitempty"`
}
