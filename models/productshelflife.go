package fhir

// ProductShelfLife ... // TODO Write proper comment
type ProductShelfLife struct {
	BackboneElement              `bson:",inline"`
	Identifier                   *Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                         *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Period                       *Quantity         `bson:"period,omitempty" json:"period,omitempty"`
	SpecialPrecautionsForStorage []CodeableConcept `bson:"specialPrecautionsForStorage,omitempty" json:"specialPrecautionsForStorage,omitempty"`
}
