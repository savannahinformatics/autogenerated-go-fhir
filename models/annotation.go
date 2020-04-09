package fhir

// Annotation is a FHIR Annotation structure
type Annotation struct {
	AuthorReference *Reference    `bson:"authorReference,omitempty" json:"authorReference,omitempty"`
	AuthorString    string        `bson:"authorString,omitempty" json:"authorString,omitempty"`
	Time            *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
	Text            string        `bson:"text,omitempty" json:"text,omitempty"`
}
