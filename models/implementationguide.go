package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ImplementationGuide ... // TODO Write proper comment
type ImplementationGuide struct {
	DomainResource `bson:",inline"`
	Url            string                                  `bson:"url,omitempty" json:"url,omitempty"`
	Version        string                                  `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                  `bson:"name,omitempty" json:"name,omitempty"`
	Title          string                                  `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                                  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date           *FHIRDateTime                           `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                                  `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept                       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright      string                                  `bson:"copyright,omitempty" json:"copyright,omitempty"`
	PackageId      string                                  `bson:"packageId,omitempty" json:"packageId,omitempty"`
	License        string                                  `bson:"license,omitempty" json:"license,omitempty"`
	FhirVersion    []string                                `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	DependsOn      []ImplementationGuideDependsOnComponent `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	Global         []ImplementationGuideGlobalComponent    `bson:"global,omitempty" json:"global,omitempty"`
	Definition     *ImplementationGuideDefinitionComponent `bson:"definition,omitempty" json:"definition,omitempty"`
	Manifest       *ImplementationGuideManifestComponent   `bson:"manifest,omitempty" json:"manifest,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *ImplementationGuide) MarshalJSON() ([]byte, error) {
	x.ResourceType = "ImplementationGuide"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "implementationGuide" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type implementationGuide ImplementationGuide

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ImplementationGuide) UnmarshalJSON(data []byte) (err error) {
	x2 := implementationGuide{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImplementationGuide(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImplementationGuide) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImplementationGuide"
	} else if x.ResourceType != "ImplementationGuide" {
		return fmt.Errorf("Expected resourceType to be ImplementationGuide, instead received %s", x.ResourceType)
	}
	return nil
}

// ImplementationGuideDependsOnComponent ... // TODO Write proper comment
type ImplementationGuideDependsOnComponent struct {
	BackboneElement `bson:",inline"`
	Uri             *Canonical `bson:"uri,omitempty" json:"uri,omitempty"`
	PackageId       string     `bson:"packageId,omitempty" json:"packageId,omitempty"`
	Version         string     `bson:"version,omitempty" json:"version,omitempty"`
}

// ImplementationGuideGlobalComponent ... // TODO Write proper comment
type ImplementationGuideGlobalComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Profile         *Canonical `bson:"profile,omitempty" json:"profile,omitempty"`
}

// ImplementationGuideDefinitionComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionComponent struct {
	BackboneElement `bson:",inline"`
	Grouping        []ImplementationGuideDefinitionGroupingComponent  `bson:"grouping,omitempty" json:"grouping,omitempty"`
	Resource        []ImplementationGuideDefinitionResourceComponent  `bson:"resource,omitempty" json:"resource,omitempty"`
	Page            *ImplementationGuideDefinitionPageComponent       `bson:"page,omitempty" json:"page,omitempty"`
	Parameter       []ImplementationGuideDefinitionParameterComponent `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Template        []ImplementationGuideDefinitionTemplateComponent  `bson:"template,omitempty" json:"template,omitempty"`
}

// ImplementationGuideDefinitionGroupingComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionGroupingComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
}

// ImplementationGuideDefinitionResourceComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionResourceComponent struct {
	BackboneElement  `bson:",inline"`
	Reference        *Reference `bson:"reference,omitempty" json:"reference,omitempty"`
	FhirVersion      []string   `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Name             string     `bson:"name,omitempty" json:"name,omitempty"`
	Description      string     `bson:"description,omitempty" json:"description,omitempty"`
	ExampleBoolean   *bool      `bson:"exampleBoolean,omitempty" json:"exampleBoolean,omitempty"`
	ExampleCanonical *Canonical `bson:"exampleCanonical,omitempty" json:"exampleCanonical,omitempty"`
	GroupingId       string     `bson:"groupingId,omitempty" json:"groupingId,omitempty"`
}

// ImplementationGuideDefinitionPageComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionPageComponent struct {
	BackboneElement `bson:",inline"`
	NameUrl         *URL                                         `bson:"nameUrl,omitempty" json:"nameUrl,omitempty"`
	NameReference   *Reference                                   `bson:"nameReference,omitempty" json:"nameReference,omitempty"`
	Title           string                                       `bson:"title,omitempty" json:"title,omitempty"`
	Generation      string                                       `bson:"generation,omitempty" json:"generation,omitempty"`
	Page            []ImplementationGuideDefinitionPageComponent `bson:"page,omitempty" json:"page,omitempty"`
}

// ImplementationGuideDefinitionParameterComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionParameterComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

// ImplementationGuideDefinitionTemplateComponent ... // TODO Write proper comment
type ImplementationGuideDefinitionTemplateComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Source          string `bson:"source,omitempty" json:"source,omitempty"`
	Scope           string `bson:"scope,omitempty" json:"scope,omitempty"`
}

// ImplementationGuideManifestComponent ... // TODO Write proper comment
type ImplementationGuideManifestComponent struct {
	BackboneElement `bson:",inline"`
	Rendering       *URL                                           `bson:"rendering,omitempty" json:"rendering,omitempty"`
	Resource        []ImplementationGuideManifestResourceComponent `bson:"resource,omitempty" json:"resource,omitempty"`
	Page            []ImplementationGuideManifestPageComponent     `bson:"page,omitempty" json:"page,omitempty"`
	Image           []string                                       `bson:"image,omitempty" json:"image,omitempty"`
	Other           []string                                       `bson:"other,omitempty" json:"other,omitempty"`
}

// ImplementationGuideManifestResourceComponent ... // TODO Write proper comment
type ImplementationGuideManifestResourceComponent struct {
	BackboneElement  `bson:",inline"`
	Reference        *Reference `bson:"reference,omitempty" json:"reference,omitempty"`
	ExampleBoolean   *bool      `bson:"exampleBoolean,omitempty" json:"exampleBoolean,omitempty"`
	ExampleCanonical *Canonical `bson:"exampleCanonical,omitempty" json:"exampleCanonical,omitempty"`
	RelativePath     *URL       `bson:"relativePath,omitempty" json:"relativePath,omitempty"`
}

// ImplementationGuideManifestPageComponent ... // TODO Write proper comment
type ImplementationGuideManifestPageComponent struct {
	BackboneElement `bson:",inline"`
	Name            string   `bson:"name,omitempty" json:"name,omitempty"`
	Title           string   `bson:"title,omitempty" json:"title,omitempty"`
	Anchor          []string `bson:"anchor,omitempty" json:"anchor,omitempty"`
}

// ImplementationGuidePlus ... // TODO Write proper comment
type ImplementationGuidePlus struct {
	ImplementationGuide                     `bson:",inline"`
	ImplementationGuidePlusRelatedResources `bson:",inline"`
}

// ImplementationGuidePlusRelatedResources ... // TODO Write proper comment
type ImplementationGuidePlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedByGlobal                 *[]StructureDefinition        `bson:"_includedStructureDefinitionResourcesReferencedByGlobal,omitempty"`
	IncludedImplementationGuideResourcesReferencedByDependson              *[]ImplementationGuide        `bson:"_includedImplementationGuideResourcesReferencedByDependson,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedCapabilityStatementResourcesReferencingGuide                *[]CapabilityStatement        `bson:"_revIncludedCapabilityStatementResourcesReferencingGuide,omitempty"`
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
	RevIncludedImplementationGuideResourcesReferencingDependson            *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingDependson,omitempty"`
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

// GetIncludedStructureDefinitionResourceReferencedByGlobal ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByGlobal() (structureDefinition *StructureDefinition, err error) {
	if i.IncludedStructureDefinitionResourcesReferencedByGlobal == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*i.IncludedStructureDefinitionResourcesReferencedByGlobal) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*i.IncludedStructureDefinitionResourcesReferencedByGlobal))
	} else if len(*i.IncludedStructureDefinitionResourcesReferencedByGlobal) == 1 {
		structureDefinition = &(*i.IncludedStructureDefinitionResourcesReferencedByGlobal)[0]
	}
	return
}

// GetIncludedImplementationGuideResourceReferencedByDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetIncludedImplementationGuideResourceReferencedByDependson() (implementationGuide *ImplementationGuide, err error) {
	if i.IncludedImplementationGuideResourcesReferencedByDependson == nil {
		err = errors.New("Included implementationguides not requested")
	} else if len(*i.IncludedImplementationGuideResourcesReferencedByDependson) > 1 {
		err = fmt.Errorf("Expected 0 or 1 implementationGuide, but found %d", len(*i.IncludedImplementationGuideResourcesReferencedByDependson))
	} else if len(*i.IncludedImplementationGuideResourcesReferencedByDependson) == 1 {
		implementationGuide = &(*i.IncludedImplementationGuideResourcesReferencedByDependson)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *i.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if i.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *i.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedCapabilityStatementResourcesReferencingGuide ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCapabilityStatementResourcesReferencingGuide() (capabilityStatements []CapabilityStatement, err error) {
	if i.RevIncludedCapabilityStatementResourcesReferencingGuide == nil {
		err = errors.New("RevIncluded capabilityStatements not requested")
	} else {
		capabilityStatements = *i.RevIncludedCapabilityStatementResourcesReferencingGuide
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if i.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *i.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingDependson() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingDependson
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *i.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if i.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *i.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedStructureDefinitionResourcesReferencedByGlobal != nil {
		for idx := range *i.IncludedStructureDefinitionResourcesReferencedByGlobal {
			rsc := (*i.IncludedStructureDefinitionResourcesReferencedByGlobal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImplementationGuideResourcesReferencedByDependson != nil {
		for idx := range *i.IncludedImplementationGuideResourcesReferencedByDependson {
			rsc := (*i.IncludedImplementationGuideResourcesReferencedByDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCapabilityStatementResourcesReferencingGuide != nil {
		for idx := range *i.RevIncludedCapabilityStatementResourcesReferencingGuide {
			rsc := (*i.RevIncludedCapabilityStatementResourcesReferencingGuide)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingDependson {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ... // TODO Write proper comment
func (i *ImplementationGuidePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedStructureDefinitionResourcesReferencedByGlobal != nil {
		for idx := range *i.IncludedStructureDefinitionResourcesReferencedByGlobal {
			rsc := (*i.IncludedStructureDefinitionResourcesReferencedByGlobal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImplementationGuideResourcesReferencedByDependson != nil {
		for idx := range *i.IncludedImplementationGuideResourcesReferencedByDependson {
			rsc := (*i.IncludedImplementationGuideResourcesReferencedByDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCapabilityStatementResourcesReferencingGuide != nil {
		for idx := range *i.RevIncludedCapabilityStatementResourcesReferencingGuide {
			rsc := (*i.RevIncludedCapabilityStatementResourcesReferencingGuide)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingDependson {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
