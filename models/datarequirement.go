package fhir

// DataRequirement ... // TODO Write proper comment
type DataRequirement struct {
	Type                   string                               `bson:"type,omitempty" json:"type,omitempty"`
	Profile                []Canonical                          `bson:"profile,omitempty" json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept                     `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                           `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	MustSupport            []string                             `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilterComponent `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilterComponent `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
	Limit                  *uint32                              `bson:"limit,omitempty" json:"limit,omitempty"`
	Sort                   []DataRequirementSortComponent       `bson:"sort,omitempty" json:"sort,omitempty"`
}

// DataRequirementCodeFilterComponent ... // TODO Write proper comment
type DataRequirementCodeFilterComponent struct {
	BackboneElement `bson:",inline"`
	Path            string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam     string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueSet        *Canonical `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Code            []Coding   `bson:"code,omitempty" json:"code,omitempty"`
}

// DataRequirementDateFilterComponent ... // TODO Write proper comment
type DataRequirementDateFilterComponent struct {
	BackboneElement `bson:",inline"`
	Path            string        `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam     string        `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod     *Period       `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration   *Quantity     `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}

// DataRequirementSortComponent ... // TODO Write proper comment
type DataRequirementSortComponent struct {
	BackboneElement `bson:",inline"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Direction       string `bson:"direction,omitempty" json:"direction,omitempty"`
}
