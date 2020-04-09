package fhir

// Element ...
type Element struct {
	ID        string      `bson:"_id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}
