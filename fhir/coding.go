package fhir

// Coding ...
type Coding struct {
	System       string `bson:"system,omitempty" json:"system,omitempty"`
	Version      string `bson:"version,omitempty" json:"version,omitempty"`
	Code         string `bson:"code,omitempty" json:"code,omitempty"`
	Display      string `bson:"display,omitempty" json:"display,omitempty"`
	UserSelected *bool  `bson:"userSelected,omitempty" json:"userSelected,omitempty"`
}
