package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// SpecimenDefinition ...
type SpecimenDefinition struct {
	DomainResource     `bson:",inline"`
	Identifier         *Identifier                             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	TypeCollected      *CodeableConcept                        `bson:"typeCollected,omitempty" json:"typeCollected,omitempty"`
	PatientPreparation []CodeableConcept                       `bson:"patientPreparation,omitempty" json:"patientPreparation,omitempty"`
	TimeAspect         string                                  `bson:"timeAspect,omitempty" json:"timeAspect,omitempty"`
	Collection         []CodeableConcept                       `bson:"collection,omitempty" json:"collection,omitempty"`
	TypeTested         []SpecimenDefinitionTypeTestedComponent `bson:"typeTested,omitempty" json:"typeTested,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *SpecimenDefinition) MarshalJSON() ([]byte, error) {
	x.ResourceType = "SpecimenDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "specimenDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type specimenDefinition SpecimenDefinition

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *SpecimenDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := specimenDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = SpecimenDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *SpecimenDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "SpecimenDefinition"
	} else if x.ResourceType != "SpecimenDefinition" {
		return fmt.Errorf("Expected resourceType to be SpecimenDefinition, instead received %s", x.ResourceType)
	}
	return nil
}

// SpecimenDefinitionTypeTestedComponent ...
type SpecimenDefinitionTypeTestedComponent struct {
	BackboneElement    `bson:",inline"`
	IsDerived          *bool                                           `bson:"isDerived,omitempty" json:"isDerived,omitempty"`
	Type               *CodeableConcept                                `bson:"type,omitempty" json:"type,omitempty"`
	Preference         string                                          `bson:"preference,omitempty" json:"preference,omitempty"`
	Container          *SpecimenDefinitionTypeTestedContainerComponent `bson:"container,omitempty" json:"container,omitempty"`
	Requirement        string                                          `bson:"requirement,omitempty" json:"requirement,omitempty"`
	RetentionTime      *Quantity                                       `bson:"retentionTime,omitempty" json:"retentionTime,omitempty"`
	RejectionCriterion []CodeableConcept                               `bson:"rejectionCriterion,omitempty" json:"rejectionCriterion,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandlingComponent `bson:"handling,omitempty" json:"handling,omitempty"`
}

// SpecimenDefinitionTypeTestedContainerComponent ...
type SpecimenDefinitionTypeTestedContainerComponent struct {
	BackboneElement             `bson:",inline"`
	Material                    *CodeableConcept                                         `bson:"material,omitempty" json:"material,omitempty"`
	Type                        *CodeableConcept                                         `bson:"type,omitempty" json:"type,omitempty"`
	Cap                         *CodeableConcept                                         `bson:"cap,omitempty" json:"cap,omitempty"`
	Description                 string                                                   `bson:"description,omitempty" json:"description,omitempty"`
	Capacity                    *Quantity                                                `bson:"capacity,omitempty" json:"capacity,omitempty"`
	MinimumVolumeSimpleQuantity *Quantity                                                `bson:"minimumVolumeSimpleQuantity,omitempty" json:"minimumVolumeSimpleQuantity,omitempty"`
	MinimumVolumeString         string                                                   `bson:"minimumVolumeString,omitempty" json:"minimumVolumeString,omitempty"`
	Additive                    []SpecimenDefinitionTypeTestedContainerAdditiveComponent `bson:"additive,omitempty" json:"additive,omitempty"`
	Preparation                 string                                                   `bson:"preparation,omitempty" json:"preparation,omitempty"`
}

// SpecimenDefinitionTypeTestedContainerAdditiveComponent ...
type SpecimenDefinitionTypeTestedContainerAdditiveComponent struct {
	BackboneElement         `bson:",inline"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

// SpecimenDefinitionTypeTestedHandlingComponent ...
type SpecimenDefinitionTypeTestedHandlingComponent struct {
	BackboneElement      `bson:",inline"`
	TemperatureQualifier *CodeableConcept `bson:"temperatureQualifier,omitempty" json:"temperatureQualifier,omitempty"`
	TemperatureRange     *Range           `bson:"temperatureRange,omitempty" json:"temperatureRange,omitempty"`
	MaxDuration          *Quantity        `bson:"maxDuration,omitempty" json:"maxDuration,omitempty"`
	Instruction          string           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

// SpecimenDefinitionPlus ...
type SpecimenDefinitionPlus struct {
	SpecimenDefinition                     `bson:",inline"`
	SpecimenDefinitionPlusRelatedResources `bson:",inline"`
}

// SpecimenDefinitionPlusRelatedResources ...
type SpecimenDefinitionPlusRelatedResources struct {
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
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *s.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if s.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *s.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if s.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *s.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *s.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if s.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *s.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (s *SpecimenDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ...
func (s *SpecimenDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (s *SpecimenDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
