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

type PractitionerRole struct {
	DomainResource         `bson:",inline"`
	Identifier             []Identifier                             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                                    `bson:"active,omitempty" json:"active,omitempty"`
	Period                 *Period                                  `bson:"period,omitempty" json:"period,omitempty"`
	Practitioner           *Reference                               `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Organization           *Reference                               `bson:"organization,omitempty" json:"organization,omitempty"`
	Code                   []CodeableConcept                        `bson:"code,omitempty" json:"code,omitempty"`
	Specialty              []CodeableConcept                        `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                              `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService      []Reference                              `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                           `bson:"telecom,omitempty" json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTimeComponent `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailableComponent  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions string                                   `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                              `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *PractitionerRole) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "PractitionerRole"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "practitionerRole" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type practitionerRole PractitionerRole

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *PractitionerRole) UnmarshalJSON(data []byte) (err error) {
	x2 := practitionerRole{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = PractitionerRole(x2)
		return x.checkResourceType()
	}
	return
}

func (x *PractitionerRole) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "PractitionerRole"
	} else if x.ResourceType != "PractitionerRole" {
		return errors.New(fmt.Sprintf("Expected resourceType to be PractitionerRole, instead received %s", x.ResourceType))
	}
	return nil
}

type PractitionerRoleAvailableTimeComponent struct {
	BackboneElement    `bson:",inline"`
	DaysOfWeek         []string      `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *FHIRDateTime `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *FHIRDateTime `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}

type PractitionerRoleNotAvailableComponent struct {
	BackboneElement `bson:",inline"`
	Description     string  `bson:"description,omitempty" json:"description,omitempty"`
	During          *Period `bson:"during,omitempty" json:"during,omitempty"`
}

type PractitionerRolePlus struct {
	PractitionerRole                     `bson:",inline"`
	PractitionerRolePlusRelatedResources `bson:",inline"`
}

type PractitionerRolePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPractitioner                  *[]Practitioner                `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedEndpointResourcesReferencedByEndpoint                          *[]Endpoint                    `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByService                  *[]HealthcareService           `bson:"_includedHealthcareServiceResourcesReferencedByService,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization                `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                    `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                          *[]Account                     `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedInvoiceResourcesReferencingParticipant                      *[]Invoice                     `bson:"_revIncludedInvoiceResourcesReferencingParticipant,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                  *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient               *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                            *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                        *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator          *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                 *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor    *[]CoverageEligibilityResponse `bson:"_revIncludedCoverageEligibilityResponseResourcesReferencingRequestor,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingReporter                   *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingReporter,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                  *[]SupplyRequest               `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult          `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                    `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingSigner                          *[]Contract                    `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer                 *[]RiskAssessment              `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedGroupResourcesReferencingManagingentity                     *[]Group                       `bson:"_revIncludedGroupResourcesReferencingManagingentity,omitempty"`
	RevIncludedGroupResourcesReferencingMember                             *[]Group                       `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice               `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingProvider                   *[]PaymentNotice               `bson:"_revIncludedPaymentNoticeResourcesReferencingProvider,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice               `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                     *[]CareTeam                    `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide         `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                   *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedImagingStudyResourcesReferencingInterpreter                 *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingInterpreter,omitempty"`
	RevIncludedImagingStudyResourcesReferencingReferrer                    *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingReferrer,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                       *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformeractor                *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingPerformeractor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingParticipant                    *[]Encounter                   `bson:"_revIncludedEncounterResourcesReferencingParticipant,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                     *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                  *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                     `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingAuthor                           *[]Linkage                     `bson:"_revIncludedLinkageResourcesReferencingAuthor,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                     `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedRequestGroupResourcesReferencingAuthor                      *[]RequestGroup                `bson:"_revIncludedRequestGroupResourcesReferencingAuthor,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant                 *[]RequestGroup                `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingRequester                  *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer                  *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver                   *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingAuthor                     *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingAuthor,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingSender                     *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingSender,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible                *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingEnterer                    *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingEnterer,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation  `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                         *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                               *[]Task                        `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                           *[]Task                        `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                        `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                        `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                        `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingCareteam            *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingCareteam,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee               *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingProvider            *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingProvider,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingEnterer             *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingEnterer,omitempty"`
	RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator      *[]ResearchStudy               `bson:"_revIncludedResearchStudyResourcesReferencingPrincipalinvestigator,omitempty"`
	RevIncludedSpecimenResourcesReferencingCollector                       *[]Specimen                    `bson:"_revIncludedSpecimenResourcesReferencingCollector,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder              *[]AllergyIntolerance          `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingAsserter              *[]AllergyIntolerance          `bson:"_revIncludedAllergyIntoleranceResourcesReferencingAsserter,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                       *[]CarePlan                    `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                      *[]Procedure                   `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                        `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSource                              *[]List                        `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedImmunizationResourcesReferencingPerformer                   *[]Immunization                `bson:"_revIncludedImmunizationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester              *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendedperformer      *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingIntendedperformer,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPrescriber            *[]VisionPrescription          `bson:"_revIncludedVisionPrescriptionResourcesReferencingPrescriber,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                       `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                           *[]Media                       `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                              *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse         `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder                    *[]AdverseEvent                `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                    *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer       *[]MedicationAdministration    `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource               *[]MedicationStatement         `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester           *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender              *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                       `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                             *[]Basic                       `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedClaimResponseResourcesReferencingRequestor                  *[]ClaimResponse               `bson:"_revIncludedClaimResponseResourcesReferencingRequestor,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer             *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingResponsibleparty      *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingResponsibleparty,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer               *[]DiagnosticReport            `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter      *[]DiagnosticReport            `bson:"_revIncludedDiagnosticReportResourcesReferencingResultsinterpreter,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingProvider                  *[]NutritionOrder              `bson:"_revIncludedNutritionOrderResourcesReferencingProvider,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                         *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingSource                        *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingRequestor          *[]PaymentReconciliation       `bson:"_revIncludedPaymentReconciliationResourcesReferencingRequestor,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                       *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                       *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                     *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor                     *[]DetectedIssue               `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue               `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingGeneralpractitioner              *[]Patient                     `bson:"_revIncludedPatientResourcesReferencingGeneralpractitioner,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor             *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource             *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedCoverageEligibilityRequestResourcesReferencingProvider      *[]CoverageEligibilityRequest  `bson:"_revIncludedCoverageEligibilityRequestResourcesReferencingProvider,omitempty"`
	RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer       *[]CoverageEligibilityRequest  `bson:"_revIncludedCoverageEligibilityRequestResourcesReferencingEnterer,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                    `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingReceiver                  *[]SupplyDelivery              `bson:"_revIncludedSupplyDeliveryResourcesReferencingReceiver,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingSupplier                  *[]SupplyDelivery              `bson:"_revIncludedSupplyDeliveryResourcesReferencingSupplier,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAssessor              *[]ClinicalImpression          `bson:"_revIncludedClinicalImpressionResourcesReferencingAssessor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression          `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingCareteam                           *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingCareteam,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                              *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
	RevIncludedClaimResourcesReferencingProvider                           *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingProvider,omitempty"`
	RevIncludedClaimResourcesReferencingEnterer                            *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingEnterer,omitempty"`
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedPractitionerResourceReferencedByPractitioner() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByPractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResourcesReferencedByPractitioner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResourcesReferencedByPractitioner))
	} else if len(*p.IncludedPractitionerResourcesReferencedByPractitioner) == 1 {
		practitioner = &(*p.IncludedPractitionerResourcesReferencedByPractitioner)[0]
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if p.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *p.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByService() (healthcareServices []HealthcareService, err error) {
	if p.IncludedHealthcareServiceResourcesReferencedByService == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *p.IncludedHealthcareServiceResourcesReferencedByService
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedLocationResourcesReferencedByLocation() (locations []Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *p.IncludedLocationResourcesReferencedByLocation
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCoverageEligibilityResponseResourcesReferencingRequestor() (coverageEligibilityResponses []CoverageEligibilityResponse, err error) {
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded coverageEligibilityResponses not requested")
	} else {
		coverageEligibilityResponses = *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingReporter() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingReporter
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if p.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *p.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedGroupResourcesReferencingManagingentity() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingManagingentity == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingManagingentity
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingProvider() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingProvider
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingInterpreter() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingInterpreter
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingReferrer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingReferrer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingReferrer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingParticipant() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingAuthor() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingAuthor() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingAuthor() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingSender() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingSender == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingSender
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingEnterer() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingCareteam() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingProvider() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingEnterer() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingPrincipalinvestigator() (researchStudies []ResearchStudy, err error) {
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingCollector() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingCollector == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingCollector
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPerformer() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingPrescriber() (visionPrescriptions []VisionPrescription, err error) {
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequestor() (claimResponses []ClaimResponse, err error) {
	if p.RevIncludedClaimResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *p.RevIncludedClaimResponseResourcesReferencingRequestor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingResponsibleparty() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResultsinterpreter() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingProvider() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingProvider
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequestor() (paymentReconciliations []PaymentReconciliation, err error) {
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingAuthor() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPatientResourcesReferencingGeneralpractitioner() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingGeneralpractitioner
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingProvider() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingEnterer() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if p.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *p.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingReceiver() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingSupplier() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAssessor() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingAssessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClaimResourcesReferencingCareteam() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingCareteam
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClaimResourcesReferencingProvider() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingProvider
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedClaimResourcesReferencingEnterer() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingEnterer
	}
	return
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *p.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*p.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByService != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByService {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PractitionerRolePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActor {
			rsc := (*p.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*p.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingInterpreter {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingInterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingReferrer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingReferrer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingReferrer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PractitionerRolePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *p.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*p.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByService != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByService {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActor {
			rsc := (*p.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*p.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingInterpreter {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingInterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingReferrer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingReferrer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingReferrer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
