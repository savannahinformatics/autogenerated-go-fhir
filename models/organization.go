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

type Organization struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                          `bson:"active,omitempty" json:"active,omitempty"`
	Type           []CodeableConcept              `bson:"type,omitempty" json:"type,omitempty"`
	Name           string                         `bson:"name,omitempty" json:"name,omitempty"`
	Alias          []string                       `bson:"alias,omitempty" json:"alias,omitempty"`
	Telecom        []ContactPoint                 `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                      `bson:"address,omitempty" json:"address,omitempty"`
	PartOf         *Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Contact        []OrganizationContactComponent `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint       []Reference                    `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Organization) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Organization"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "organization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type organization Organization

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Organization) UnmarshalJSON(data []byte) (err error) {
	x2 := organization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Organization(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Organization) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Organization"
	} else if x.ResourceType != "Organization" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Organization, instead received %s", x.ResourceType))
	}
	return nil
}

type OrganizationContactComponent struct {
	BackboneElement `bson:",inline"`
	Purpose         *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name            *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address         *Address         `bson:"address,omitempty" json:"address,omitempty"`
}

type OrganizationPlus struct {
	Organization                     `bson:",inline"`
	OrganizationPlusRelatedResources `bson:",inline"`
}

type OrganizationPlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByPartof                                 *[]Organization                  `bson:"_includedOrganizationResourcesReferencedByPartof,omitempty"`
	IncludedEndpointResourcesReferencedByEndpoint                                   *[]Endpoint                      `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo                        *[]Appointment                   `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedAccountResourcesReferencingOwner                                     *[]Account                       `bson:"_revIncludedAccountResourcesReferencingOwner,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                                   *[]Account                       `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedInvoiceResourcesReferencingIssuer                                    *[]Invoice                       `bson:"_revIncludedInvoiceResourcesReferencingIssuer,omitempty"`
	RevIncludedInvoiceResourcesReferencingParticipant                               *[]Invoice                       `bson:"_revIncludedInvoiceResourcesReferencingParticipant,omitempty"`
	RevIncludedInvoiceResourcesReferencingRecipient                                 *[]Invoice                       `bson:"_revIncludedInvoiceResourcesReferencingRecipient,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                         *[]EventDefinition               `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom                       *[]EventDefinition               `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor                       *[]EventDefinition               `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof                        *[]EventDefinition               `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                         *[]EventDefinition               `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                             *[]DocumentManifest              `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                           *[]DocumentManifest              `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref                       *[]DocumentManifest              `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient                        *[]DocumentManifest              `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                                      *[]Goal                          `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEndpointResourcesReferencingOrganization                             *[]Endpoint                      `bson:"_revIncludedEndpointResourcesReferencingOrganization,omitempty"`
	RevIncludedConsentResourcesReferencingData                                      *[]Consent                       `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                                     *[]Consent                       `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingOrganization                              *[]Consent                       `bson:"_revIncludedConsentResourcesReferencingOrganization,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                                 *[]Consent                       `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMedicationResourcesReferencingManufacturer                           *[]Medication                    `bson:"_revIncludedMedicationResourcesReferencingManufacturer,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                                 *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                               *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                               *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                                *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                            *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                            *[]Measure                       `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                         *[]DocumentReference             `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator                   *[]DocumentReference             `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingCustodian                       *[]DocumentReference             `bson:"_revIncludedDocumentReferenceResourcesReferencingCustodian,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                          *[]DocumentReference             `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer               *[]CoverageEligibilityResponse   `bson:"_revIncludedCoverageEligibilityResponseResourcesReferencingInsurer,omitempty"`
	RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor             *[]CoverageEligibilityResponse   `bson:"_revIncludedCoverageEligibilityResponseResourcesReferencingRequestor,omitempty"`
	RevIncludedMeasureReportResourcesReferencingReporter                            *[]MeasureReport                 `bson:"_revIncludedMeasureReportResourcesReferencingReporter,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource                   *[]MeasureReport                 `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingOrganization                     *[]PractitionerRole              `bson:"_revIncludedPractitionerRoleResourcesReferencingOrganization,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                          *[]ServiceRequest                `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                          *[]ServiceRequest                `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                           *[]SupplyRequest                 `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSubject                             *[]SupplyRequest                 `bson:"_revIncludedSupplyRequestResourcesReferencingSubject,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSupplier                            *[]SupplyRequest                 `bson:"_revIncludedSupplyRequestResourcesReferencingSupplier,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                         *[]VerificationResult            `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedPersonResourcesReferencingOrganization                               *[]Person                        `bson:"_revIncludedPersonResourcesReferencingOrganization,omitempty"`
	RevIncludedContractResourcesReferencingSubject                                  *[]Contract                      `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingAuthority                                *[]Contract                      `bson:"_revIncludedContractResourcesReferencingAuthority,omitempty"`
	RevIncludedContractResourcesReferencingSigner                                   *[]Contract                      `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedGroupResourcesReferencingManagingentity                              *[]Group                         `bson:"_revIncludedGroupResourcesReferencingManagingentity,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                             *[]PaymentNotice                 `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingProvider                            *[]PaymentNotice                 `bson:"_revIncludedPaymentNoticeResourcesReferencingProvider,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                            *[]PaymentNotice                 `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor                      *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom                    *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor                    *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof                     *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1                 *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2                 *[]ResearchDefinition            `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedOrganizationResourcesReferencingPartof                               *[]Organization                  `bson:"_revIncludedOrganizationResourcesReferencingPartof,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                              *[]CareTeam                      `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource                      *[]ImplementationGuide           `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                            *[]ImagingStudy                  `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformingorganization                 *[]ChargeItem                    `bson:"_revIncludedChargeItemResourcesReferencingPerformingorganization,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                                *[]ChargeItem                    `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformeractor                         *[]ChargeItem                    `bson:"_revIncludedChargeItemResourcesReferencingPerformeractor,omitempty"`
	RevIncludedChargeItemResourcesReferencingRequestingorganization                 *[]ChargeItem                    `bson:"_revIncludedChargeItemResourcesReferencingRequestingorganization,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor               *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom             *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor             *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof              *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1          *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2          *[]ResearchElementDefinition     `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingServiceprovider                         *[]Encounter                     `bson:"_revIncludedEncounterResourcesReferencingServiceprovider,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                              *[]Communication                 `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                             *[]Communication                 `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                              *[]Communication                 `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                           *[]Communication                 `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor                      *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom                    *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor                    *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof                     *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1                 *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2                 *[]ActivityDefinition            `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedInsurancePlanResourcesReferencingOwnedby                             *[]InsurancePlan                 `bson:"_revIncludedInsurancePlanResourcesReferencingOwnedby,omitempty"`
	RevIncludedInsurancePlanResourcesReferencingAdministeredby                      *[]InsurancePlan                 `bson:"_revIncludedInsurancePlanResourcesReferencingAdministeredby,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                                      *[]Linkage                       `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingAuthor                                    *[]Linkage                       `bson:"_revIncludedLinkageResourcesReferencingAuthor,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                                    *[]Linkage                       `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingRequester                           *[]DeviceRequest                 `bson:"_revIncludedDeviceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer                           *[]DeviceRequest                 `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                             *[]DeviceRequest                 `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest                        *[]DeviceRequest                 `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingReceiver                            *[]MessageHeader                 `bson:"_revIncludedMessageHeaderResourcesReferencingReceiver,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                               *[]MessageHeader                 `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingSender                              *[]MessageHeader                 `bson:"_revIncludedMessageHeaderResourcesReferencingSender,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingResponsible                         *[]MessageHeader                 `bson:"_revIncludedMessageHeaderResourcesReferencingResponsible,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation            *[]ImmunizationRecommendation    `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                                  *[]Provenance                    `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                                 *[]Provenance                    `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                                 *[]Provenance                    `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                                        *[]Task                          `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                                    *[]Task                          `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                                      *[]Task                          `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                                        *[]Task                          `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                                      *[]Task                          `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingCareteam                     *[]ExplanationOfBenefit          `bson:"_revIncludedExplanationOfBenefitResourcesReferencingCareteam,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee                        *[]ExplanationOfBenefit          `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingProvider                     *[]ExplanationOfBenefit          `bson:"_revIncludedExplanationOfBenefitResourcesReferencingProvider,omitempty"`
	RevIncludedResearchStudyResourcesReferencingSponsor                             *[]ResearchStudy                 `bson:"_revIncludedResearchStudyResourcesReferencingSponsor,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                                *[]CarePlan                      `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingOrganization                        *[]EpisodeOfCare                 `bson:"_revIncludedEpisodeOfCareResourcesReferencingOrganization,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                               *[]Procedure                     `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                         *[]List                          `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedImmunizationResourcesReferencingPerformer                            *[]Immunization                  `bson:"_revIncludedImmunizationResourcesReferencingPerformer,omitempty"`
	RevIncludedImmunizationResourcesReferencingManufacturer                         *[]Immunization                  `bson:"_revIncludedImmunizationResourcesReferencingManufacturer,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester                       *[]MedicationRequest             `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendeddispenser               *[]MedicationRequest             `bson:"_revIncludedMedicationRequestResourcesReferencingIntendeddispenser,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendedperformer               *[]MedicationRequest             `bson:"_revIncludedMedicationRequestResourcesReferencingIntendedperformer,omitempty"`
	RevIncludedDeviceResourcesReferencingOrganization                               *[]Device                        `bson:"_revIncludedDeviceResourcesReferencingOrganization,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                                    *[]Media                         `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor                        *[]EvidenceVariable              `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom                      *[]EvidenceVariable              `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor                      *[]EvidenceVariable              `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof                       *[]EvidenceVariable              `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson                        *[]EvidenceVariable              `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                                      *[]Flag                          `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                                       *[]Flag                          `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                                 *[]Observation                   `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                             *[]Observation                   `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                                 *[]Library                       `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                               *[]Library                       `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                               *[]Library                       `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                                *[]Library                       `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                                 *[]Library                       `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource                        *[]MedicationStatement           `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester                    *[]CommunicationRequest          `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon                      *[]CommunicationRequest          `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender                       *[]CommunicationRequest          `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient                    *[]CommunicationRequest          `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                                     *[]Basic                         `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                                      *[]Basic                         `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedClaimResponseResourcesReferencingInsurer                             *[]ClaimResponse                 `bson:"_revIncludedClaimResponseResourcesReferencingInsurer,omitempty"`
	RevIncludedClaimResponseResourcesReferencingRequestor                           *[]ClaimResponse                 `bson:"_revIncludedClaimResponseResourcesReferencingRequestor,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer                      *[]MedicationDispense            `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPerformer                        *[]DiagnosticReport              `bson:"_revIncludedDiagnosticReportResourcesReferencingPerformer,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter               *[]DiagnosticReport              `bson:"_revIncludedDiagnosticReportResourcesReferencingResultsinterpreter,omitempty"`
	RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization       *[]OrganizationAffiliation       `bson:"_revIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization,omitempty"`
	RevIncludedOrganizationAffiliationResourcesReferencingNetwork                   *[]OrganizationAffiliation       `bson:"_revIncludedOrganizationAffiliationResourcesReferencingNetwork,omitempty"`
	RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization *[]OrganizationAffiliation       `bson:"_revIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingOrganization                    *[]HealthcareService             `bson:"_revIncludedHealthcareServiceResourcesReferencingOrganization,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                                *[]Evidence                      `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                              *[]Evidence                      `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                              *[]Evidence                      `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                               *[]Evidence                      `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                                *[]Evidence                      `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                                  *[]AuditEvent                    `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingSource                                 *[]AuditEvent                    `bson:"_revIncludedAuditEventResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                                 *[]AuditEvent                    `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer               *[]PaymentReconciliation         `bson:"_revIncludedPaymentReconciliationResourcesReferencingPaymentissuer,omitempty"`
	RevIncludedPaymentReconciliationResourcesReferencingRequestor                   *[]PaymentReconciliation         `bson:"_revIncludedPaymentReconciliationResourcesReferencingRequestor,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                          *[]Condition                     `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                               *[]Composition                   `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                                *[]Composition                   `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                              *[]Composition                   `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                                 *[]Composition                   `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                          *[]DetectedIssue                 `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedMedicationKnowledgeResourcesReferencingManufacturer                  *[]MedicationKnowledge           `bson:"_revIncludedMedicationKnowledgeResourcesReferencingManufacturer,omitempty"`
	RevIncludedPatientResourcesReferencingGeneralpractitioner                       *[]Patient                       `bson:"_revIncludedPatientResourcesReferencingGeneralpractitioner,omitempty"`
	RevIncludedPatientResourcesReferencingOrganization                              *[]Patient                       `bson:"_revIncludedPatientResourcesReferencingOrganization,omitempty"`
	RevIncludedCoverageResourcesReferencingPayor                                    *[]Coverage                      `bson:"_revIncludedCoverageResourcesReferencingPayor,omitempty"`
	RevIncludedCoverageResourcesReferencingPolicyholder                             *[]Coverage                      `bson:"_revIncludedCoverageResourcesReferencingPolicyholder,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject                     *[]QuestionnaireResponse         `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor                      *[]QuestionnaireResponse         `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedCoverageEligibilityRequestResourcesReferencingProvider               *[]CoverageEligibilityRequest    `bson:"_revIncludedCoverageEligibilityRequestResourcesReferencingProvider,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingSupplier                           *[]SupplyDelivery                `bson:"_revIncludedSupplyDeliveryResourcesReferencingSupplier,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo                 *[]ClinicalImpression            `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                          *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom                        *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor                        *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                         *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1                     *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2                     *[]PlanDefinition                `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder              *[]MedicinalProductAuthorization `bson:"_revIncludedMedicinalProductAuthorizationResourcesReferencingHolder,omitempty"`
	RevIncludedClaimResourcesReferencingCareteam                                    *[]Claim                         `bson:"_revIncludedClaimResourcesReferencingCareteam,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                                       *[]Claim                         `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
	RevIncludedClaimResourcesReferencingProvider                                    *[]Claim                         `bson:"_revIncludedClaimResourcesReferencingProvider,omitempty"`
	RevIncludedClaimResourcesReferencingInsurer                                     *[]Claim                         `bson:"_revIncludedClaimResourcesReferencingInsurer,omitempty"`
	RevIncludedLocationResourcesReferencingOrganization                             *[]Location                      `bson:"_revIncludedLocationResourcesReferencingOrganization,omitempty"`
}

func (o *OrganizationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPartof() (organization *Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPartof == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*o.IncludedOrganizationResourcesReferencedByPartof))
	} else if len(*o.IncludedOrganizationResourcesReferencedByPartof) == 1 {
		organization = &(*o.IncludedOrganizationResourcesReferencedByPartof)[0]
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if o.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *o.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *o.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingOwner() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingOwner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if o.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *o.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingIssuer() (invoices []Invoice, err error) {
	if o.RevIncludedInvoiceResourcesReferencingIssuer == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *o.RevIncludedInvoiceResourcesReferencingIssuer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if o.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *o.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingRecipient() (invoices []Invoice, err error) {
	if o.RevIncludedInvoiceResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *o.RevIncludedInvoiceResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if o.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *o.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEndpointResourcesReferencingOrganization() (endpoints []Endpoint, err error) {
	if o.RevIncludedEndpointResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded endpoints not requested")
	} else {
		endpoints = *o.RevIncludedEndpointResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingOrganization() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationResourcesReferencingManufacturer() (medications []Medication, err error) {
	if o.RevIncludedMedicationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded medications not requested")
	} else {
		medications = *o.RevIncludedMedicationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingCustodian() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingCustodian
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageEligibilityResponseResourcesReferencingInsurer() (coverageEligibilityResponses []CoverageEligibilityResponse, err error) {
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded coverageEligibilityResponses not requested")
	} else {
		coverageEligibilityResponses = *o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageEligibilityResponseResourcesReferencingRequestor() (coverageEligibilityResponses []CoverageEligibilityResponse, err error) {
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded coverageEligibilityResponses not requested")
	} else {
		coverageEligibilityResponses = *o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingReporter() (measureReports []MeasureReport, err error) {
	if o.RevIncludedMeasureReportResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *o.RevIncludedMeasureReportResourcesReferencingReporter
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingOrganization() (practitionerRoles []PractitionerRole, err error) {
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *o.RevIncludedPractitionerRoleResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if o.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *o.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if o.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *o.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSubject() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSupplier() (supplyRequests []SupplyRequest, err error) {
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *o.RevIncludedSupplyRequestResourcesReferencingSupplier
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if o.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *o.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPersonResourcesReferencingOrganization() (people []Person, err error) {
	if o.RevIncludedPersonResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *o.RevIncludedPersonResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingAuthority() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingAuthority == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingAuthority
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedGroupResourcesReferencingManagingentity() (groups []Group, err error) {
	if o.RevIncludedGroupResourcesReferencingManagingentity == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *o.RevIncludedGroupResourcesReferencingManagingentity
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingProvider() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingProvider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationResourcesReferencingPartof() (organizations []Organization, err error) {
	if o.RevIncludedOrganizationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded organizations not requested")
	} else {
		organizations = *o.RevIncludedOrganizationResourcesReferencingPartof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if o.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *o.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if o.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *o.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformingorganization() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingPerformingorganization == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingPerformingorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingRequestingorganization() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingRequestingorganization == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingRequestingorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingServiceprovider() (encounters []Encounter, err error) {
	if o.RevIncludedEncounterResourcesReferencingServiceprovider == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *o.RevIncludedEncounterResourcesReferencingServiceprovider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedInsurancePlanResourcesReferencingOwnedby() (insurancePlans []InsurancePlan, err error) {
	if o.RevIncludedInsurancePlanResourcesReferencingOwnedby == nil {
		err = errors.New("RevIncluded insurancePlans not requested")
	} else {
		insurancePlans = *o.RevIncludedInsurancePlanResourcesReferencingOwnedby
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedInsurancePlanResourcesReferencingAdministeredby() (insurancePlans []InsurancePlan, err error) {
	if o.RevIncludedInsurancePlanResourcesReferencingAdministeredby == nil {
		err = errors.New("RevIncluded insurancePlans not requested")
	} else {
		insurancePlans = *o.RevIncludedInsurancePlanResourcesReferencingAdministeredby
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingAuthor() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingSender() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingSender == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingCareteam() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if o.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *o.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingProvider() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if o.RevIncludedExplanationOfBenefitResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *o.RevIncludedExplanationOfBenefitResourcesReferencingProvider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingSponsor() (researchStudies []ResearchStudy, err error) {
	if o.RevIncludedResearchStudyResourcesReferencingSponsor == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *o.RevIncludedResearchStudyResourcesReferencingSponsor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if o.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *o.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingOrganization() (episodeOfCares []EpisodeOfCare, err error) {
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPerformer() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingManufacturer() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if o.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *o.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendeddispenser() (medicationRequests []MedicationRequest, err error) {
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingOrganization() (devices []Device, err error) {
	if o.RevIncludedDeviceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *o.RevIncludedDeviceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if o.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *o.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if o.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *o.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if o.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *o.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingInsurer() (claimResponses []ClaimResponse, err error) {
	if o.RevIncludedClaimResponseResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *o.RevIncludedClaimResponseResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequestor() (claimResponses []ClaimResponse, err error) {
	if o.RevIncludedClaimResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *o.RevIncludedClaimResponseResourcesReferencingRequestor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if o.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *o.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResultsinterpreter() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization() (organizationAffiliations []OrganizationAffiliation, err error) {
	if o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingNetwork() (organizationAffiliations []OrganizationAffiliation, err error) {
	if o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization() (organizationAffiliations []OrganizationAffiliation, err error) {
	if o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingOrganization() (healthcareServices []HealthcareService, err error) {
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *o.RevIncludedHealthcareServiceResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingPaymentissuer() (paymentReconciliations []PaymentReconciliation, err error) {
	if o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequestor() (paymentReconciliations []PaymentReconciliation, err error) {
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *o.RevIncludedPaymentReconciliationResourcesReferencingRequestor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if o.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *o.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicationKnowledgeResourcesReferencingManufacturer() (medicationKnowledges []MedicationKnowledge, err error) {
	if o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer == nil {
		err = errors.New("RevIncluded medicationKnowledges not requested")
	} else {
		medicationKnowledges = *o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingGeneralpractitioner() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingGeneralpractitioner
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPatientResourcesReferencingOrganization() (patients []Patient, err error) {
	if o.RevIncludedPatientResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *o.RevIncludedPatientResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPayor() (coverages []Coverage, err error) {
	if o.RevIncludedCoverageResourcesReferencingPayor == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *o.RevIncludedCoverageResourcesReferencingPayor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPolicyholder() (coverages []Coverage, err error) {
	if o.RevIncludedCoverageResourcesReferencingPolicyholder == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *o.RevIncludedCoverageResourcesReferencingPolicyholder
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingProvider() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingSupplier() (supplyDeliveries []SupplyDelivery, err error) {
	if o.RevIncludedSupplyDeliveryResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *o.RevIncludedSupplyDeliveryResourcesReferencingSupplier
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedMedicinalProductAuthorizationResourcesReferencingHolder() (medicinalProductAuthorizations []MedicinalProductAuthorization, err error) {
	if o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder == nil {
		err = errors.New("RevIncluded medicinalProductAuthorizations not requested")
	} else {
		medicinalProductAuthorizations = *o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingCareteam() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingCareteam
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingProvider() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingProvider
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingInsurer() (claims []Claim, err error) {
	if o.RevIncludedClaimResourcesReferencingInsurer == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *o.RevIncludedClaimResourcesReferencingInsurer
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedLocationResourcesReferencingOrganization() (locations []Location, err error) {
	if o.RevIncludedLocationResourcesReferencingOrganization == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *o.RevIncludedLocationResourcesReferencingOrganization
	}
	return
}

func (o *OrganizationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *o.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*o.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingIssuer != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingIssuer {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingIssuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
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
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEndpointResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEndpointResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEndpointResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingActor {
			rsc := (*o.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingOrganization {
			rsc := (*o.RevIncludedConsentResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*o.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
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
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerRoleResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerRoleResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAuthority != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAuthority {
			rsc := (*o.RevIncludedContractResourcesReferencingAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *o.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*o.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
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
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*o.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingPerformingorganization != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingPerformingorganization {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingPerformingorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingRequestingorganization != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingRequestingorganization {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingRequestingorganization)[idx]
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
	if o.RevIncludedEncounterResourcesReferencingServiceprovider != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingServiceprovider {
			rsc := (*o.RevIncludedEncounterResourcesReferencingServiceprovider)[idx]
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
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
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
	if o.RevIncludedInsurancePlanResourcesReferencingOwnedby != nil {
		for idx := range *o.RevIncludedInsurancePlanResourcesReferencingOwnedby {
			rsc := (*o.RevIncludedInsurancePlanResourcesReferencingOwnedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInsurancePlanResourcesReferencingAdministeredby != nil {
		for idx := range *o.RevIncludedInsurancePlanResourcesReferencingAdministeredby {
			rsc := (*o.RevIncludedInsurancePlanResourcesReferencingAdministeredby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*o.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
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
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
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
	if o.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*o.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*o.RevIncludedTaskResourcesReferencingRequester)[idx]
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
	if o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchStudyResourcesReferencingSponsor != nil {
		for idx := range *o.RevIncludedResearchStudyResourcesReferencingSponsor {
			rsc := (*o.RevIncludedResearchStudyResourcesReferencingSponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *o.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*o.RevIncludedMediaResourcesReferencingOperator)[idx]
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
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
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
	if o.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*o.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*o.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
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
	if o.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
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
	if o.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
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
	if o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*o.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
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
	if o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder != nil {
		for idx := range *o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder {
			rsc := (*o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*o.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*o.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*o.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (o *OrganizationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedOrganizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPartof {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *o.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*o.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingOwner {
			rsc := (*o.RevIncludedAccountResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*o.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingIssuer != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingIssuer {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingIssuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*o.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
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
	if o.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*o.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEndpointResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEndpointResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEndpointResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingActor {
			rsc := (*o.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingOrganization {
			rsc := (*o.RevIncludedConsentResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*o.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationResourcesReferencingManufacturer)[idx]
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
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingCustodian != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingCustodian {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingCustodian)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedCoverageEligibilityResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*o.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPractitionerRoleResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPractitionerRoleResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPractitionerRoleResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPersonResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPersonResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPersonResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingAuthority != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingAuthority {
			rsc := (*o.RevIncludedContractResourcesReferencingAuthority)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSigner {
			rsc := (*o.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *o.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*o.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
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
	if o.RevIncludedOrganizationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedOrganizationResourcesReferencingPartof {
			rsc := (*o.RevIncludedOrganizationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *o.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*o.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*o.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingPerformingorganization != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingPerformingorganization {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingPerformingorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingRequestingorganization != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingRequestingorganization {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingRequestingorganization)[idx]
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
	if o.RevIncludedEncounterResourcesReferencingServiceprovider != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingServiceprovider {
			rsc := (*o.RevIncludedEncounterResourcesReferencingServiceprovider)[idx]
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
	if o.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
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
	if o.RevIncludedInsurancePlanResourcesReferencingOwnedby != nil {
		for idx := range *o.RevIncludedInsurancePlanResourcesReferencingOwnedby {
			rsc := (*o.RevIncludedInsurancePlanResourcesReferencingOwnedby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedInsurancePlanResourcesReferencingAdministeredby != nil {
		for idx := range *o.RevIncludedInsurancePlanResourcesReferencingAdministeredby {
			rsc := (*o.RevIncludedInsurancePlanResourcesReferencingAdministeredby)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*o.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
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
	if o.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingAgent)[idx]
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
	if o.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*o.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*o.RevIncludedTaskResourcesReferencingRequester)[idx]
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
	if o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*o.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedResearchStudyResourcesReferencingSponsor != nil {
		for idx := range *o.RevIncludedResearchStudyResourcesReferencingSponsor {
			rsc := (*o.RevIncludedResearchStudyResourcesReferencingSponsor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*o.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedEpisodeOfCareResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedEpisodeOfCareResourcesReferencingOrganization {
			rsc := (*o.RevIncludedEpisodeOfCareResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendeddispenser)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*o.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDeviceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedDeviceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedDeviceResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *o.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*o.RevIncludedMediaResourcesReferencingOperator)[idx]
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
	if o.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*o.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*o.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*o.RevIncludedObservationResourcesReferencingPerformer)[idx]
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
	if o.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*o.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*o.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*o.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingPrimaryorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingNetwork)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization != nil {
		for idx := range *o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization {
			rsc := (*o.RevIncludedOrganizationAffiliationResourcesReferencingParticipatingorganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedHealthcareServiceResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedHealthcareServiceResourcesReferencingOrganization {
			rsc := (*o.RevIncludedHealthcareServiceResourcesReferencingOrganization)[idx]
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
	if o.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingPaymentissuer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *o.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*o.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
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
	if o.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*o.RevIncludedCompositionResourcesReferencingAttester)[idx]
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
	if o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer != nil {
		for idx := range *o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer {
			rsc := (*o.RevIncludedMedicationKnowledgeResourcesReferencingManufacturer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*o.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedPatientResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedPatientResourcesReferencingOrganization {
			rsc := (*o.RevIncludedPatientResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *o.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*o.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*o.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *o.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*o.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
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
	if o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder != nil {
		for idx := range *o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder {
			rsc := (*o.RevIncludedMedicinalProductAuthorizationResourcesReferencingHolder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*o.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*o.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*o.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedClaimResourcesReferencingInsurer != nil {
		for idx := range *o.RevIncludedClaimResourcesReferencingInsurer {
			rsc := (*o.RevIncludedClaimResourcesReferencingInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if o.RevIncludedLocationResourcesReferencingOrganization != nil {
		for idx := range *o.RevIncludedLocationResourcesReferencingOrganization {
			rsc := (*o.RevIncludedLocationResourcesReferencingOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
