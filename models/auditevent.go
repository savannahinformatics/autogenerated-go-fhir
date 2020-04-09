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

type AuditEvent struct {
	DomainResource `bson:",inline"`
	Type           *Coding                     `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []Coding                    `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string                      `bson:"action,omitempty" json:"action,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome        string                      `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string                      `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []CodeableConcept           `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent          []AuditEventAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source         *AuditEventSourceComponent  `bson:"source,omitempty" json:"source,omitempty"`
	Entity         []AuditEventEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *AuditEvent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AuditEvent"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "auditEvent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type auditEvent AuditEvent

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *AuditEvent) UnmarshalJSON(data []byte) (err error) {
	x2 := auditEvent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AuditEvent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AuditEvent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AuditEvent"
	} else if x.ResourceType != "AuditEvent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AuditEvent, instead received %s", x.ResourceType))
	}
	return nil
}

type AuditEventAgentComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept                 `bson:"type,omitempty" json:"type,omitempty"`
	Role            []CodeableConcept                `bson:"role,omitempty" json:"role,omitempty"`
	Who             *Reference                       `bson:"who,omitempty" json:"who,omitempty"`
	AltId           string                           `bson:"altId,omitempty" json:"altId,omitempty"`
	Name            string                           `bson:"name,omitempty" json:"name,omitempty"`
	Requestor       *bool                            `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location        *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	Policy          []string                         `bson:"policy,omitempty" json:"policy,omitempty"`
	Media           *Coding                          `bson:"media,omitempty" json:"media,omitempty"`
	Network         *AuditEventAgentNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse    []CodeableConcept                `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventAgentNetworkComponent struct {
	BackboneElement `bson:",inline"`
	Address         string `bson:"address,omitempty" json:"address,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	BackboneElement `bson:",inline"`
	Site            string     `bson:"site,omitempty" json:"site,omitempty"`
	Observer        *Reference `bson:"observer,omitempty" json:"observer,omitempty"`
	Type            []Coding   `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventEntityComponent struct {
	BackboneElement `bson:",inline"`
	What            *Reference                        `bson:"what,omitempty" json:"what,omitempty"`
	Type            *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	Role            *Coding                           `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle       *Coding                           `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel   []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name            string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query           string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail          []AuditEventEntityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventEntityDetailComponent struct {
	BackboneElement   `bson:",inline"`
	Type              string `bson:"type,omitempty" json:"type,omitempty"`
	ValueString       string `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBase64Binary string `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
}

type AuditEventPlus struct {
	AuditEvent                     `bson:",inline"`
	AuditEventPlusRelatedResources `bson:",inline"`
}

type AuditEventPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAgent                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAgent,omitempty"`
	IncludedOrganizationResourcesReferencedByAgent                         *[]Organization               `bson:"_includedOrganizationResourcesReferencedByAgent,omitempty"`
	IncludedDeviceResourcesReferencedByAgent                               *[]Device                     `bson:"_includedDeviceResourcesReferencedByAgent,omitempty"`
	IncludedPatientResourcesReferencedByAgent                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByAgent,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByAgent                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByAgent,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAgent                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByAgent,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedOrganizationResourcesReferencedBySource                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedBySource,omitempty"`
	IncludedDeviceResourcesReferencedBySource                              *[]Device                     `bson:"_includedDeviceResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                             *[]Patient                    `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedPractitionerRoleResourcesReferencedBySource                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedBySource,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySource                       *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath1                       *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatientPath1,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath2                       *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatientPath2,omitempty"`
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

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAgent() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByAgent == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByAgent))
	} else if len(*a.IncludedPractitionerResourcesReferencedByAgent) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAgent() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedByAgent == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedByAgent))
	} else if len(*a.IncludedOrganizationResourcesReferencedByAgent) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedDeviceResourceReferencedByAgent() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedByAgent == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedByAgent))
	} else if len(*a.IncludedDeviceResourcesReferencedByAgent) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByAgent() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByAgent == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByAgent))
	} else if len(*a.IncludedPatientResourcesReferencedByAgent) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByAgent() (practitionerRole *PractitionerRole, err error) {
	if a.IncludedPractitionerRoleResourcesReferencedByAgent == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*a.IncludedPractitionerRoleResourcesReferencedByAgent))
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByAgent) == 1 {
		practitionerRole = &(*a.IncludedPractitionerRoleResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAgent() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByAgent == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAgent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByAgent))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByAgent) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByAgent)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*a.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySource() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedBySource == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedBySource))
	} else if len(*a.IncludedOrganizationResourcesReferencedBySource) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedDeviceResourceReferencedBySource() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedBySource == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedBySource))
	} else if len(*a.IncludedDeviceResourcesReferencedBySource) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedBySource))
	} else if len(*a.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedBySource() (practitionerRole *PractitionerRole, err error) {
	if a.IncludedPractitionerRoleResourcesReferencedBySource == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*a.IncludedPractitionerRoleResourcesReferencedBySource))
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedBySource) == 1 {
		practitionerRole = &(*a.IncludedPractitionerRoleResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath1() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath1 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath1))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath1)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath2() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath2 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath2))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath2)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *a.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if a.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *a.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if a.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *a.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *a.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if a.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *a.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*a.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*a.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAgent {
			rsc := (*a.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*a.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*a.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySource != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySource {
			rsc := (*a.IncludedDeviceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySource {
			rsc := (*a.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedByAgent {
			rsc := (*a.IncludedOrganizationResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByAgent {
			rsc := (*a.IncludedDeviceResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByAgent {
			rsc := (*a.IncludedPatientResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByAgent {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByAgent != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByAgent {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*a.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *a.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*a.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedBySource != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedBySource {
			rsc := (*a.IncludedDeviceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPatientResourcesReferencedBySource {
			rsc := (*a.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			rsc := (*a.IncludedPatientResourcesReferencedByPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
