package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// TerminologyCapabilities ...
type TerminologyCapabilities struct {
	DomainResource `bson:",inline"`
	URL            string                                          `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                          `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                          `bson:"name,omitempty" json:"name,omitempty"`
	Title          string                                          `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                                          `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date           *FHIRDateTime                                   `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept                               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose        string                                          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      string                                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind           string                                          `bson:"kind,omitempty" json:"kind,omitempty"`
	Software       *TerminologyCapabilitiesSoftwareComponent       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation *TerminologyCapabilitiesImplementationComponent `bson:"implementation,omitempty" json:"implementation,omitempty"`
	LockedDate     *bool                                           `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	CodeSystem     []TerminologyCapabilitiesCodeSystemComponent    `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	Expansion      *TerminologyCapabilitiesExpansionComponent      `bson:"expansion,omitempty" json:"expansion,omitempty"`
	CodeSearch     string                                          `bson:"codeSearch,omitempty" json:"codeSearch,omitempty"`
	ValidateCode   *TerminologyCapabilitiesValidateCodeComponent   `bson:"validateCode,omitempty" json:"validateCode,omitempty"`
	Translation    *TerminologyCapabilitiesTranslationComponent    `bson:"translation,omitempty" json:"translation,omitempty"`
	Closure        *TerminologyCapabilitiesClosureComponent        `bson:"closure,omitempty" json:"closure,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	x.ResourceType = "TerminologyCapabilities"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "terminologyCapabilities" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type terminologyCapabilities TerminologyCapabilities

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *TerminologyCapabilities) UnmarshalJSON(data []byte) (err error) {
	x2 := terminologyCapabilities{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = TerminologyCapabilities(x2)
		return x.checkResourceType()
	}
	return
}

func (x *TerminologyCapabilities) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "TerminologyCapabilities"
	} else if x.ResourceType != "TerminologyCapabilities" {
		return fmt.Errorf("Expected resourceType to be TerminologyCapabilities, instead received %s", x.ResourceType)
	}
	return nil
}

// TerminologyCapabilitiesSoftwareComponent ...
type TerminologyCapabilitiesSoftwareComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Version         string `bson:"version,omitempty" json:"version,omitempty"`
}

// TerminologyCapabilitiesImplementationComponent ...
type TerminologyCapabilitiesImplementationComponent struct {
	BackboneElement `bson:",inline"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	URL             *URL   `bson:"url,omitempty" json:"url,omitempty"`
}

// TerminologyCapabilitiesCodeSystemComponent ...
type TerminologyCapabilitiesCodeSystemComponent struct {
	BackboneElement `bson:",inline"`
	Uri             *Canonical                                          `bson:"uri,omitempty" json:"uri,omitempty"`
	Version         []TerminologyCapabilitiesCodeSystemVersionComponent `bson:"version,omitempty" json:"version,omitempty"`
	Subsumption     *bool                                               `bson:"subsumption,omitempty" json:"subsumption,omitempty"`
}

// TerminologyCapabilitiesCodeSystemVersionComponent ...
type TerminologyCapabilitiesCodeSystemVersionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string                                                    `bson:"code,omitempty" json:"code,omitempty"`
	IsDefault       *bool                                                     `bson:"isDefault,omitempty" json:"isDefault,omitempty"`
	Compositional   *bool                                                     `bson:"compositional,omitempty" json:"compositional,omitempty"`
	Language        []string                                                  `bson:"language,omitempty" json:"language,omitempty"`
	Filter          []TerminologyCapabilitiesCodeSystemVersionFilterComponent `bson:"filter,omitempty" json:"filter,omitempty"`
	Property        []string                                                  `bson:"property,omitempty" json:"property,omitempty"`
}

// TerminologyCapabilitiesCodeSystemVersionFilterComponent ...
type TerminologyCapabilitiesCodeSystemVersionFilterComponent struct {
	BackboneElement `bson:",inline"`
	Code            string   `bson:"code,omitempty" json:"code,omitempty"`
	Op              []string `bson:"op,omitempty" json:"op,omitempty"`
}

// TerminologyCapabilitiesExpansionComponent ...
type TerminologyCapabilitiesExpansionComponent struct {
	BackboneElement `bson:",inline"`
	Hierarchical    *bool                                                `bson:"hierarchical,omitempty" json:"hierarchical,omitempty"`
	Paging          *bool                                                `bson:"paging,omitempty" json:"paging,omitempty"`
	Incomplete      *bool                                                `bson:"incomplete,omitempty" json:"incomplete,omitempty"`
	Parameter       []TerminologyCapabilitiesExpansionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	TextFilter      string                                               `bson:"textFilter,omitempty" json:"textFilter,omitempty"`
}

// TerminologyCapabilitiesExpansionParameterComponent ...
type TerminologyCapabilitiesExpansionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// TerminologyCapabilitiesValidateCodeComponent ...
type TerminologyCapabilitiesValidateCodeComponent struct {
	BackboneElement `bson:",inline"`
	Translations    *bool `bson:"translations,omitempty" json:"translations,omitempty"`
}

// TerminologyCapabilitiesTranslationComponent ...
type TerminologyCapabilitiesTranslationComponent struct {
	BackboneElement `bson:",inline"`
	NeedsMap        *bool `bson:"needsMap,omitempty" json:"needsMap,omitempty"`
}

// TerminologyCapabilitiesClosureComponent ...
type TerminologyCapabilitiesClosureComponent struct {
	BackboneElement `bson:",inline"`
	Translation     *bool `bson:"translation,omitempty" json:"translation,omitempty"`
}

// TerminologyCapabilitiesPlus ...
type TerminologyCapabilitiesPlus struct {
	TerminologyCapabilities                     `bson:",inline"`
	TerminologyCapabilitiesPlusRelatedResources `bson:",inline"`
}

// TerminologyCapabilitiesPlusRelatedResources ...
type TerminologyCapabilitiesPlusRelatedResources struct {
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
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *t.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if t.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *t.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if t.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *t.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *t.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if t.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *t.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (t *TerminologyCapabilitiesPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
