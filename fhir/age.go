package fhir

// Age is a FHIR age structure
type Age struct {
	Quantity `bson:",inline"`
}
