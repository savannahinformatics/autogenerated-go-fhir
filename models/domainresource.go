package fhir

// DomainResource ... // TODO Write proper comment
type DomainResource struct {
	Resource          `bson:",inline"`
	Text              *Narrative    `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []interface{} `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
