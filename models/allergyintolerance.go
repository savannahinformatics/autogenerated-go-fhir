package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AllergyIntolerance struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept                      `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept                      `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               string                                `bson:"type,omitempty" json:"type,omitempty"`
	Category           []string                              `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        string                                `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept                      `bson:"code,omitempty" json:"code,omitempty"`
	Patient            *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter          *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OnsetDateTime      *FHIRDateTime                         `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Quantity                             `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                               `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                                `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        string                                `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	RecordedDate       *FHIRDateTime                         `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder           *Reference                            `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference                            `bson:"asserter,omitempty" json:"asserter,omitempty"`
	LastOccurrence     *FHIRDateTime                         `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                          `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReactionComponent `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *AllergyIntolerance) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AllergyIntolerance"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "allergyIntolerance" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type allergyIntolerance AllergyIntolerance

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *AllergyIntolerance) UnmarshalJSON(data []byte) (err error) {
	x2 := allergyIntolerance{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AllergyIntolerance(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AllergyIntolerance) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AllergyIntolerance"
	} else if x.ResourceType != "AllergyIntolerance" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AllergyIntolerance, instead received %s", x.ResourceType))
	}
	return nil
}

type AllergyIntoleranceReactionComponent struct {
	BackboneElement `bson:",inline"`
	Substance       *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Manifestation   []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description     string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset           *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity        string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute   *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note            []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

type AllergyIntolerancePlus struct {
	AllergyIntolerance                     `bson:",inline"`
	AllergyIntolerancePlusRelatedResources `bson:",inline"`
}

type AllergyIntolerancePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRecorder                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRecorder,omitempty"`
	IncludedPatientResourcesReferencedByRecorder                           *[]Patient                    `bson:"_includedPatientResourcesReferencedByRecorder,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByRecorder                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByRecorder,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRecorder                     *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByRecorder,omitempty"`
	IncludedPractitionerResourcesReferencedByAsserter                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAsserter,omitempty"`
	IncludedPatientResourcesReferencedByAsserter                           *[]Patient                    `bson:"_includedPatientResourcesReferencedByAsserter,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByAsserter                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByAsserter,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAsserter                     *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByAsserter,omitempty"`
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
	RevIncludedClinicalImpressionResourcesReferencingProblem               *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingProblem,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByRecorder() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByRecorder == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByRecorder))
	} else if len(*a.IncludedPractitionerResourcesReferencedByRecorder) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByRecorder() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByRecorder == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByRecorder))
	} else if len(*a.IncludedPatientResourcesReferencedByRecorder) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByRecorder() (practitionerRole *PractitionerRole, err error) {
	if a.IncludedPractitionerRoleResourcesReferencedByRecorder == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*a.IncludedPractitionerRoleResourcesReferencedByRecorder))
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByRecorder) == 1 {
		practitionerRole = &(*a.IncludedPractitionerRoleResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRecorder() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByRecorder == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByRecorder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByRecorder))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByRecorder) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByRecorder)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAsserter() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByAsserter == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByAsserter))
	} else if len(*a.IncludedPractitionerResourcesReferencedByAsserter) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByAsserter() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByAsserter == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByAsserter))
	} else if len(*a.IncludedPatientResourcesReferencedByAsserter) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByAsserter() (practitionerRole *PractitionerRole, err error) {
	if a.IncludedPractitionerRoleResourcesReferencedByAsserter == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*a.IncludedPractitionerRoleResourcesReferencedByAsserter))
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByAsserter) == 1 {
		practitionerRole = &(*a.IncludedPractitionerRoleResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAsserter() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByAsserter == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAsserter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByAsserter))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAsserter) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByAsserter)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *a.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if a.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *a.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if a.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *a.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *a.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if a.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *a.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingProblem() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingProblem
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByRecorder {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAsserter {
			rsc := (*a.IncludedPatientResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAsserter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AllergyIntolerancePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByRecorder {
			rsc := (*a.IncludedPatientResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByRecorder {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByRecorder != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByRecorder {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAsserter {
			rsc := (*a.IncludedPatientResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByAsserter {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAsserter != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAsserter {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingProblem != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingProblem {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingProblem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
