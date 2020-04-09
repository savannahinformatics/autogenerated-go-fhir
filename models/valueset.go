package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ValueSet ... // TODO Write proper comment
type ValueSet struct {
	DomainResource `bson:",inline"`
	URL            string                      `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                      `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                      `bson:"name,omitempty" json:"name,omitempty"`
	Title          string                      `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date           *FHIRDateTime               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                      `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail             `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                      `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext              `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept           `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable      *bool                       `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose        string                      `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      string                      `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Compose        *ValueSetComposeComponent   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion      *ValueSetExpansionComponent `bson:"expansion,omitempty" json:"expansion,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *ValueSet) MarshalJSON() ([]byte, error) {
	x.ResourceType = "ValueSet"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "valueSet" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type valueSet ValueSet

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ValueSet) UnmarshalJSON(data []byte) (err error) {
	x2 := valueSet{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ValueSet(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ValueSet) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ValueSet"
	} else if x.ResourceType != "ValueSet" {
		return fmt.Errorf("Expected resourceType to be ValueSet, instead received %s", x.ResourceType)
	}
	return nil
}

// ValueSetComposeComponent ... // TODO Write proper comment
type ValueSetComposeComponent struct {
	BackboneElement `bson:",inline"`
	LockedDate      *FHIRDateTime                 `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive        *bool                         `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include         []ValueSetConceptSetComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude         []ValueSetConceptSetComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

// ValueSetConceptSetComponent ... // TODO Write proper comment
type ValueSetConceptSetComponent struct {
	BackboneElement `bson:",inline"`
	System          string                              `bson:"system,omitempty" json:"system,omitempty"`
	Version         string                              `bson:"version,omitempty" json:"version,omitempty"`
	Concept         []ValueSetConceptReferenceComponent `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter          []ValueSetConceptSetFilterComponent `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet        []Canonical                         `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}

// ValueSetConceptReferenceComponent ... // TODO Write proper comment
type ValueSetConceptReferenceComponent struct {
	BackboneElement `bson:",inline"`
	Code            string                                         `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                                         `bson:"display,omitempty" json:"display,omitempty"`
	Designation     []ValueSetConceptReferenceDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

// ValueSetConceptReferenceDesignationComponent ... // TODO Write proper comment
type ValueSetConceptReferenceDesignationComponent struct {
	BackboneElement `bson:",inline"`
	Language        string  `bson:"language,omitempty" json:"language,omitempty"`
	Use             *Coding `bson:"use,omitempty" json:"use,omitempty"`
	Value           string  `bson:"value,omitempty" json:"value,omitempty"`
}

// ValueSetConceptSetFilterComponent ... // TODO Write proper comment
type ValueSetConceptSetFilterComponent struct {
	BackboneElement `bson:",inline"`
	Property        string `bson:"property,omitempty" json:"property,omitempty"`
	Op              string `bson:"op,omitempty" json:"op,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

// ValueSetExpansionComponent ... // TODO Write proper comment
type ValueSetExpansionComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      string                                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp       *FHIRDateTime                         `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total           *int32                                `bson:"total,omitempty" json:"total,omitempty"`
	Offset          *int32                                `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter       []ValueSetExpansionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Contains        []ValueSetExpansionContainsComponent  `bson:"contains,omitempty" json:"contains,omitempty"`
}

// ValueSetExpansionParameterComponent ... // TODO Write proper comment
type ValueSetExpansionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Name            string        `bson:"name,omitempty" json:"name,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean    *bool         `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger    *int32        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal    *float64      `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueURI        string        `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode       string        `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
}

// ValueSetExpansionContainsComponent ... // TODO Write proper comment
type ValueSetExpansionContainsComponent struct {
	BackboneElement `bson:",inline"`
	System          string                                         `bson:"system,omitempty" json:"system,omitempty"`
	Abstract        *bool                                          `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive        *bool                                          `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version         string                                         `bson:"version,omitempty" json:"version,omitempty"`
	Code            string                                         `bson:"code,omitempty" json:"code,omitempty"`
	Display         string                                         `bson:"display,omitempty" json:"display,omitempty"`
	Designation     []ValueSetConceptReferenceDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	Contains        []ValueSetExpansionContainsComponent           `bson:"contains,omitempty" json:"contains,omitempty"`
}

// ValueSetPlus ... // TODO Write proper comment
type ValueSetPlus struct {
	ValueSet                     `bson:",inline"`
	ValueSetPlusRelatedResources `bson:",inline"`
}

// ValueSetPlusRelatedResources ... // TODO Write proper comment
type ValueSetPlusRelatedResources struct {
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
	RevIncludedStructureDefinitionResourcesReferencingValueset             *[]StructureDefinition        `bson:"_revIncludedStructureDefinitionResourcesReferencingValueset,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedConceptMapResourcesReferencingSource                        *[]ConceptMap                 `bson:"_revIncludedConceptMapResourcesReferencingSource,omitempty"`
	RevIncludedConceptMapResourcesReferencingSourceuri                     *[]ConceptMap                 `bson:"_revIncludedConceptMapResourcesReferencingSourceuri,omitempty"`
	RevIncludedConceptMapResourcesReferencingTargeturi                     *[]ConceptMap                 `bson:"_revIncludedConceptMapResourcesReferencingTargeturi,omitempty"`
	RevIncludedConceptMapResourcesReferencingTarget                        *[]ConceptMap                 `bson:"_revIncludedConceptMapResourcesReferencingTarget,omitempty"`
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

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if v.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *v.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if v.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *v.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if v.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *v.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if v.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *v.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if v.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *v.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if v.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *v.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedStructureDefinitionResourcesReferencingValueset ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedStructureDefinitionResourcesReferencingValueset() (structureDefinitions []StructureDefinition, err error) {
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset == nil {
		err = errors.New("RevIncluded structureDefinitions not requested")
	} else {
		structureDefinitions = *v.RevIncludedStructureDefinitionResourcesReferencingValueset
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if v.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *v.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedConceptMapResourcesReferencingSource ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSource() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingSource == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingSource
	}
	return
}

// GetRevIncludedConceptMapResourcesReferencingSourceuri ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingSourceuri() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingSourceuri == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingSourceuri
	}
	return
}

// GetRevIncludedConceptMapResourcesReferencingTargeturi ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingTargeturi() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingTargeturi == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingTargeturi
	}
	return
}

// GetRevIncludedConceptMapResourcesReferencingTarget ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConceptMapResourcesReferencingTarget() (conceptMaps []ConceptMap, err error) {
	if v.RevIncludedConceptMapResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded conceptMaps not requested")
	} else {
		conceptMaps = *v.RevIncludedConceptMapResourcesReferencingTarget
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if v.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *v.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *v.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if v.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *v.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if v.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *v.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *v.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *v.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*v.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingData {
			rsc := (*v.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*v.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*v.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*v.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *v.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*v.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset != nil {
		for idx := range *v.RevIncludedStructureDefinitionResourcesReferencingValueset {
			rsc := (*v.RevIncludedStructureDefinitionResourcesReferencingValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTargeturi != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTargeturi {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTargeturi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*v.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ... // TODO Write proper comment
func (v *ValueSetPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *v.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*v.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*v.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*v.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *v.RevIncludedConsentResourcesReferencingData {
			rsc := (*v.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*v.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *v.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*v.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*v.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedContractResourcesReferencingSubject {
			rsc := (*v.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*v.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *v.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*v.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*v.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*v.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*v.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*v.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *v.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*v.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*v.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*v.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*v.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*v.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedStructureDefinitionResourcesReferencingValueset != nil {
		for idx := range *v.RevIncludedStructureDefinitionResourcesReferencingValueset {
			rsc := (*v.RevIncludedStructureDefinitionResourcesReferencingValueset)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSource != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSource {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingSourceuri != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingSourceuri {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingSourceuri)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTargeturi != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTargeturi {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTargeturi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConceptMapResourcesReferencingTarget != nil {
		for idx := range *v.RevIncludedConceptMapResourcesReferencingTarget {
			rsc := (*v.RevIncludedConceptMapResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*v.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *v.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*v.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*v.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*v.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *v.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*v.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*v.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *v.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*v.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *v.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*v.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *v.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*v.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*v.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *v.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*v.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*v.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*v.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
