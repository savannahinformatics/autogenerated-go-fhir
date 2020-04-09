package fhir

// Dosage ...
type Dosage struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *int32                       `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Text                     string                       `bson:"text,omitempty" json:"text,omitempty"`
	AdditionalInstruction    []CodeableConcept            `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	PatientInstruction       string                       `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	Timing                   *Timing                      `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean          *bool                        `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept  *CodeableConcept             `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	Site                     *CodeableConcept             `bson:"site,omitempty" json:"site,omitempty"`
	Route                    *CodeableConcept             `bson:"route,omitempty" json:"route,omitempty"`
	Method                   *CodeableConcept             `bson:"method,omitempty" json:"method,omitempty"`
	DoseAndRate              []DosageDoseAndRateComponent `bson:"doseAndRate,omitempty" json:"doseAndRate,omitempty"`
	MaxDosePerPeriod         *Ratio                       `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration *Quantity                    `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *Quantity                    `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
}

// DosageDoseAndRateComponent ...
type DosageDoseAndRateComponent struct {
	BackboneElement    `bson:",inline"`
	Type               *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	DoseRange          *Range           `bson:"doseRange,omitempty" json:"doseRange,omitempty"`
	DoseSimpleQuantity *Quantity        `bson:"doseSimpleQuantity,omitempty" json:"doseSimpleQuantity,omitempty"`
	RateRatio          *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange          *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	RateSimpleQuantity *Quantity        `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
}
