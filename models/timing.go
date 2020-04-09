package fhir

// Timing ...
type Timing struct {
	BackboneElement `bson:",inline"`
	Event           []FHIRDateTime         `bson:"event,omitempty" json:"event,omitempty"`
	Repeat          *TimingRepeatComponent `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Code            *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
}

// TimingRepeatComponent ...
type TimingRepeatComponent struct {
	BackboneElement `bson:",inline"`
	BoundsDuration  *Quantity      `bson:"boundsDuration,omitempty" json:"boundsDuration,omitempty"`
	BoundsRange     *Range         `bson:"boundsRange,omitempty" json:"boundsRange,omitempty"`
	BoundsPeriod    *Period        `bson:"boundsPeriod,omitempty" json:"boundsPeriod,omitempty"`
	Count           *uint32        `bson:"count,omitempty" json:"count,omitempty"`
	CountMax        *uint32        `bson:"countMax,omitempty" json:"countMax,omitempty"`
	Duration        *float64       `bson:"duration,omitempty" json:"duration,omitempty"`
	DurationMax     *float64       `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	DurationUnit    string         `bson:"durationUnit,omitempty" json:"durationUnit,omitempty"`
	Frequency       *uint32        `bson:"frequency,omitempty" json:"frequency,omitempty"`
	FrequencyMax    *uint32        `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	Period          *float64       `bson:"period,omitempty" json:"period,omitempty"`
	PeriodMax       *float64       `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	PeriodUnit      string         `bson:"periodUnit,omitempty" json:"periodUnit,omitempty"`
	DayOfWeek       []string       `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	TimeOfDay       []FHIRDateTime `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty"`
	When            []string       `bson:"when,omitempty" json:"when,omitempty"`
	Offset          *uint32        `bson:"offset,omitempty" json:"offset,omitempty"`
}
