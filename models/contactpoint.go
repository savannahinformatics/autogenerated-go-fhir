package fhir

// ContactPoint ... // TODO Write proper comment
type ContactPoint struct {
	System string  `bson:"system,omitempty" json:"system,omitempty"`
	Value  string  `bson:"value,omitempty" json:"value,omitempty"`
	Use    string  `bson:"use,omitempty" json:"use,omitempty"`
	Rank   *uint32 `bson:"rank,omitempty" json:"rank,omitempty"`
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}
