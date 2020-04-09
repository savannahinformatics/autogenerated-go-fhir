// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MedicinalProduct struct {
	DomainResource                 `bson:",inline"`
	Identifier                     []Identifier                                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                           *CodeableConcept                                          `bson:"type,omitempty" json:"type,omitempty"`
	Domain                         *Coding                                                   `bson:"domain,omitempty" json:"domain,omitempty"`
	CombinedPharmaceuticalDoseForm *CodeableConcept                                          `bson:"combinedPharmaceuticalDoseForm,omitempty" json:"combinedPharmaceuticalDoseForm,omitempty"`
	LegalStatusOfSupply            *CodeableConcept                                          `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	AdditionalMonitoringIndicator  *CodeableConcept                                          `bson:"additionalMonitoringIndicator,omitempty" json:"additionalMonitoringIndicator,omitempty"`
	SpecialMeasures                []string                                                  `bson:"specialMeasures,omitempty" json:"specialMeasures,omitempty"`
	PaediatricUseIndicator         *CodeableConcept                                          `bson:"paediatricUseIndicator,omitempty" json:"paediatricUseIndicator,omitempty"`
	ProductClassification          []CodeableConcept                                         `bson:"productClassification,omitempty" json:"productClassification,omitempty"`
	MarketingStatus                []MarketingStatus                                         `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	PharmaceuticalProduct          []Reference                                               `bson:"pharmaceuticalProduct,omitempty" json:"pharmaceuticalProduct,omitempty"`
	PackagedMedicinalProduct       []Reference                                               `bson:"packagedMedicinalProduct,omitempty" json:"packagedMedicinalProduct,omitempty"`
	AttachedDocument               []Reference                                               `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	MasterFile                     []Reference                                               `bson:"masterFile,omitempty" json:"masterFile,omitempty"`
	Contact                        []Reference                                               `bson:"contact,omitempty" json:"contact,omitempty"`
	ClinicalTrial                  []Reference                                               `bson:"clinicalTrial,omitempty" json:"clinicalTrial,omitempty"`
	Name                           []MedicinalProductNameComponent                           `bson:"name,omitempty" json:"name,omitempty"`
	CrossReference                 []Identifier                                              `bson:"crossReference,omitempty" json:"crossReference,omitempty"`
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperationComponent `bson:"manufacturingBusinessOperation,omitempty" json:"manufacturingBusinessOperation,omitempty"`
	SpecialDesignation             []MedicinalProductSpecialDesignationComponent             `bson:"specialDesignation,omitempty" json:"specialDesignation,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicinalProduct) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicinalProduct"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "medicinalProduct" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicinalProduct MedicinalProduct

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *MedicinalProduct) UnmarshalJSON(data []byte) (err error) {
	x2 := medicinalProduct{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicinalProduct(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicinalProduct) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicinalProduct"
	} else if x.ResourceType != "MedicinalProduct" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicinalProduct, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicinalProductNameComponent struct {
	BackboneElement `bson:",inline"`
	ProductName     string                                         `bson:"productName,omitempty" json:"productName,omitempty"`
	NamePart        []MedicinalProductNameNamePartComponent        `bson:"namePart,omitempty" json:"namePart,omitempty"`
	CountryLanguage []MedicinalProductNameCountryLanguageComponent `bson:"countryLanguage,omitempty" json:"countryLanguage,omitempty"`
}

type MedicinalProductNameNamePartComponent struct {
	BackboneElement `bson:",inline"`
	Part            string  `bson:"part,omitempty" json:"part,omitempty"`
	Type            *Coding `bson:"type,omitempty" json:"type,omitempty"`
}

type MedicinalProductNameCountryLanguageComponent struct {
	BackboneElement `bson:",inline"`
	Country         *CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
	Jurisdiction    *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

type MedicinalProductManufacturingBusinessOperationComponent struct {
	BackboneElement              `bson:",inline"`
	OperationType                *CodeableConcept `bson:"operationType,omitempty" json:"operationType,omitempty"`
	AuthorisationReferenceNumber *Identifier      `bson:"authorisationReferenceNumber,omitempty" json:"authorisationReferenceNumber,omitempty"`
	EffectiveDate                *FHIRDateTime    `bson:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
	ConfidentialityIndicator     *CodeableConcept `bson:"confidentialityIndicator,omitempty" json:"confidentialityIndicator,omitempty"`
	Manufacturer                 []Reference      `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Regulator                    *Reference       `bson:"regulator,omitempty" json:"regulator,omitempty"`
}

type MedicinalProductSpecialDesignationComponent struct {
	BackboneElement           `bson:",inline"`
	Identifier                []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                      *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	IntendedUse               *CodeableConcept `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
	IndicationCodeableConcept *CodeableConcept `bson:"indicationCodeableConcept,omitempty" json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference       `bson:"indicationReference,omitempty" json:"indicationReference,omitempty"`
	Status                    *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	Date                      *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Species                   *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
}

type MedicinalProductPlus struct {
	MedicinalProduct                     `bson:",inline"`
	MedicinalProductPlusRelatedResources `bson:",inline"`
}

type MedicinalProductPlusRelatedResources struct {
	RevIncludedAppointmentResourcesReferencingSupportinginfo                *[]Appointment                       `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                 *[]EventDefinition                   `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom               *[]EventDefinition                   `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor               *[]EventDefinition                   `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof                *[]EventDefinition                   `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                 *[]EventDefinition                   `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                     *[]DocumentManifest                  `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref               *[]DocumentManifest                  `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedMedicinalProductPackagedResourcesReferencingSubject          *[]MedicinalProductPackaged          `bson:"_revIncludedMedicinalProductPackagedResourcesReferencingSubject,omitempty"`
	RevIncludedConsentResourcesReferencingData                              *[]Consent                           `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                         *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                       *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                       *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                        *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                    *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                    *[]Measure                           `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                 *[]DocumentReference                 `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource           *[]MeasureReport                     `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                 *[]VerificationResult                `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                          *[]Contract                          `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                     *[]PaymentNotice                     `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                    *[]PaymentNotice                     `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor              *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom            *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor            *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof             *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1         *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2         *[]ResearchDefinition                `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource              *[]ImplementationGuide               `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor       *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom     *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor     *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof      *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1  *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2  *[]ResearchElementDefinition         `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                      *[]Communication                     `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                     *[]Communication                     `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor              *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom            *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor            *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof             *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1         *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2         *[]ActivityDefinition                `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                              *[]Linkage                           `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                            *[]Linkage                           `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                     *[]DeviceRequest                     `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest                *[]DeviceRequest                     `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                       *[]MessageHeader                     `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation    *[]ImmunizationRecommendation        `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                         *[]Provenance                        `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                         *[]Provenance                        `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                              *[]Task                              `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                                *[]Task                              `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                              *[]Task                              `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                                 *[]List                              `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedMedicinalProductContraindicationResourcesReferencingSubject  *[]MedicinalProductContraindication  `bson:"_revIncludedMedicinalProductContraindicationResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor                *[]EvidenceVariable                  `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom              *[]EvidenceVariable                  `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor              *[]EvidenceVariable                  `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof               *[]EvidenceVariable                  `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson                *[]EvidenceVariable                  `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                         *[]Observation                       `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                         *[]Library                           `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                       *[]Library                           `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                       *[]Library                           `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                        *[]Library                           `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                         *[]Library                           `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicinalProductInteractionResourcesReferencingSubject       *[]MedicinalProductInteraction       `bson:"_revIncludedMedicinalProductInteractionResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon              *[]CommunicationRequest              `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                             *[]Basic                             `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicinalProductIndicationResourcesReferencingSubject        *[]MedicinalProductIndication        `bson:"_revIncludedMedicinalProductIndicationResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                        *[]Evidence                          `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                      *[]Evidence                          `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                      *[]Evidence                          `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                       *[]Evidence                          `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                        *[]Evidence                          `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                         *[]AuditEvent                        `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                  *[]Condition                         `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                       *[]Composition                       `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                         *[]Composition                       `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                  *[]DetectedIssue                     `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject             *[]QuestionnaireResponse             `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject *[]MedicinalProductUndesirableEffect `bson:"_revIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo         *[]ClinicalImpression                `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                  *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom                *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor                *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                 *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1             *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2             *[]PlanDefinition                    `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject     *[]MedicinalProductAuthorization     `bson:"_revIncludedMedicinalProductAuthorizationResourcesReferencingSubject,omitempty"`
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *m.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductPackagedResourcesReferencingSubject() (medicinalProductPackageds []MedicinalProductPackaged, err error) {
	if m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductPackageds not requested")
	} else {
		medicinalProductPackageds = *m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if m.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *m.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductContraindicationResourcesReferencingSubject() (medicinalProductContraindications []MedicinalProductContraindication, err error) {
	if m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductContraindications not requested")
	} else {
		medicinalProductContraindications = *m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if m.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *m.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductInteractionResourcesReferencingSubject() (medicinalProductInteractions []MedicinalProductInteraction, err error) {
	if m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductInteractions not requested")
	} else {
		medicinalProductInteractions = *m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *m.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductIndicationResourcesReferencingSubject() (medicinalProductIndications []MedicinalProductIndication, err error) {
	if m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductIndications not requested")
	} else {
		medicinalProductIndications = *m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if m.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *m.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject() (medicinalProductUndesirableEffects []MedicinalProductUndesirableEffect, err error) {
	if m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductUndesirableEffects not requested")
	} else {
		medicinalProductUndesirableEffects = *m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedMedicinalProductAuthorizationResourcesReferencingSubject() (medicinalProductAuthorizations []MedicinalProductAuthorization, err error) {
	if m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicinalProductAuthorizations not requested")
	} else {
		medicinalProductAuthorizations = *m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject
	}
	return
}

func (m *MedicinalProductPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (m *MedicinalProductPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicinalProductPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductPackagedResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductContraindicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductInteractionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductIndicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductUndesirableEffectResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject {
			rsc := (*m.RevIncludedMedicinalProductAuthorizationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
