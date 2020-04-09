package fhir

// Reference is a FHIR resource Reference
type Reference struct {
	Reference    string `bson:"reference,omitempty" json:"reference,omitempty"`
	Display      string `bson:"display,omitempty" json:"display,omitempty"`
	Type         string `bson:"type,omitempty" json:"type,omitempty"`
	ReferencedID string `bson:"referenceid,omitempty" json:"referenceid,omitempty"`
	External     *bool  `bson:"external,omitempty" json:"external,omitempty"`
}
