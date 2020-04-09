package fhir

// MarketingStatus ... // TODO Write proper comment
type MarketingStatus struct {
	BackboneElement `bson:",inline"`
	Country         *CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
	Jurisdiction    *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Status          *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	DateRange       *Period          `bson:"dateRange,omitempty" json:"dateRange,omitempty"`
	RestoreDate     *FHIRDateTime    `bson:"restoreDate,omitempty" json:"restoreDate,omitempty"`
}
