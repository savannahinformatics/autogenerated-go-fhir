package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FamilyMemberHistory ... // TODO Write proper comment
type FamilyMemberHistory struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []Canonical                             `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                                `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Status                string                                  `bson:"status,omitempty" json:"status,omitempty"`
	DataAbsentReason      *CodeableConcept                        `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Patient               *Reference                              `bson:"patient,omitempty" json:"patient,omitempty"`
	Date                  *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Name                  string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Relationship          *CodeableConcept                        `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Sex                   *CodeableConcept                        `bson:"sex,omitempty" json:"sex,omitempty"`
	BornPeriod            *Period                                 `bson:"bornPeriod,omitempty" json:"bornPeriod,omitempty"`
	BornDate              *FHIRDateTime                           `bson:"bornDate,omitempty" json:"bornDate,omitempty"`
	BornString            string                                  `bson:"bornString,omitempty" json:"bornString,omitempty"`
	AgeAge                *Quantity                               `bson:"ageAge,omitempty" json:"ageAge,omitempty"`
	AgeRange              *Range                                  `bson:"ageRange,omitempty" json:"ageRange,omitempty"`
	AgeString             string                                  `bson:"ageString,omitempty" json:"ageString,omitempty"`
	EstimatedAge          *bool                                   `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                                   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Quantity                               `bson:"deceasedAge,omitempty" json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                                  `bson:"deceasedRange,omitempty" json:"deceasedRange,omitempty"`
	DeceasedDate          *FHIRDateTime                           `bson:"deceasedDate,omitempty" json:"deceasedDate,omitempty"`
	DeceasedString        string                                  `bson:"deceasedString,omitempty" json:"deceasedString,omitempty"`
	ReasonCode            []CodeableConcept                       `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference                             `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation                            `bson:"note,omitempty" json:"note,omitempty"`
	Condition             []FamilyMemberHistoryConditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	x.ResourceType = "FamilyMemberHistory"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "familyMemberHistory" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type familyMemberHistory FamilyMemberHistory

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *FamilyMemberHistory) UnmarshalJSON(data []byte) (err error) {
	x2 := familyMemberHistory{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = FamilyMemberHistory(x2)
		return x.checkResourceType()
	}
	return
}

func (x *FamilyMemberHistory) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "FamilyMemberHistory"
	} else if x.ResourceType != "FamilyMemberHistory" {
		return fmt.Errorf("Expected resourceType to be FamilyMemberHistory, instead received %s", x.ResourceType)
	}
	return nil
}

// FamilyMemberHistoryConditionComponent ... // TODO Write proper comment
type FamilyMemberHistoryConditionComponent struct {
	BackboneElement    `bson:",inline"`
	Code               *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Outcome            *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ContributedToDeath *bool            `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty"`
	OnsetAge           *Quantity        `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetRange         *Range           `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetPeriod        *Period          `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetString        string           `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	Note               []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
}

// FamilyMemberHistoryPlus ... // TODO Write proper comment
type FamilyMemberHistoryPlus struct {
	FamilyMemberHistory                     `bson:",inline"`
	FamilyMemberHistoryPlusRelatedResources `bson:",inline"`
}

// FamilyMemberHistoryPlusRelatedResources ... // TODO Write proper comment
type FamilyMemberHistoryPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedQuestionnaireResourcesReferencedByInstantiatescanonical        *[]Questionnaire              `bson:"_includedQuestionnaireResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedMeasureResourcesReferencedByInstantiatescanonical              *[]Measure                    `bson:"_includedMeasureResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical       *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical  *[]OperationDefinition        `bson:"_includedOperationDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical   *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
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
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
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
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

// GetIncludedPatientResourceReferencedByPatient ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByPatient))
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical() (questionnaires []Questionnaire, err error) {
	if f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included questionnaires not requested")
	} else {
		questionnaires = *f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedMeasureResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedMeasureResourcesReferencedByInstantiatescanonical() (measures []Measure, err error) {
	if f.IncludedMeasureResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included measures not requested")
	} else {
		measures = *f.IncludedMeasureResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical() (operationDefinitions []OperationDefinition, err error) {
	if f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included operationDefinitions not requested")
	} else {
		operationDefinitions = *f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *f.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if f.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *f.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *f.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if f.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *f.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if f.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *f.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if f.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *f.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if f.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *f.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if f.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *f.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if f.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *f.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if f.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *f.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if f.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *f.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *f.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if f.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *f.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if f.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *f.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if f.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *f.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *f.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingInvestigation ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingData {
			rsc := (*f.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*f.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*f.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*f.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*f.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*f.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ... // TODO Write proper comment
func (f *FamilyMemberHistoryPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*f.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingData {
			rsc := (*f.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*f.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*f.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*f.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*f.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*f.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
