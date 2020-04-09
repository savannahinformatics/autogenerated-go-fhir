package models

type ProdCharacteristic struct {
	BackboneElement  `bson:",inline"`
	Height           *Quantity        `bson:"height,omitempty" json:"height,omitempty"`
	Width            *Quantity        `bson:"width,omitempty" json:"width,omitempty"`
	Depth            *Quantity        `bson:"depth,omitempty" json:"depth,omitempty"`
	Weight           *Quantity        `bson:"weight,omitempty" json:"weight,omitempty"`
	NominalVolume    *Quantity        `bson:"nominalVolume,omitempty" json:"nominalVolume,omitempty"`
	ExternalDiameter *Quantity        `bson:"externalDiameter,omitempty" json:"externalDiameter,omitempty"`
	Shape            string           `bson:"shape,omitempty" json:"shape,omitempty"`
	Color            []string         `bson:"color,omitempty" json:"color,omitempty"`
	Imprint          []string         `bson:"imprint,omitempty" json:"imprint,omitempty"`
	Image            []Attachment     `bson:"image,omitempty" json:"image,omitempty"`
	Scoring          *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty"`
}
