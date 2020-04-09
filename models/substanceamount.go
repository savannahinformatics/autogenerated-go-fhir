package models

type SubstanceAmount struct {
	BackboneElement `bson:",inline"`
	AmountQuantity  *Quantity                               `bson:"amountQuantity,omitempty" json:"amountQuantity,omitempty"`
	AmountRange     *Range                                  `bson:"amountRange,omitempty" json:"amountRange,omitempty"`
	AmountString    string                                  `bson:"amountString,omitempty" json:"amountString,omitempty"`
	AmountType      *CodeableConcept                        `bson:"amountType,omitempty" json:"amountType,omitempty"`
	AmountText      string                                  `bson:"amountText,omitempty" json:"amountText,omitempty"`
	ReferenceRange  *SubstanceAmountReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

type SubstanceAmountReferenceRangeComponent struct {
	BackboneElement `bson:",inline"`
	LowLimit        *Quantity `bson:"lowLimit,omitempty" json:"lowLimit,omitempty"`
	HighLimit       *Quantity `bson:"highLimit,omitempty" json:"highLimit,omitempty"`
}
