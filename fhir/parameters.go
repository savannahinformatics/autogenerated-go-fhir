package fhir

import (
	"encoding/json"
	"fmt"
)

// Parameters ...
type Parameters struct {
	Resource  `bson:",inline"`
	Parameter []ParametersParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Parameters) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Parameters"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "parameters" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type parameters Parameters

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Parameters) UnmarshalJSON(data []byte) (err error) {
	x2 := parameters{}
	if err = json.Unmarshal(data, &x2); err == nil {
		*x = Parameters(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Parameters) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Parameters"
	} else if x.ResourceType != "Parameters" {
		return fmt.Errorf("Expected resourceType to be Parameters, instead received %s", x.ResourceType)
	}
	return nil
}

// ParametersParameterComponent ...
type ParametersParameterComponent struct {
	BackboneElement      `bson:",inline"`
	Name                 string                         `bson:"name,omitempty" json:"name,omitempty"`
	ValueAddress         *Address                       `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation                    `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment                    `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string                         `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool                          `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string                         `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept               `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding                        `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint                  `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime                  `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime                  `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64                       `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName                     `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueID              string                         `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier                    `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime                  `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32                         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string                         `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta                          `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOID             string                         `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period                        `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32                        `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity                      `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range                         `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                         `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference                     `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData                   `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature                     `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string                         `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime                  `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing                        `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32                        `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueURI             string                         `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	Resource             interface{}                    `bson:"resource,omitempty" json:"resource,omitempty"`
	Part                 []ParametersParameterComponent `bson:"part,omitempty" json:"part,omitempty"`
}

// "parametersParameterComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type parametersParameterComponent ParametersParameterComponent

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ParametersParameterComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := parametersParameterComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Resource != nil {
			x2.Resource = MapToResource(x2.Resource, true)
		}
		*x = ParametersParameterComponent(x2)
	}
	return
}
