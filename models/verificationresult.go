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

type VerificationResult struct {
	DomainResource    `bson:",inline"`
	Target            []Reference                                `bson:"target,omitempty" json:"target,omitempty"`
	TargetLocation    []string                                   `bson:"targetLocation,omitempty" json:"targetLocation,omitempty"`
	Need              *CodeableConcept                           `bson:"need,omitempty" json:"need,omitempty"`
	Status            string                                     `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate        *FHIRDateTime                              `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                           `bson:"validationType,omitempty" json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                          `bson:"validationProcess,omitempty" json:"validationProcess,omitempty"`
	Frequency         *Timing                                    `bson:"frequency,omitempty" json:"frequency,omitempty"`
	LastPerformed     *FHIRDateTime                              `bson:"lastPerformed,omitempty" json:"lastPerformed,omitempty"`
	NextScheduled     *FHIRDateTime                              `bson:"nextScheduled,omitempty" json:"nextScheduled,omitempty"`
	FailureAction     *CodeableConcept                           `bson:"failureAction,omitempty" json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySourceComponent `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestationComponent    `bson:"attestation,omitempty" json:"attestation,omitempty"`
	Validator         []VerificationResultValidatorComponent     `bson:"validator,omitempty" json:"validator,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *VerificationResult) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "VerificationResult"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "verificationResult" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type verificationResult VerificationResult

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *VerificationResult) UnmarshalJSON(data []byte) (err error) {
	x2 := verificationResult{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = VerificationResult(x2)
		return x.checkResourceType()
	}
	return
}

func (x *VerificationResult) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "VerificationResult"
	} else if x.ResourceType != "VerificationResult" {
		return errors.New(fmt.Sprintf("Expected resourceType to be VerificationResult, instead received %s", x.ResourceType))
	}
	return nil
}

type VerificationResultPrimarySourceComponent struct {
	BackboneElement     `bson:",inline"`
	Who                 *Reference        `bson:"who,omitempty" json:"who,omitempty"`
	Type                []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `bson:"validationStatus,omitempty" json:"validationStatus,omitempty"`
	ValidationDate      *FHIRDateTime     `bson:"validationDate,omitempty" json:"validationDate,omitempty"`
	CanPushUpdates      *CodeableConcept  `bson:"canPushUpdates,omitempty" json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `bson:"pushTypeAvailable,omitempty" json:"pushTypeAvailable,omitempty"`
}

type VerificationResultAttestationComponent struct {
	BackboneElement           `bson:",inline"`
	Who                       *Reference       `bson:"who,omitempty" json:"who,omitempty"`
	OnBehalfOf                *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	Date                      *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	SourceIdentityCertificate string           `bson:"sourceIdentityCertificate,omitempty" json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  string           `bson:"proxyIdentityCertificate,omitempty" json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `bson:"proxySignature,omitempty" json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `bson:"sourceSignature,omitempty" json:"sourceSignature,omitempty"`
}

type VerificationResultValidatorComponent struct {
	BackboneElement      `bson:",inline"`
	Organization         *Reference `bson:"organization,omitempty" json:"organization,omitempty"`
	IdentityCertificate  string     `bson:"identityCertificate,omitempty" json:"identityCertificate,omitempty"`
	AttestationSignature *Signature `bson:"attestationSignature,omitempty" json:"attestationSignature,omitempty"`
}

type VerificationResultPlus struct {
	VerificationResult                     `bson:",inline"`
	VerificationResultPlusRelatedResources `bson:",inline"`
}

type VerificationResultPlusRelatedResources struct {
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

func (v *VerificationResultPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if v.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *v.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if v.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *v.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if v.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *v.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if v.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *v.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *v.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if v.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *v.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if v.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *v.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *v.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if v.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *v.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *v.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if v.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *v.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *v.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if v.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *v.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if v.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *v.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if v.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *v.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if v.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *v.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if v.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *v.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if v.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *v.RevIncludedListResourcesReferencingItem
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if v.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *v.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if v.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *v.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if v.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *v.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if v.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *v.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if v.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *v.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if v.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *v.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if v.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *v.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *v.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *v.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *v.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (v *VerificationResultPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (v *VerificationResultPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
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

func (v *VerificationResultPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
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
	if v.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *v.RevIncludedListResourcesReferencingItem {
			rsc := (*v.RevIncludedListResourcesReferencingItem)[idx]
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
