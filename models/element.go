package fhir

// Element ... // TODO Write proper comment
type Element struct {
	Id        string      `bson:"_id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}
