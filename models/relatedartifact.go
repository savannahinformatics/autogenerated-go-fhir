package models

type RelatedArtifact struct {
	Type     string      `bson:"type,omitempty" json:"type,omitempty"`
	Label    string      `bson:"label,omitempty" json:"label,omitempty"`
	Display  string      `bson:"display,omitempty" json:"display,omitempty"`
	Citation string      `bson:"citation,omitempty" json:"citation,omitempty"`
	Url      *URL        `bson:"url,omitempty" json:"url,omitempty"`
	Document *Attachment `bson:"document,omitempty" json:"document,omitempty"`
	Resource *Canonical  `bson:"resource,omitempty" json:"resource,omitempty"`
}
