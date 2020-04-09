package models

type HumanName struct {
	Use    string   `bson:"use,omitempty" json:"use,omitempty"`
	Text   string   `bson:"text,omitempty" json:"text,omitempty"`
	Family string   `bson:"family,omitempty" json:"family,omitempty"`
	Given  []string `bson:"given,omitempty" json:"given,omitempty"`
	Prefix []string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Suffix []string `bson:"suffix,omitempty" json:"suffix,omitempty"`
	Period *Period  `bson:"period,omitempty" json:"period,omitempty"`
}
