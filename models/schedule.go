package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Schedule struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active          *bool             `bson:"active,omitempty" json:"active,omitempty"`
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType     []CodeableConcept `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty       []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Actor           []Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	PlanningHorizon *Period           `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty"`
	Comment         string            `bson:"comment,omitempty" json:"comment,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Schedule) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Schedule"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "schedule" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type schedule Schedule

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Schedule) UnmarshalJSON(data []byte) (err error) {
	x2 := schedule{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Schedule(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Schedule) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Schedule"
	} else if x.ResourceType != "Schedule" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Schedule, instead received %s", x.ResourceType))
	}
	return nil
}

type SchedulePlus struct {
	Schedule                     `bson:",inline"`
	SchedulePlusRelatedResources `bson:",inline"`
}

type SchedulePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByActor                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByActor,omitempty"`
	IncludedDeviceResourcesReferencedByActor                               *[]Device                     `bson:"_includedDeviceResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByActor                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByActor,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByActor                    *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByActor,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByActor                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByActor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByActor                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByActor,omitempty"`
	IncludedLocationResourcesReferencedByActor                             *[]Location                   `bson:"_includedLocationResourcesReferencedByActor,omitempty"`
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
	RevIncludedSlotResourcesReferencingSchedule                            *[]Slot                       `bson:"_revIncludedSlotResourcesReferencingSchedule,omitempty"`
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
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (s *SchedulePlusRelatedResources) GetIncludedPractitionerResourcesReferencedByActor() (practitioners []Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedByActor == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *s.IncludedPractitionerResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedDeviceResourcesReferencedByActor() (devices []Device, err error) {
	if s.IncludedDeviceResourcesReferencedByActor == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *s.IncludedDeviceResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedPatientResourcesReferencedByActor() (patients []Patient, err error) {
	if s.IncludedPatientResourcesReferencedByActor == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *s.IncludedPatientResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByActor() (healthcareServices []HealthcareService, err error) {
	if s.IncludedHealthcareServiceResourcesReferencedByActor == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *s.IncludedHealthcareServiceResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByActor() (practitionerRoles []PractitionerRole, err error) {
	if s.IncludedPractitionerRoleResourcesReferencedByActor == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *s.IncludedPractitionerRoleResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByActor() (relatedPeople []RelatedPerson, err error) {
	if s.IncludedRelatedPersonResourcesReferencedByActor == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *s.IncludedRelatedPersonResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedLocationResourcesReferencedByActor() (locations []Location, err error) {
	if s.IncludedLocationResourcesReferencedByActor == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *s.IncludedLocationResourcesReferencedByActor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *s.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if s.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *s.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedSlotResourcesReferencingSchedule() (slots []Slot, err error) {
	if s.RevIncludedSlotResourcesReferencingSchedule == nil {
		err = errors.New("RevIncluded slots not requested")
	} else {
		slots = *s.RevIncludedSlotResourcesReferencingSchedule
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if s.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *s.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *s.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if s.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *s.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SchedulePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SchedulePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*s.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByActor {
			rsc := (*s.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByActor {
			rsc := (*s.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *s.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*s.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByActor {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *s.IncludedLocationResourcesReferencedByActor {
			rsc := (*s.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SchedulePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSlotResourcesReferencingSchedule != nil {
		for idx := range *s.RevIncludedSlotResourcesReferencingSchedule {
			rsc := (*s.RevIncludedSlotResourcesReferencingSchedule)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SchedulePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*s.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByActor {
			rsc := (*s.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByActor {
			rsc := (*s.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *s.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*s.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByActor != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByActor {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *s.IncludedLocationResourcesReferencedByActor {
			rsc := (*s.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSlotResourcesReferencingSchedule != nil {
		for idx := range *s.RevIncludedSlotResourcesReferencingSchedule {
			rsc := (*s.RevIncludedSlotResourcesReferencingSchedule)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
