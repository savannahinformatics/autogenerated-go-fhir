package models

type UsageContext struct {
	Code                 *Coding          `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
}
