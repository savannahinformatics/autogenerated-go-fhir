package fhir

// Identifier ... // TODO Write proper comment
type Identifier struct {
	Use      string           `bson:"use,omitempty" json:"use,omitempty"`
	Type     *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	System   string           `bson:"system,omitempty" json:"system,omitempty"`
	Value    string           `bson:"value,omitempty" json:"value,omitempty"`
	Period   *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Assigner *Reference       `bson:"assigner,omitempty" json:"assigner,omitempty"`
}
