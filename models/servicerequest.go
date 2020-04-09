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

type ServiceRequest struct {
	DomainResource          `bson:",inline"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical   []canonical       `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string          `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces                []Reference       `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Requisition             *Identifier       `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status                  string            `bson:"status,omitempty" json:"status,omitempty"`
	Intent                  string            `bson:"intent,omitempty" json:"intent,omitempty"`
	Category                []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Priority                string            `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform            *bool             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Code                    *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	OrderDetail             []CodeableConcept `bson:"orderDetail,omitempty" json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity         `bson:"quantityQuantity,omitempty" json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio            `bson:"quantityRatio,omitempty" json:"quantityRatio,omitempty"`
	QuantityRange           *Range            `bson:"quantityRange,omitempty" json:"quantityRange,omitempty"`
	Subject                 *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter               *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime      *FHIRDateTime     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period           `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing           `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *FHIRDateTime     `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester               *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType           *CodeableConcept  `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer               []Reference       `bson:"performer,omitempty" json:"performer,omitempty"`
	LocationCode            []CodeableConcept `bson:"locationCode,omitempty" json:"locationCode,omitempty"`
	LocationReference       []Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	ReasonCode              []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference         []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Insurance               []Reference       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	SupportingInfo          []Reference       `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Specimen                []Reference       `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite                []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note                    []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	PatientInstruction      string            `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference       `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *ServiceRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ServiceRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "serviceRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type serviceRequest ServiceRequest

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ServiceRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := serviceRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ServiceRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ServiceRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ServiceRequest"
	} else if x.ResourceType != "ServiceRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ServiceRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type ServiceRequestPlus struct {
	ServiceRequest                     `bson:",inline"`
	ServiceRequestPlusRelatedResources `bson:",inline"`
}

type ServiceRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByRequester                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByRequester,omitempty"`
	IncludedOrganizationResourcesReferencedByRequester                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByRequester,omitempty"`
	IncludedDeviceResourcesReferencedByRequester                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByRequester,omitempty"`
	IncludedPatientResourcesReferencedByRequester                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByRequester,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByRequester                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByRequester,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRequester                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByRequester,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                         *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByPerformer                *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedServiceRequestResourcesReferencedByReplaces                    *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByReplaces,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                             *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                           *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical       *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical   *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedMedicationRequestResourcesReferencedByBasedon                  *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByBasedon,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimen                          *[]Specimen                   `bson:"_includedSpecimenResourcesReferencedBySpecimen,omitempty"`
	RevIncludedAppointmentResourcesReferencingBasedon                      *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingBasedon,omitempty"`
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
	RevIncludedServiceRequestResourcesReferencingReplaces                  *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingReplaces,omitempty"`
	RevIncludedServiceRequestResourcesReferencingBasedon                   *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingBasedon,omitempty"`
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
	RevIncludedImagingStudyResourcesReferencingBasedon                     *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingBasedon                        *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingBasedon,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingActivityreference               *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral           *[]EpisodeOfCare              `bson:"_revIncludedEpisodeOfCareResourcesReferencingIncomingreferral,omitempty"`
	RevIncludedProcedureResourcesReferencingBasedon                        *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedMediaResourcesReferencingBasedon                            *[]Media                      `bson:"_revIncludedMediaResourcesReferencingBasedon,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
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
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequester() (practitioner *Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedByRequester == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedPractitionerResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedPractitionerResourcesReferencedByRequester))
	} else if len(*s.IncludedPractitionerResourcesReferencedByRequester) == 1 {
		practitioner = &(*s.IncludedPractitionerResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequester() (organization *Organization, err error) {
	if s.IncludedOrganizationResourcesReferencedByRequester == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*s.IncludedOrganizationResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*s.IncludedOrganizationResourcesReferencedByRequester))
	} else if len(*s.IncludedOrganizationResourcesReferencedByRequester) == 1 {
		organization = &(*s.IncludedOrganizationResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedByRequester() (device *Device, err error) {
	if s.IncludedDeviceResourcesReferencedByRequester == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedDeviceResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedDeviceResourcesReferencedByRequester))
	} else if len(*s.IncludedDeviceResourcesReferencedByRequester) == 1 {
		device = &(*s.IncludedDeviceResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByRequester() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedByRequester == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedByRequester))
	} else if len(*s.IncludedPatientResourcesReferencedByRequester) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByRequester() (practitionerRole *PractitionerRole, err error) {
	if s.IncludedPractitionerRoleResourcesReferencedByRequester == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*s.IncludedPractitionerRoleResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*s.IncludedPractitionerRoleResourcesReferencedByRequester))
	} else if len(*s.IncludedPractitionerRoleResourcesReferencedByRequester) == 1 {
		practitionerRole = &(*s.IncludedPractitionerRoleResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByRequester() (relatedPerson *RelatedPerson, err error) {
	if s.IncludedRelatedPersonResourcesReferencedByRequester == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*s.IncludedRelatedPersonResourcesReferencedByRequester) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*s.IncludedRelatedPersonResourcesReferencedByRequester))
	} else if len(*s.IncludedRelatedPersonResourcesReferencedByRequester) == 1 {
		relatedPerson = &(*s.IncludedRelatedPersonResourcesReferencedByRequester)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *s.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if s.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *s.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByPerformer() (careTeams []CareTeam, err error) {
	if s.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *s.IncludedCareTeamResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedDeviceResourcesReferencedByPerformer() (devices []Device, err error) {
	if s.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *s.IncludedDeviceResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if s.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *s.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByPerformer() (healthcareServices []HealthcareService, err error) {
	if s.IncludedHealthcareServiceResourcesReferencedByPerformer == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *s.IncludedHealthcareServiceResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByPerformer() (practitionerRoles []PractitionerRole, err error) {
	if s.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *s.IncludedPractitionerRoleResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if s.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *s.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByReplaces() (serviceRequests []ServiceRequest, err error) {
	if s.IncludedServiceRequestResourcesReferencedByReplaces == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *s.IncludedServiceRequestResourcesReferencedByReplaces
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if s.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*s.IncludedGroupResourcesReferencedBySubject))
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*s.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if s.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*s.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedBySubject))
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if s.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*s.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*s.IncludedLocationResourcesReferencedBySubject))
	} else if len(*s.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*s.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if s.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*s.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*s.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*s.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*s.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if s.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *s.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedMedicationRequestResourcesReferencedByBasedon() (medicationRequests []MedicationRequest, err error) {
	if s.IncludedMedicationRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included medicationRequests not requested")
	} else {
		medicationRequests = *s.IncludedMedicationRequestResourcesReferencedByBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if s.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *s.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedByPatient))
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimen() (specimen []Specimen, err error) {
	if s.IncludedSpecimenResourcesReferencedBySpecimen == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *s.IncludedSpecimenResourcesReferencedBySpecimen
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingBasedon() (appointments []Appointment, err error) {
	if s.RevIncludedAppointmentResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *s.RevIncludedAppointmentResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *s.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingReplaces() (serviceRequests []ServiceRequest, err error) {
	if s.RevIncludedServiceRequestResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *s.RevIncludedServiceRequestResourcesReferencingReplaces
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingBasedon() (serviceRequests []ServiceRequest, err error) {
	if s.RevIncludedServiceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *s.RevIncludedServiceRequestResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if s.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *s.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if s.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *s.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if s.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *s.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingBasedon() (encounters []Encounter, err error) {
	if s.RevIncludedEncounterResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *s.RevIncludedEncounterResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if s.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *s.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingIncomingreferral() (episodeOfCares []EpisodeOfCare, err error) {
	if s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingBasedon() (procedures []Procedure, err error) {
	if s.RevIncludedProcedureResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *s.RevIncludedProcedureResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedMediaResourcesReferencingBasedon() (media []Media, err error) {
	if s.RevIncludedMediaResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *s.RevIncludedMediaResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *s.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if s.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *s.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if s.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *s.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*s.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*s.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*s.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByRequester {
			rsc := (*s.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByRequester {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*s.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*s.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*s.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*s.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*s.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*s.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedServiceRequestResourcesReferencedByReplaces != nil {
		for idx := range *s.IncludedServiceRequestResourcesReferencedByReplaces {
			rsc := (*s.IncludedServiceRequestResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedGroupResourcesReferencedBySubject {
			rsc := (*s.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*s.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedPatientResourcesReferencedBySubject {
			rsc := (*s.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedLocationResourcesReferencedBySubject {
			rsc := (*s.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *s.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*s.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*s.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*s.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*s.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPatient {
			rsc := (*s.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *s.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*s.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *ServiceRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedAppointmentResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingBasedon {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingReplaces != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingReplaces {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*s.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEncounterResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedEncounterResourcesReferencingBasedon {
			rsc := (*s.RevIncludedEncounterResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *s.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*s.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for idx := range *s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			rsc := (*s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*s.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMediaResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedMediaResourcesReferencingBasedon {
			rsc := (*s.RevIncludedMediaResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *ServiceRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedPractitionerResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByRequester {
			rsc := (*s.IncludedPractitionerResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedOrganizationResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedOrganizationResourcesReferencedByRequester {
			rsc := (*s.IncludedOrganizationResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByRequester {
			rsc := (*s.IncludedDeviceResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByRequester {
			rsc := (*s.IncludedPatientResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByRequester {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByRequester != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByRequester {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*s.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*s.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*s.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*s.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*s.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*s.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *s.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*s.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedServiceRequestResourcesReferencedByReplaces != nil {
		for idx := range *s.IncludedServiceRequestResourcesReferencedByReplaces {
			rsc := (*s.IncludedServiceRequestResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedGroupResourcesReferencedBySubject {
			rsc := (*s.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*s.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedPatientResourcesReferencedBySubject {
			rsc := (*s.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedLocationResourcesReferencedBySubject {
			rsc := (*s.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*s.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*s.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *s.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*s.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*s.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*s.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *s.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*s.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPatient {
			rsc := (*s.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *s.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*s.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAppointmentResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingBasedon {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingReplaces != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingReplaces {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*s.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEncounterResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedEncounterResourcesReferencingBasedon {
			rsc := (*s.RevIncludedEncounterResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *s.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*s.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral != nil {
		for idx := range *s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral {
			rsc := (*s.RevIncludedEpisodeOfCareResourcesReferencingIncomingreferral)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*s.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMediaResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedMediaResourcesReferencingBasedon {
			rsc := (*s.RevIncludedMediaResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
