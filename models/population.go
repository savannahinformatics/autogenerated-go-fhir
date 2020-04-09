package models

type Population struct {
	BackboneElement        `bson:",inline"`
	AgeRange               *Range           `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeCodeableConcept     *CodeableConcept `bson:"ageCodeableConcept,omitempty" json:"ageCodeableConcept,omitempty"`
	Gender                 *CodeableConcept `bson:"gender,omitempty" json:"gender,omitempty"`
	Race                   *CodeableConcept `bson:"race,omitempty" json:"race,omitempty"`
	PhysiologicalCondition *CodeableConcept `bson:"physiologicalCondition,omitempty" json:"physiologicalCondition,omitempty"`
}
