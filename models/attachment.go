package models

// Attachment is a FHIR attachment resource
type Attachment struct {
	ContentType string        `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Language    string        `bson:"language,omitempty" json:"language,omitempty"`
	Data        string        `bson:"data,omitempty" json:"data,omitempty"`
	Url         *URL          `bson:"url,omitempty" json:"url,omitempty"`
	Size        *uint32       `bson:"size,omitempty" json:"size,omitempty"`
	Hash        string        `bson:"hash,omitempty" json:"hash,omitempty"`
	Title       string        `bson:"title,omitempty" json:"title,omitempty"`
	Creation    *FHIRDateTime `bson:"creation,omitempty" json:"creation,omitempty"`
}
