package fhir

// ParameterDefinition ... // TODO Write proper comment
type ParameterDefinition struct {
	Name          string     `bson:"name,omitempty" json:"name,omitempty"`
	Use           string     `bson:"use,omitempty" json:"use,omitempty"`
	Min           *int32     `bson:"min,omitempty" json:"min,omitempty"`
	Max           string     `bson:"max,omitempty" json:"max,omitempty"`
	Documentation string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type          string     `bson:"type,omitempty" json:"type,omitempty"`
	Profile       *Canonical `bson:"profile,omitempty" json:"profile,omitempty"`
}
