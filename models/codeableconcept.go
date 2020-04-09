package fhir

// CodeableConcept ...
type CodeableConcept struct {
	Coding []Coding `bson:"coding,omitempty" json:"coding,omitempty"`
	Text   string   `bson:"text,omitempty" json:"text,omitempty"`
}
