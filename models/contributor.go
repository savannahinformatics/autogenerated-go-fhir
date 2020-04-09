package fhir

// Contributor ... // TODO Write proper comment
type Contributor struct {
	Type    string          `bson:"type,omitempty" json:"type,omitempty"`
	Name    string          `bson:"name,omitempty" json:"name,omitempty"`
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
}
