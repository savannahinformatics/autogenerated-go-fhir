package fhir

// BackboneElement ... // TODO Write proper comment
type BackboneElement struct {
	Element           `bson:",inline"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
