package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NamingSystem ...
type NamingSystem struct {
	DomainResource `bson:",inline"`
	Name           string                          `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                          `bson:"status,omitempty" json:"status,omitempty"`
	Kind           string                          `bson:"kind,omitempty" json:"kind,omitempty"`
	Date           *FHIRDateTime                   `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible    string                          `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Type           *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Description    string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Usage          string                          `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId       []NamingSystemUniqueIdComponent `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *NamingSystem) MarshalJSON() ([]byte, error) {
	x.ResourceType = "NamingSystem"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "namingSystem" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type namingSystem NamingSystem

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *NamingSystem) UnmarshalJSON(data []byte) (err error) {
	x2 := namingSystem{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NamingSystem(x2)
		return x.checkResourceType()
	}
	return
}

func (x *NamingSystem) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "NamingSystem"
	} else if x.ResourceType != "NamingSystem" {
		return fmt.Errorf("Expected resourceType to be NamingSystem, instead received %s", x.ResourceType)
	}
	return nil
}

// NamingSystemUniqueIdComponent ...
type NamingSystemUniqueIdComponent struct {
	BackboneElement `bson:",inline"`
	Type            string  `bson:"type,omitempty" json:"type,omitempty"`
	Value           string  `bson:"value,omitempty" json:"value,omitempty"`
	Preferred       *bool   `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Comment         string  `bson:"comment,omitempty" json:"comment,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

// NamingSystemPlus ...
type NamingSystemPlus struct {
	NamingSystem                     `bson:",inline"`
	NamingSystemPlusRelatedResources `bson:",inline"`
}

// NamingSystemPlusRelatedResources ...
type NamingSystemPlusRelatedResources struct {
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
func (n *NamingSystemPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *n.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if n.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *n.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *n.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if n.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *n.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if n.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *n.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if n.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *n.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if n.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *n.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if n.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *n.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if n.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *n.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *n.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if n.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *n.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if n.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *n.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if n.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *n.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *n.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (n *NamingSystemPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ...
func (n *NamingSystemPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (n *NamingSystemPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
