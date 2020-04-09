package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Task struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical *Canonical                `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       string                    `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier               `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	PartOf                []Reference               `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                    `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason          *CodeableConcept          `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	BusinessStatus        *CodeableConcept          `bson:"businessStatus,omitempty" json:"businessStatus,omitempty"`
	Intent                string                    `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority              string                    `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                  *CodeableConcept          `bson:"code,omitempty" json:"code,omitempty"`
	Description           string                    `bson:"description,omitempty" json:"description,omitempty"`
	Focus                 *Reference                `bson:"focus,omitempty" json:"focus,omitempty"`
	For                   *Reference                `bson:"for,omitempty" json:"for,omitempty"`
	Encounter             *Reference                `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ExecutionPeriod       *Period                   `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty"`
	AuthoredOn            *FHIRDateTime             `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	LastModified          *FHIRDateTime             `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Requester             *Reference                `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType         []CodeableConcept         `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Owner                 *Reference                `bson:"owner,omitempty" json:"owner,omitempty"`
	Location              *Reference                `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode            *CodeableConcept          `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       *Reference                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Insurance             []Reference               `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Note                  []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory       []Reference               `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Restriction           *TaskRestrictionComponent `bson:"restriction,omitempty" json:"restriction,omitempty"`
	Input                 []TaskParameterComponent  `bson:"input,omitempty" json:"input,omitempty"`
	Output                []TaskOutputComponent     `bson:"output,omitempty" json:"output,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Task) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Task"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "task" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type task Task

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Task) UnmarshalJSON(data []byte) (err error) {
	x2 := task{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Task(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Task) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Task"
	} else if x.ResourceType != "Task" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Task, instead received %s", x.ResourceType))
	}
	return nil
}

type TaskRestrictionComponent struct {
	BackboneElement `bson:",inline"`
	Repetitions     *uint32     `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	Period          *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Recipient       []Reference `bson:"recipient,omitempty" json:"recipient,omitempty"`
}

type TaskParameterComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskOutputComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueAddress         *Address         `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAnnotation      *Annotation      `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueBase64Binary    string           `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCode            string           `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCoding          *Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint    *ContactPoint    `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueDate            *FHIRDateTime    `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime        *FHIRDateTime    `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal         *float64         `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueHumanName       *HumanName       `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueId              string           `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueIdentifier      *Identifier      `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueInstant         *FHIRDateTime    `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger         *int32           `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown        string           `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueMeta            *Meta            `bson:"valueMeta,omitempty" json:"valueMeta,omitempty"`
	ValueOid             string           `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePeriod          *Period          `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValuePositiveInt     *uint32          `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio           `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData     *SampledData     `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature       *Signature       `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime            *FHIRDateTime    `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueTiming          *Timing          `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueUnsignedInt     *uint32          `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri             string           `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
}

type TaskPlus struct {
	Task                     `bson:",inline"`
	TaskPlusRelatedResources `bson:",inline"`
}

type TaskPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByOwner                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByOwner,omitempty"`
	IncludedOrganizationResourcesReferencedByOwner                         *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOwner,omitempty"`
	IncludedCareTeamResourcesReferencedByOwner                             *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByOwner,omitempty"`
	IncludedDeviceResourcesReferencedByOwner                               *[]Device                     `bson:"_includedDeviceResourcesReferencedByOwner,omitempty"`
	IncludedPatientResourcesReferencedByOwner                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByOwner,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByOwner                    *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByOwner,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByOwner                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByOwner,omitempty"`
	IncludedRelatedPersonResourcesReferencedByOwner                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByOwner,omitempty"`
	IncludedPractitionerResourcesReferencedByRequester                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByRequester                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByRequester,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRequester                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByRequester,omitempty"`
	IncludedTaskResourcesReferencedByPartof                                *[]Task                       `bson:"_includedTaskResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon                     *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingPartof                              *[]Task                       `bson:"_revIncludedTaskResourcesReferencingPartof,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference               *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingRequest            *[]PaymentReconciliation      `bson:"_revIncludedPaymentReconciliationResourcesReferencingRequest,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByOwner() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByOwner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByOwner))
	} else if len(*t.IncludedPractitionerResourcesReferencedByOwner) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOwner() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByOwner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByOwner))
	} else if len(*t.IncludedOrganizationResourcesReferencedByOwner) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedCareTeamResourceReferencedByOwner() (careTeam *CareTeam, err error) {
	if t.IncludedCareTeamResourcesReferencedByOwner == nil {
		err = errors.New("Included careteams not requested")
	} else if len(*t.IncludedCareTeamResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 careTeam, but found %d", len(*t.IncludedCareTeamResourcesReferencedByOwner))
	} else if len(*t.IncludedCareTeamResourcesReferencedByOwner) == 1 {
		careTeam = &(*t.IncludedCareTeamResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByOwner() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByOwner == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByOwner))
	} else if len(*t.IncludedDeviceResourcesReferencedByOwner) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByOwner() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByOwner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByOwner))
	} else if len(*t.IncludedPatientResourcesReferencedByOwner) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedHealthcareServiceResourceReferencedByOwner() (healthcareService *HealthcareService, err error) {
	if t.IncludedHealthcareServiceResourcesReferencedByOwner == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*t.IncludedHealthcareServiceResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*t.IncludedHealthcareServiceResourcesReferencedByOwner))
	} else if len(*t.IncludedHealthcareServiceResourcesReferencedByOwner) == 1 {
		healthcareService = &(*t.IncludedHealthcareServiceResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByOwner() (practitionerRole *PractitionerRole, err error) {
	if t.IncludedPractitionerRoleResourcesReferencedByOwner == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*t.IncludedPractitionerRoleResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*t.IncludedPractitionerRoleResourcesReferencedByOwner))
	} else if len(*t.IncludedPractitionerRoleResourcesReferencedByOwner) == 1 {
		practitionerRole = &(*t.IncludedPractitionerRoleResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByOwner() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByOwner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByOwner))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByOwner) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByOwner)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if t.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*t.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*t.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*t.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*t.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if t.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*t.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*t.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*t.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*t.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if t.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*t.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*t.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*t.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*t.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByRequester))
	} else if len(*t.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByRequester() (practitionerRole *PractitionerRole, err error) {
	if t.IncludedPractitionerRoleResourcesReferencedByRequester == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*t.IncludedPractitionerRoleResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*t.IncludedPractitionerRoleResourcesReferencedByRequester))
	} else if len(*t.IncludedPractitionerRoleResourcesReferencedByRequester) == 1 {
		practitionerRole = &(*t.IncludedPractitionerRoleResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRequester() (relatedPerson *RelatedPerson, err error) {
	if t.IncludedRelatedPersonResourcesReferencedByRequester == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*t.IncludedRelatedPersonResourcesReferencedByRequester))
	} else if len(*t.IncludedRelatedPersonResourcesReferencedByRequester) == 1 {
		relatedPerson = &(*t.IncludedRelatedPersonResourcesReferencedByRequester)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedTaskResourcesReferencedByPartof() (tasks []Task, err error) {
	if t.IncludedTaskResourcesReferencedByPartof == nil {
		err = errors.New("Included tasks not requested")
	} else {
		tasks = *t.IncludedTaskResourcesReferencedByPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if t.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*t.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*t.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*t.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*t.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if t.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*t.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*t.IncludedPatientResourcesReferencedByPatient))
	} else if len(*t.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*t.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *t.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if t.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *t.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if t.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *t.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingPartof() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingPartof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if t.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *t.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if t.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *t.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *t.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequest() (paymentReconciliations []PaymentReconciliation, err error) {
	if t.RevIncludedPaymentReconciliationResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *t.RevIncludedPaymentReconciliationResourcesReferencingRequest
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if t.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *t.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TaskPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TaskPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedCareTeamResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedCareTeamResourcesReferencedByOwner {
			rsc := (*t.IncludedCareTeamResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedHealthcareServiceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedHealthcareServiceResourcesReferencedByOwner {
			rsc := (*t.IncludedHealthcareServiceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerRoleResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerRoleResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerRoleResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*t.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*t.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByRequester {
			rsc := (*t.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerRoleResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerRoleResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerRoleResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByPartof != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByPartof {
			rsc := (*t.IncludedTaskResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *t.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*t.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByPatient {
			rsc := (*t.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*t.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingPartof {
			rsc := (*t.RevIncludedTaskResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *t.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*t.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentReconciliationResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentReconciliationResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentReconciliationResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TaskPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.IncludedPractitionerResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByOwner {
			rsc := (*t.IncludedOrganizationResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedCareTeamResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedCareTeamResourcesReferencedByOwner {
			rsc := (*t.IncludedCareTeamResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByOwner {
			rsc := (*t.IncludedDeviceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByOwner {
			rsc := (*t.IncludedPatientResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedHealthcareServiceResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedHealthcareServiceResourcesReferencedByOwner {
			rsc := (*t.IncludedHealthcareServiceResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerRoleResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedPractitionerRoleResourcesReferencedByOwner {
			rsc := (*t.IncludedPractitionerRoleResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByOwner != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByOwner {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*t.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*t.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByRequester {
			rsc := (*t.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPractitionerRoleResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedPractitionerRoleResourcesReferencedByRequester {
			rsc := (*t.IncludedPractitionerRoleResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *t.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*t.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedTaskResourcesReferencedByPartof != nil {
		for idx := range *t.IncludedTaskResourcesReferencedByPartof {
			rsc := (*t.IncludedTaskResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *t.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*t.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *t.IncludedPatientResourcesReferencedByPatient {
			rsc := (*t.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*t.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingPartof {
			rsc := (*t.RevIncludedTaskResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *t.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*t.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentReconciliationResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentReconciliationResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentReconciliationResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
