package fhir

// Resource ...
type Resource struct {
	ResourceType  string `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	ID            string `bson:"_id,omitempty" json:"id,omitempty"`
	Meta          *Meta  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      string `bson:"language,omitempty" json:"language,omitempty"`
}
