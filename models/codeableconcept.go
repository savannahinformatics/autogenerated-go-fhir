package fhir

// CodeableConcept ... // TODO Write proper comment
type CodeableConcept struct {
	Coding []Coding `bson:"coding,omitempty" json:"coding,omitempty"`
	Text   string   `bson:"text,omitempty" json:"text,omitempty"`
}
