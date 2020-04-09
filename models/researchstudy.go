package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ResearchStudy struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title                 string                            `bson:"title,omitempty" json:"title,omitempty"`
	Protocol              []Reference                       `bson:"protocol,omitempty" json:"protocol,omitempty"`
	PartOf                []Reference                       `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                            `bson:"status,omitempty" json:"status,omitempty"`
	PrimaryPurposeType    *CodeableConcept                  `bson:"primaryPurposeType,omitempty" json:"primaryPurposeType,omitempty"`
	Phase                 *CodeableConcept                  `bson:"phase,omitempty" json:"phase,omitempty"`
	Category              []CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Focus                 []CodeableConcept                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Condition             []CodeableConcept                 `bson:"condition,omitempty" json:"condition,omitempty"`
	Contact               []ContactDetail                   `bson:"contact,omitempty" json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact                 `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept                 `bson:"keyword,omitempty" json:"keyword,omitempty"`
	Location              []CodeableConcept                 `bson:"location,omitempty" json:"location,omitempty"`
	Description           string                            `bson:"description,omitempty" json:"description,omitempty"`
	Enrollment            []Reference                       `bson:"enrollment,omitempty" json:"enrollment,omitempty"`
	Period                *Period                           `bson:"period,omitempty" json:"period,omitempty"`
	Sponsor               *Reference                        `bson:"sponsor,omitempty" json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference                        `bson:"principalInvestigator,omitempty" json:"principalInvestigator,omitempty"`
	Site                  []Reference                       `bson:"site,omitempty" json:"site,omitempty"`
	ReasonStopped         *CodeableConcept                  `bson:"reasonStopped,omitempty" json:"reasonStopped,omitempty"`
	Note                  []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	Arm                   []ResearchStudyArmComponent       `bson:"arm,omitempty" json:"arm,omitempty"`
	Objective             []ResearchStudyObjectiveComponent `bson:"objective,omitempty" json:"objective,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *ResearchStudy) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ResearchStudy"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "researchStudy" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type researchStudy ResearchStudy

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ResearchStudy) UnmarshalJSON(data []byte) (err error) {
	x2 := researchStudy{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ResearchStudy(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ResearchStudy) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ResearchStudy"
	} else if x.ResourceType != "ResearchStudy" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ResearchStudy, instead received %s", x.ResourceType))
	}
	return nil
}

type ResearchStudyArmComponent struct {
	BackboneElement `bson:",inline"`
	Name            string           `bson:"name,omitempty" json:"name,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Description     string           `bson:"description,omitempty" json:"description,omitempty"`
}

type ResearchStudyObjectiveComponent struct {
	BackboneElement `bson:",inline"`
	Name            string           `bson:"name,omitempty" json:"name,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

type ResearchStudyPlus struct {
	ResearchStudy                     `bson:",inline"`
	ResearchStudyPlusRelatedResources `bson:",inline"`
}

type ResearchStudyPlusRelatedResources struct {
	IncludedResearchStudyResourcesReferencedByPartof                       *[]ResearchStudy              `bson:"_includedResearchStudyResourcesReferencedByPartof,omitempty"`
	IncludedOrganizationResourcesReferencedBySponsor                       *[]Organization               `bson:"_includedOrganizationResourcesReferencedBySponsor,omitempty"`
	IncludedPractitionerResourcesReferencedByPrincipalinvestigator         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPrincipalinvestigator,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPrincipalinvestigator,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByProtocol                    *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByProtocol,omitempty"`
	IncludedLocationResourcesReferencedBySite                              *[]Location                   `bson:"_includedLocationResourcesReferencedBySite,omitempty"`
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
	RevIncludedResearchSubjectResourcesReferencingStudy                    *[]ResearchSubject            `bson:"_revIncludedResearchSubjectResourcesReferencingStudy,omitempty"`
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
	RevIncludedResearchStudyResourcesReferencingPartof                     *[]ResearchStudy              `bson:"_revIncludedResearchStudyResourcesReferencingPartof,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedAdverseEventResourcesReferencingStudy                       *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingStudy,omitempty"`
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

func (r *ResearchStudyPlusRelatedResources) GetIncludedResearchStudyResourcesReferencedByPartof() (researchStudies []ResearchStudy, err error) {
	if r.IncludedResearchStudyResourcesReferencedByPartof == nil {
		err = errors.New("Included researchStudies not requested")
	} else {
		researchStudies = *r.IncludedResearchStudyResourcesReferencedByPartof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySponsor() (organization *Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedBySponsor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*r.IncludedOrganizationResourcesReferencedBySponsor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*r.IncludedOrganizationResourcesReferencedBySponsor))
	} else if len(*r.IncludedOrganizationResourcesReferencedBySponsor) == 1 {
		organization = &(*r.IncludedOrganizationResourcesReferencedBySponsor)[0]
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPrincipalinvestigator() (practitioner *Practitioner, err error) {
	if r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator))
	} else if len(*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator) == 1 {
		practitioner = &(*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator)[0]
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPrincipalinvestigator() (practitionerRole *PractitionerRole, err error) {
	if r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator))
	} else if len(*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator) == 1 {
		practitionerRole = &(*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator)[0]
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByProtocol() (planDefinitions []PlanDefinition, err error) {
	if r.IncludedPlanDefinitionResourcesReferencedByProtocol == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *r.IncludedPlanDefinitionResourcesReferencedByProtocol
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedLocationResourcesReferencedBySite() (locations []Location, err error) {
	if r.IncludedLocationResourcesReferencedBySite == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *r.IncludedLocationResourcesReferencedBySite
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchSubjectResourcesReferencingStudy() (researchSubjects []ResearchSubject, err error) {
	if r.RevIncludedResearchSubjectResourcesReferencingStudy == nil {
		err = errors.New("RevIncluded researchSubjects not requested")
	} else {
		researchSubjects = *r.RevIncludedResearchSubjectResourcesReferencingStudy
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if r.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *r.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingPartof() (researchStudies []ResearchStudy, err error) {
	if r.RevIncludedResearchStudyResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *r.RevIncludedResearchStudyResourcesReferencingPartof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingStudy() (adverseEvents []AdverseEvent, err error) {
	if r.RevIncludedAdverseEventResourcesReferencingStudy == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *r.RevIncludedAdverseEventResourcesReferencingStudy
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if r.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *r.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedResearchStudyResourcesReferencedByPartof != nil {
		for idx := range *r.IncludedResearchStudyResourcesReferencedByPartof {
			rsc := (*r.IncludedResearchStudyResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedBySponsor != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedBySponsor {
			rsc := (*r.IncludedOrganizationResourcesReferencedBySponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator {
			rsc := (*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator != nil {
		for idx := range *r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator {
			rsc := (*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPlanDefinitionResourcesReferencedByProtocol != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedByProtocol {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedByProtocol)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedLocationResourcesReferencedBySite != nil {
		for idx := range *r.IncludedLocationResourcesReferencedBySite {
			rsc := (*r.IncludedLocationResourcesReferencedBySite)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ResearchStudyPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchSubjectResourcesReferencingStudy != nil {
		for idx := range *r.RevIncludedResearchSubjectResourcesReferencingStudy {
			rsc := (*r.RevIncludedResearchSubjectResourcesReferencingStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchStudyResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedResearchStudyResourcesReferencingPartof {
			rsc := (*r.RevIncludedResearchStudyResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingStudy != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingStudy {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *ResearchStudyPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedResearchStudyResourcesReferencedByPartof != nil {
		for idx := range *r.IncludedResearchStudyResourcesReferencedByPartof {
			rsc := (*r.IncludedResearchStudyResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedBySponsor != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedBySponsor {
			rsc := (*r.IncludedOrganizationResourcesReferencedBySponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator != nil {
		for idx := range *r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator {
			rsc := (*r.IncludedPractitionerResourcesReferencedByPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator != nil {
		for idx := range *r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator {
			rsc := (*r.IncludedPractitionerRoleResourcesReferencedByPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPlanDefinitionResourcesReferencedByProtocol != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedByProtocol {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedByProtocol)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedLocationResourcesReferencedBySite != nil {
		for idx := range *r.IncludedLocationResourcesReferencedBySite {
			rsc := (*r.IncludedLocationResourcesReferencedBySite)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchSubjectResourcesReferencingStudy != nil {
		for idx := range *r.RevIncludedResearchSubjectResourcesReferencingStudy {
			rsc := (*r.RevIncludedResearchSubjectResourcesReferencingStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchStudyResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedResearchStudyResourcesReferencingPartof {
			rsc := (*r.RevIncludedResearchStudyResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingStudy != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingStudy {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingStudy)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
