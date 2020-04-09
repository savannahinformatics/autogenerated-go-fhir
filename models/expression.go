package models

type Expression struct {
	Element     `bson:",inline"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Language    string `bson:"language,omitempty" json:"language,omitempty"`
	Expression  string `bson:"expression,omitempty" json:"expression,omitempty"`
	Reference   string `bson:"reference,omitempty" json:"reference,omitempty"`
}
