package fhir

// Meta ... // TODO Write proper comment
type Meta struct {
	Element     `bson:",inline"`
	VersionId   string        `bson:"versionId,omitempty" json:"versionId,omitempty"`
	LastUpdated *FHIRDateTime `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Source      string        `bson:"source,omitempty" json:"source,omitempty"`
	Profile     []Canonical   `bson:"profile,omitempty" json:"profile,omitempty"`
	Security    []Coding      `bson:"security,omitempty" json:"security,omitempty"`
	Tag         []Coding      `bson:"tag,omitempty" json:"tag,omitempty"`
}
