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
	"gopkg.in/mgo.v2/bson"
)

type Bundle struct {
	Resource   `bson:",inline"`
	Identifier *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type       string                 `bson:"type,omitempty" json:"type,omitempty"`
	Timestamp  *FHIRDateTime          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total      *uint32                `bson:"total,omitempty" json:"total,omitempty"`
	Link       []BundleLinkComponent  `bson:"link,omitempty" json:"link,omitempty"`
	Entry      []BundleEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	Signature  *Signature             `bson:"signature,omitempty" json:"signature,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Bundle) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Bundle"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "bundle" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundle Bundle

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Bundle) UnmarshalJSON(data []byte) (err error) {
	x2 := bundle{}
	if err = json.Unmarshal(data, &x2); err == nil {
		*x = Bundle(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Bundle) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Bundle"
	} else if x.ResourceType != "Bundle" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Bundle, instead received %s", x.ResourceType))
	}
	return nil
}

type BundleLinkComponent struct {
	BackboneElement `bson:",inline"`
	Relation        string `bson:"relation,omitempty" json:"relation,omitempty"`
	Url             string `bson:"url,omitempty" json:"url,omitempty"`
}

type BundleEntryComponent struct {
	BackboneElement `bson:",inline"`
	Link            []BundleLinkComponent         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl         string                        `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource        interface{}                   `bson:"resource,omitempty" json:"resource,omitempty"`
	Search          *BundleEntrySearchComponent   `bson:"search,omitempty" json:"search,omitempty"`
	Request         *BundleEntryRequestComponent  `bson:"request,omitempty" json:"request,omitempty"`
	Response        *BundleEntryResponseComponent `bson:"response,omitempty" json:"response,omitempty"`
}

// "bundleEntryComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundleEntryComponent BundleEntryComponent

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *BundleEntryComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := bundleEntryComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Resource != nil {
			x2.Resource = MapToResource(x2.Resource, true)
		}
		*x = BundleEntryComponent(x2)
	}
	return
}

type BundleEntrySearchComponent struct {
	BackboneElement `bson:",inline"`
	Mode            string   `bson:"mode,omitempty" json:"mode,omitempty"`
	Score           *float64 `bson:"score,omitempty" json:"score,omitempty"`
}

type BundleEntryRequestComponent struct {
	BackboneElement `bson:",inline"`
	Method          string        `bson:"method,omitempty" json:"method,omitempty"`
	Url             string        `bson:"url,omitempty" json:"url,omitempty"`
	IfNoneMatch     string        `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince *FHIRDateTime `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch         string        `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist     string        `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}

type BundleEntryResponseComponent struct {
	BackboneElement `bson:",inline"`
	Status          string        `bson:"status,omitempty" json:"status,omitempty"`
	Location        string        `bson:"location,omitempty" json:"location,omitempty"`
	Etag            string        `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified    *FHIRDateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Outcome         interface{}   `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

// "bundleEntryResponseComponent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type bundleEntryResponseComponent BundleEntryResponseComponent

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *BundleEntryResponseComponent) UnmarshalJSON(data []byte) (err error) {
	x2 := bundleEntryResponseComponent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Outcome != nil {
			x2.Outcome = MapToResource(x2.Outcome, true)
		}
		*x = BundleEntryResponseComponent(x2)
	}
	return
}

type BundlePlus struct {
	Bundle                     `bson:",inline"`
	BundlePlusRelatedResources `bson:",inline"`
}

type BundlePlusRelatedResources struct {
	IncludedCompositionResourcesReferencedByComposition                    *[]Composition                `bson:"_includedCompositionResourcesReferencedByComposition,omitempty"`
	IncludedMessageHeaderResourcesReferencedByMessage                      *[]MessageHeader              `bson:"_includedMessageHeaderResourcesReferencedByMessage,omitempty"`
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

func (b *BundlePlusRelatedResources) GetIncludedCompositionResourceReferencedByComposition() (composition *Composition, err error) {
	if b.IncludedCompositionResourcesReferencedByComposition == nil {
		err = errors.New("Included compositions not requested")
	} else if len(*b.IncludedCompositionResourcesReferencedByComposition) > 1 {
		err = fmt.Errorf("Expected 0 or 1 composition, but found %d", len(*b.IncludedCompositionResourcesReferencedByComposition))
	} else if len(*b.IncludedCompositionResourcesReferencedByComposition) == 1 {
		composition = &(*b.IncludedCompositionResourcesReferencedByComposition)[0]
	}
	return
}

func (b *BundlePlusRelatedResources) GetIncludedMessageHeaderResourceReferencedByMessage() (messageHeader *MessageHeader, err error) {
	if b.IncludedMessageHeaderResourcesReferencedByMessage == nil {
		err = errors.New("Included messageheaders not requested")
	} else if len(*b.IncludedMessageHeaderResourcesReferencedByMessage) > 1 {
		err = fmt.Errorf("Expected 0 or 1 messageHeader, but found %d", len(*b.IncludedMessageHeaderResourcesReferencedByMessage))
	} else if len(*b.IncludedMessageHeaderResourcesReferencedByMessage) == 1 {
		messageHeader = &(*b.IncludedMessageHeaderResourcesReferencedByMessage)[0]
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *b.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if b.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *b.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *b.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if b.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *b.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if b.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *b.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if b.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *b.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if b.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *b.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if b.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *b.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if b.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *b.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if b.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *b.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *b.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if b.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *b.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if b.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *b.RevIncludedListResourcesReferencingItem
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if b.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *b.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *b.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if b.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *b.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if b.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *b.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if b.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *b.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *b.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BundlePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BundlePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedCompositionResourcesReferencedByComposition != nil {
		for idx := range *b.IncludedCompositionResourcesReferencedByComposition {
			rsc := (*b.IncludedCompositionResourcesReferencedByComposition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedMessageHeaderResourcesReferencedByMessage != nil {
		for idx := range *b.IncludedMessageHeaderResourcesReferencedByMessage {
			rsc := (*b.IncludedMessageHeaderResourcesReferencedByMessage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BundlePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*b.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*b.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*b.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*b.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*b.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*b.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*b.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *b.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*b.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BundlePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedCompositionResourcesReferencedByComposition != nil {
		for idx := range *b.IncludedCompositionResourcesReferencedByComposition {
			rsc := (*b.IncludedCompositionResourcesReferencedByComposition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedMessageHeaderResourcesReferencedByMessage != nil {
		for idx := range *b.IncludedMessageHeaderResourcesReferencedByMessage {
			rsc := (*b.IncludedMessageHeaderResourcesReferencedByMessage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*b.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*b.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*b.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*b.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*b.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*b.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*b.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *b.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*b.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
