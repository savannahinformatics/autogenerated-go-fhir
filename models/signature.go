package fhir

// Signature ...
type Signature struct {
	Type         []Coding      `bson:"type,omitempty" json:"type,omitempty"`
	When         *FHIRDateTime `bson:"when,omitempty" json:"when,omitempty"`
	Who          *Reference    `bson:"who,omitempty" json:"who,omitempty"`
	OnBehalfOf   *Reference    `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	TargetFormat string        `bson:"targetFormat,omitempty" json:"targetFormat,omitempty"`
	SigFormat    string        `bson:"sigFormat,omitempty" json:"sigFormat,omitempty"`
	Data         string        `bson:"data,omitempty" json:"data,omitempty"`
}
