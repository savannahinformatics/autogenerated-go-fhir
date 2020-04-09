package fhir

// TriggerDefinition ... // TODO Write proper comment
type TriggerDefinition struct {
	Type            string            `bson:"type,omitempty" json:"type,omitempty"`
	Name            string            `bson:"name,omitempty" json:"name,omitempty"`
	TimingTiming    *Timing           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingReference *Reference        `bson:"timingReference,omitempty" json:"timingReference,omitempty"`
	TimingDate      *FHIRDateTime     `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingDateTime  *FHIRDateTime     `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	Data            []DataRequirement `bson:"data,omitempty" json:"data,omitempty"`
	Condition       *Expression       `bson:"condition,omitempty" json:"condition,omitempty"`
}
