package fhir

// SampledData ... // TODO Write proper comment
type SampledData struct {
	Origin     *Quantity `bson:"origin,omitempty" json:"origin,omitempty"`
	Period     *float64  `bson:"period,omitempty" json:"period,omitempty"`
	Factor     *float64  `bson:"factor,omitempty" json:"factor,omitempty"`
	LowerLimit *float64  `bson:"lowerLimit,omitempty" json:"lowerLimit,omitempty"`
	UpperLimit *float64  `bson:"upperLimit,omitempty" json:"upperLimit,omitempty"`
	Dimensions *uint32   `bson:"dimensions,omitempty" json:"dimensions,omitempty"`
	Data       string    `bson:"data,omitempty" json:"data,omitempty"`
}
