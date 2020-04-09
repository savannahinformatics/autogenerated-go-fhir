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

type ObservationDefinition struct {
	DomainResource         `bson:",inline"`
	Category               []CodeableConcept                                  `bson:"category,omitempty" json:"category,omitempty"`
	Code                   *CodeableConcept                                   `bson:"code,omitempty" json:"code,omitempty"`
	Identifier             []Identifier                                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	PermittedDataType      []string                                           `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                              `bson:"multipleResultsAllowed,omitempty" json:"multipleResultsAllowed,omitempty"`
	Method                 *CodeableConcept                                   `bson:"method,omitempty" json:"method,omitempty"`
	PreferredReportName    string                                             `bson:"preferredReportName,omitempty" json:"preferredReportName,omitempty"`
	QuantitativeDetails    *ObservationDefinitionQuantitativeDetailsComponent `bson:"quantitativeDetails,omitempty" json:"quantitativeDetails,omitempty"`
	QualifiedInterval      []ObservationDefinitionQualifiedIntervalComponent  `bson:"qualifiedInterval,omitempty" json:"qualifiedInterval,omitempty"`
	ValidCodedValueSet     *Reference                                         `bson:"validCodedValueSet,omitempty" json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet    *Reference                                         `bson:"normalCodedValueSet,omitempty" json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet  *Reference                                         `bson:"abnormalCodedValueSet,omitempty" json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet  *Reference                                         `bson:"criticalCodedValueSet,omitempty" json:"criticalCodedValueSet,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *ObservationDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ObservationDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "observationDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observationDefinition ObservationDefinition

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ObservationDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := observationDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ObservationDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ObservationDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ObservationDefinition"
	} else if x.ResourceType != "ObservationDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ObservationDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type ObservationDefinitionQuantitativeDetailsComponent struct {
	BackboneElement  `bson:",inline"`
	CustomaryUnit    *CodeableConcept `bson:"customaryUnit,omitempty" json:"customaryUnit,omitempty"`
	Unit             *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty"`
	ConversionFactor *float64         `bson:"conversionFactor,omitempty" json:"conversionFactor,omitempty"`
	DecimalPrecision *int32           `bson:"decimalPrecision,omitempty" json:"decimalPrecision,omitempty"`
}

type ObservationDefinitionQualifiedIntervalComponent struct {
	BackboneElement `bson:",inline"`
	Category        string            `bson:"category,omitempty" json:"category,omitempty"`
	Range           *Range            `bson:"range,omitempty" json:"range,omitempty"`
	Context         *CodeableConcept  `bson:"context,omitempty" json:"context,omitempty"`
	AppliesTo       []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Gender          string            `bson:"gender,omitempty" json:"gender,omitempty"`
	Age             *Range            `bson:"age,omitempty" json:"age,omitempty"`
	GestationalAge  *Range            `bson:"gestationalAge,omitempty" json:"gestationalAge,omitempty"`
	Condition       string            `bson:"condition,omitempty" json:"condition,omitempty"`
}

type ObservationDefinitionPlus struct {
	ObservationDefinition                     `bson:",inline"`
	ObservationDefinitionPlusRelatedResources `bson:",inline"`
}

type ObservationDefinitionPlusRelatedResources struct {
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

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *o.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if o.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *o.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if o.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *o.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *ObservationDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (o *ObservationDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *ObservationDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
