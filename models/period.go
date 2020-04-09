package fhir

// Period ...
type Period struct {
	Start *FHIRDateTime `bson:"start,omitempty" json:"start,omitempty"`
	End   *FHIRDateTime `bson:"end,omitempty" json:"end,omitempty"`
}
