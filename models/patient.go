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

type Patient struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                           `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName                     `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint                  `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               string                          `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *FHIRDateTime                   `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                           `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *FHIRDateTime                   `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address                       `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept                `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                           `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int32                          `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment                    `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContactComponent       `bson:"contact,omitempty" json:"contact,omitempty"`
	Communication        []PatientCommunicationComponent `bson:"communication,omitempty" json:"communication,omitempty"`
	GeneralPractitioner  []Reference                     `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference                      `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLinkComponent          `bson:"link,omitempty" json:"link,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Patient) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Patient"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "patient" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type patient Patient

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Patient) UnmarshalJSON(data []byte) (err error) {
	x2 := patient{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Patient(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Patient) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Patient"
	} else if x.ResourceType != "Patient" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Patient, instead received %s", x.ResourceType))
	}
	return nil
}

type PatientContactComponent struct {
	BackboneElement `bson:",inline"`
	Relationship    []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name            *HumanName        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom         []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address         *Address          `bson:"address,omitempty" json:"address,omitempty"`
	Gender          string            `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization    *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

type PatientCommunicationComponent struct {
	BackboneElement `bson:",inline"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred       *bool            `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PatientLinkComponent struct {
	BackboneElement `bson:",inline"`
	Other           *Reference `bson:"other,omitempty" json:"other,omitempty"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
}

type PatientPlus struct {
	Patient                     `bson:",inline"`
	PatientPlusRelatedResources `bson:",inline"`
}

type PatientPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByGeneralpractitioner           *[]Practitioner                `bson:"_includedPractitionerResourcesReferencedByGeneralpractitioner,omitempty"`
	IncludedOrganizationResourcesReferencedByGeneralpractitioner           *[]Organization                `bson:"_includedOrganizationResourcesReferencedByGeneralpractitioner,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner       *[]PractitionerRole            `bson:"_includedPractitionerRoleResourcesReferencedByGeneralpractitioner,omitempty"`
	IncludedPatientResourcesReferencedByLink                               *[]Patient                     `bson:"_includedPatientResourcesReferencedByLink,omitempty"`
	IncludedRelatedPersonResourcesReferencedByLink                         *[]RelatedPerson               `bson:"_includedRelatedPersonResourcesReferencedByLink,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization                `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingPatient                      *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingPatient,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                          *[]Account                     `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedAccountResourcesReferencingPatient                          *[]Account                     `bson:"_revIncludedAccountResourcesReferencingPatient,omitempty"`
	RevIncludedInvoiceResourcesReferencingSubject                          *[]Invoice                     `bson:"_revIncludedInvoiceResourcesReferencingSubject,omitempty"`
	RevIncludedInvoiceResourcesReferencingParticipant                      *[]Invoice                     `bson:"_revIncludedInvoiceResourcesReferencingParticipant,omitempty"`
	RevIncludedInvoiceResourcesReferencingPatient                          *[]Invoice                     `bson:"_revIncludedInvoiceResourcesReferencingPatient,omitempty"`
	RevIncludedInvoiceResourcesReferencingRecipient                        *[]Invoice                     `bson:"_revIncludedInvoiceResourcesReferencingRecipient,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition             `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject                 *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                  *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingPatient                 *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient               *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedGoalResourcesReferencingPatient                             *[]Goal                        `bson:"_revIncludedGoalResourcesReferencingPatient,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                             *[]Goal                        `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingSubject                *[]EnrollmentRequest           `bson:"_revIncludedEnrollmentRequestResourcesReferencingSubject,omitempty"`
	RevIncludedEnrollmentRequestResourcesReferencingPatient                *[]EnrollmentRequest           `bson:"_revIncludedEnrollmentRequestResourcesReferencingPatient,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                            *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingPatient                          *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingPatient,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                        *[]Consent                     `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                     `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedResearchSubjectResourcesReferencingIndividual               *[]ResearchSubject             `bson:"_revIncludedResearchSubjectResourcesReferencingIndividual,omitempty"`
	RevIncludedResearchSubjectResourcesReferencingPatient                  *[]ResearchSubject             `bson:"_revIncludedResearchSubjectResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingPatient                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingPatient,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                 *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedCoverageEligibilityResponseResourcesReferencingPatient      *[]CoverageEligibilityResponse `bson:"_revIncludedCoverageEligibilityResponseResourcesReferencingPatient,omitempty"`
	RevIncludedMeasureReportResourcesReferencingPatient                    *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingPatient,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedServiceRequestResourcesReferencingSubject                   *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPatient                   *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingPatient,omitempty"`
	RevIncludedRelatedPersonResourcesReferencingPatient                    *[]RelatedPerson               `bson:"_revIncludedRelatedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                  *[]SupplyRequest               `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSubject                    *[]SupplyRequest               `bson:"_revIncludedSupplyRequestResourcesReferencingSubject,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult          `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedBodyStructureResourcesReferencingPatient                    *[]BodyStructure               `bson:"_revIncludedBodyStructureResourcesReferencingPatient,omitempty"`
	RevIncludedPersonResourcesReferencingLink                              *[]Person                      `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedPersonResourcesReferencingPatient                           *[]Person                      `bson:"_revIncludedPersonResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingPatient                         *[]Contract                    `bson:"_revIncludedContractResourcesReferencingPatient,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                    `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingSigner                          *[]Contract                    `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject                   *[]RiskAssessment              `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPatient                   *[]RiskAssessment              `bson:"_revIncludedRiskAssessmentResourcesReferencingPatient,omitempty"`
	RevIncludedGroupResourcesReferencingMember                             *[]Group                       `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice               `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice               `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition          `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCareTeamResourcesReferencingPatient                         *[]CareTeam                    `bson:"_revIncludedCareTeamResourcesReferencingPatient,omitempty"`
	RevIncludedCareTeamResourcesReferencingSubject                         *[]CareTeam                    `bson:"_revIncludedCareTeamResourcesReferencingSubject,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                     *[]CareTeam                    `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide         `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                   *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedImagingStudyResourcesReferencingSubject                     *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingSubject,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPatient                     *[]ImagingStudy                `bson:"_revIncludedImagingStudyResourcesReferencingPatient,omitempty"`
	RevIncludedFamilyMemberHistoryResourcesReferencingPatient              *[]FamilyMemberHistory         `bson:"_revIncludedFamilyMemberHistoryResourcesReferencingPatient,omitempty"`
	RevIncludedChargeItemResourcesReferencingSubject                       *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingSubject,omitempty"`
	RevIncludedChargeItemResourcesReferencingPatient                       *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingPatient,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                       *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformeractor                *[]ChargeItem                  `bson:"_revIncludedChargeItemResourcesReferencingPerformeractor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition   `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingSubject                        *[]Encounter                   `bson:"_revIncludedEncounterResourcesReferencingSubject,omitempty"`
	RevIncludedEncounterResourcesReferencingPatient                        *[]Encounter                   `bson:"_revIncludedEncounterResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationResourcesReferencingSubject                    *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                     *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingPatient                    *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                  *[]Communication               `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition          `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                     `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                     `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedImmunizationEvaluationResourcesReferencingPatient           *[]ImmunizationEvaluation      `bson:"_revIncludedImmunizationEvaluationResourcesReferencingPatient,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingSubject               *[]DeviceUseStatement          `bson:"_revIncludedDeviceUseStatementResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingPatient               *[]DeviceUseStatement          `bson:"_revIncludedDeviceUseStatementResourcesReferencingPatient,omitempty"`
	RevIncludedRequestGroupResourcesReferencingSubject                     *[]RequestGroup                `bson:"_revIncludedRequestGroupResourcesReferencingSubject,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant                 *[]RequestGroup                `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedRequestGroupResourcesReferencingPatient                     *[]RequestGroup                `bson:"_revIncludedRequestGroupResourcesReferencingPatient,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer                  *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingSubject                    *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPatient                    *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingPatient,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest               `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader               `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingPatient       *[]ImmunizationRecommendation  `bson:"_revIncludedImmunizationRecommendationResourcesReferencingPatient,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation  `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                         *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingPatient                       *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingPatient,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                  `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                               *[]Task                        `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                           *[]Task                        `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                        `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                        `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                        `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedTaskResourcesReferencingPatient                             *[]Task                        `bson:"_revIncludedTaskResourcesReferencingPatient,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee               *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPatient             *[]ExplanationOfBenefit        `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPatient,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                         *[]Specimen                    `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedSpecimenResourcesReferencingPatient                         *[]Specimen                    `bson:"_revIncludedSpecimenResourcesReferencingPatient,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder              *[]AllergyIntolerance          `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingAsserter              *[]AllergyIntolerance          `bson:"_revIncludedAllergyIntoleranceResourcesReferencingAsserter,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingPatient               *[]AllergyIntolerance          `bson:"_revIncludedAllergyIntoleranceResourcesReferencingPatient,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                       *[]CarePlan                    `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject                         *[]CarePlan                    `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingPatient                         *[]CarePlan                    `bson:"_revIncludedCarePlanResourcesReferencingPatient,omitempty"`
	RevIncludedEpisodeOfCareResourcesReferencingPatient                    *[]EpisodeOfCare               `bson:"_revIncludedEpisodeOfCareResourcesReferencingPatient,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                      *[]Procedure                   `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject                        *[]Procedure                   `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingPatient                        *[]Procedure                   `bson:"_revIncludedProcedureResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                        `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                             *[]List                        `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingPatient                             *[]List                        `bson:"_revIncludedListResourcesReferencingPatient,omitempty"`
	RevIncludedListResourcesReferencingSource                              *[]List                        `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedImmunizationResourcesReferencingPatient                     *[]Immunization                `bson:"_revIncludedImmunizationResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester              *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingSubject                *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingPatient                *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendedperformer      *[]MedicationRequest           `bson:"_revIncludedMedicationRequestResourcesReferencingIntendedperformer,omitempty"`
	RevIncludedDeviceResourcesReferencingPatient                           *[]Device                      `bson:"_revIncludedDeviceResourcesReferencingPatient,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingPatient               *[]VisionPrescription          `bson:"_revIncludedVisionPrescriptionResourcesReferencingPatient,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                       `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                           *[]Media                       `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedMediaResourcesReferencingPatient                            *[]Media                       `bson:"_revIncludedMediaResourcesReferencingPatient,omitempty"`
	RevIncludedMolecularSequenceResourcesReferencingPatient                *[]MolecularSequence           `bson:"_revIncludedMolecularSequenceResourcesReferencingPatient,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable            `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                             *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingPatient                             *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingPatient,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                              *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse         `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPatient              *[]AppointmentResponse         `bson:"_revIncludedAppointmentResponseResourcesReferencingPatient,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder                    *[]AdverseEvent                `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject                     *[]AdverseEvent                `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingPatient                 *[]GuidanceResponse            `bson:"_revIncludedGuidanceResponseResourcesReferencingPatient,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingSubject                 *[]GuidanceResponse            `bson:"_revIncludedGuidanceResponseResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                      *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPatient                      *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingPatient,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                    *[]Observation                 `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer       *[]MedicationAdministration    `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingSubject         *[]MedicationAdministration    `bson:"_revIncludedMedicationAdministrationResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPatient         *[]MedicationAdministration    `bson:"_revIncludedMedicationAdministrationResourcesReferencingPatient,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                     `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSubject              *[]MedicationStatement         `bson:"_revIncludedMedicationStatementResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPatient              *[]MedicationStatement         `bson:"_revIncludedMedicationStatementResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource               *[]MedicationStatement         `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester           *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSubject             *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender              *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingPatient             *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingPatient,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest        `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                       `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingPatient                            *[]Basic                       `bson:"_revIncludedBasicResourcesReferencingPatient,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                             *[]Basic                       `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedClaimResponseResourcesReferencingPatient                    *[]ClaimResponse               `bson:"_revIncludedClaimResponseResourcesReferencingPatient,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer             *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingReceiver              *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingSubject               *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPatient               *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingPatient,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject                 *[]DiagnosticReport            `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingPatient                 *[]DiagnosticReport            `bson:"_revIncludedDiagnosticReportResourcesReferencingPatient,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingPatient                   *[]NutritionOrder              `bson:"_revIncludedNutritionOrderResourcesReferencingPatient,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                    `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                         *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingSource                        *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath1                  *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingPatientPath1,omitempty"`
	RevIncludedAuditEventResourcesReferencingPatientPath2                  *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingPatientPath2,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                  `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingSubject                        *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingSubject,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                       *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedConditionResourcesReferencingPatient                        *[]Condition                   `bson:"_revIncludedConditionResourcesReferencingPatient,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                       *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                     *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedCompositionResourcesReferencingPatient                      *[]Composition                 `bson:"_revIncludedCompositionResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingPatient                    *[]DetectedIssue               `bson:"_revIncludedDetectedIssueResourcesReferencingPatient,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue               `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingLink                             *[]Patient                     `bson:"_revIncludedPatientResourcesReferencingLink,omitempty"`
	RevIncludedCoverageResourcesReferencingPayor                           *[]Coverage                    `bson:"_revIncludedCoverageResourcesReferencingPayor,omitempty"`
	RevIncludedCoverageResourcesReferencingSubscriber                      *[]Coverage                    `bson:"_revIncludedCoverageResourcesReferencingSubscriber,omitempty"`
	RevIncludedCoverageResourcesReferencingBeneficiary                     *[]Coverage                    `bson:"_revIncludedCoverageResourcesReferencingBeneficiary,omitempty"`
	RevIncludedCoverageResourcesReferencingPatient                         *[]Coverage                    `bson:"_revIncludedCoverageResourcesReferencingPatient,omitempty"`
	RevIncludedCoverageResourcesReferencingPolicyholder                    *[]Coverage                    `bson:"_revIncludedCoverageResourcesReferencingPolicyholder,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor             *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingPatient            *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingPatient,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource             *[]QuestionnaireResponse       `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedCoverageEligibilityRequestResourcesReferencingPatient       *[]CoverageEligibilityRequest  `bson:"_revIncludedCoverageEligibilityRequestResourcesReferencingPatient,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                    `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedSupplyDeliveryResourcesReferencingPatient                   *[]SupplyDelivery              `bson:"_revIncludedSupplyDeliveryResourcesReferencingPatient,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSubject               *[]ClinicalImpression          `bson:"_revIncludedClinicalImpressionResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPatient               *[]ClinicalImpression          `bson:"_revIncludedClinicalImpressionResourcesReferencingPatient,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression          `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition              `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                              *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
	RevIncludedClaimResourcesReferencingPatient                            *[]Claim                       `bson:"_revIncludedClaimResourcesReferencingPatient,omitempty"`
}

func (p *PatientPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByGeneralpractitioner() (practitioners []Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByGeneralpractitioner() (organizations []Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByGeneralpractitioner() (practitionerRoles []PractitionerRole, err error) {
	if p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedPatientResourceReferencedByLink() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByLink == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByLink) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByLink))
	} else if len(*p.IncludedPatientResourcesReferencedByLink) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByLink)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByLink() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByLink == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByLink) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByLink))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByLink) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByLink)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*p.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingPatient() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAccountResourcesReferencingPatient() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingSubject() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingPatient() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingRecipient() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingPatient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGoalResourcesReferencingPatient() (goals []Goal, err error) {
	if p.RevIncludedGoalResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *p.RevIncludedGoalResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if p.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *p.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEnrollmentRequestResourcesReferencingSubject() (enrollmentRequests []EnrollmentRequest, err error) {
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded enrollmentRequests not requested")
	} else {
		enrollmentRequests = *p.RevIncludedEnrollmentRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEnrollmentRequestResourcesReferencingPatient() (enrollmentRequests []EnrollmentRequest, err error) {
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded enrollmentRequests not requested")
	} else {
		enrollmentRequests = *p.RevIncludedEnrollmentRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingPatient() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchSubjectResourcesReferencingIndividual() (researchSubjects []ResearchSubject, err error) {
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual == nil {
		err = errors.New("RevIncluded researchSubjects not requested")
	} else {
		researchSubjects = *p.RevIncludedResearchSubjectResourcesReferencingIndividual
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchSubjectResourcesReferencingPatient() (researchSubjects []ResearchSubject, err error) {
	if p.RevIncludedResearchSubjectResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded researchSubjects not requested")
	} else {
		researchSubjects = *p.RevIncludedResearchSubjectResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingPatient() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageEligibilityResponseResourcesReferencingPatient() (coverageEligibilityResponses []CoverageEligibilityResponse, err error) {
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded coverageEligibilityResponses not requested")
	} else {
		coverageEligibilityResponses = *p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingPatient() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingSubject() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPatient() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRelatedPersonResourcesReferencingPatient() (relatedPeople []RelatedPerson, err error) {
	if p.RevIncludedRelatedPersonResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded relatedPeople not requested")
	} else {
		relatedPeople = *p.RevIncludedRelatedPersonResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSubject() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if p.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *p.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBodyStructureResourcesReferencingPatient() (bodyStructures []BodyStructure, err error) {
	if p.RevIncludedBodyStructureResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded bodyStructures not requested")
	} else {
		bodyStructures = *p.RevIncludedBodyStructureResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingLink
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPersonResourcesReferencingPatient() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingPatient() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingSubject() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPatient() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingPatient() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingSubject() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingSubject() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPatient() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFamilyMemberHistoryResourcesReferencingPatient() (familyMemberHistories []FamilyMemberHistory, err error) {
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded familyMemberHistories not requested")
	} else {
		familyMemberHistories = *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingSubject() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPatient() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingSubject() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPatient() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSubject() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPatient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationEvaluationResourcesReferencingPatient() (immunizationEvaluations []ImmunizationEvaluation, err error) {
	if p.RevIncludedImmunizationEvaluationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded immunizationEvaluations not requested")
	} else {
		immunizationEvaluations = *p.RevIncludedImmunizationEvaluationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingSubject() (deviceUseStatements []DeviceUseStatement, err error) {
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *p.RevIncludedDeviceUseStatementResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingPatient() (deviceUseStatements []DeviceUseStatement, err error) {
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *p.RevIncludedDeviceUseStatementResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingSubject() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingPatient() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPatient() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingPatient() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingPatient() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedTaskResourcesReferencingPatient() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPatient() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingPatient() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingPatient() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingSubject() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPatient() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingPatient() (episodeOfCares []EpisodeOfCare, err error) {
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *p.RevIncludedEpisodeOfCareResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingSubject() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPatient() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingPatient() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPatient() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingSubject() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingPatient() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingPatient() (devices []Device, err error) {
	if p.RevIncludedDeviceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *p.RevIncludedDeviceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingPatient() (visionPrescriptions []VisionPrescription, err error) {
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *p.RevIncludedVisionPrescriptionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMediaResourcesReferencingPatient() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMolecularSequenceResourcesReferencingPatient() (molecularSequences []MolecularSequence, err error) {
	if p.RevIncludedMolecularSequenceResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded molecularSequences not requested")
	} else {
		molecularSequences = *p.RevIncludedMolecularSequenceResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingPatient() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingPatient() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingPatient() (guidanceResponses []GuidanceResponse, err error) {
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *p.RevIncludedGuidanceResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingSubject() (guidanceResponses []GuidanceResponse, err error) {
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *p.RevIncludedGuidanceResponseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPatient() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingSubject() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPatient() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSubject() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPatient() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSubject() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingPatient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingPatient() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingPatient() (claimResponses []ClaimResponse, err error) {
	if p.RevIncludedClaimResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *p.RevIncludedClaimResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingReceiver() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingReceiver
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingSubject() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPatient() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPatient() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingPatient() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingPatientPath1() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingPatientPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingPatientPath2() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingPatientPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingSubject() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedConditionResourcesReferencingPatient() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingPatient() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingPatient() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPatientResourcesReferencingLink() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingLink == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingLink
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPayor() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingPayor == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingPayor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingSubscriber() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingSubscriber == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingSubscriber
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingBeneficiary() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingBeneficiary == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingBeneficiary
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPatient() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPolicyholder() (coverages []Coverage, err error) {
	if p.RevIncludedCoverageResourcesReferencingPolicyholder == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *p.RevIncludedCoverageResourcesReferencingPolicyholder
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingPatient() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingPatient() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if p.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *p.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingPatient() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSubject() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSubject
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPatient() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (p *PatientPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPatient() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPatient
	}
	return
}

func (p *PatientPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedOrganizationResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByLink {
			rsc := (*p.IncludedPatientResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByLink != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByLink {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingPatient {
			rsc := (*p.RevIncludedAccountResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingSubject {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingPatient {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingPatient {
			rsc := (*p.RevIncludedGoalResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*p.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedConsentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingPatient {
			rsc := (*p.RevIncludedConsentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingIndividual {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchSubjectResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingPatient {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
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
	if p.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBodyStructureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBodyStructureResourcesReferencingPatient {
			rsc := (*p.RevIncludedBodyStructureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCareTeamResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingPatient {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingSubject)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPatient {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPatient)[idx]
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
	if p.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*p.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationEvaluationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationEvaluationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationEvaluationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingPatient {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedTaskResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingPatient {
			rsc := (*p.RevIncludedTaskResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSubject {
			rsc := (*p.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedListResourcesReferencingPatient {
			rsc := (*p.RevIncludedListResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingPatient {
			rsc := (*p.RevIncludedMediaResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMolecularSequenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMolecularSequenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedMolecularSequenceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingPatient {
			rsc := (*p.RevIncludedFlagResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*p.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPatient {
			rsc := (*p.RevIncludedObservationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingPatient {
			rsc := (*p.RevIncludedBasicResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*p.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingPatient {
			rsc := (*p.RevIncludedConditionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingPatient {
			rsc := (*p.RevIncludedCompositionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*p.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingBeneficiary != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingBeneficiary {
			rsc := (*p.RevIncludedCoverageResourcesReferencingBeneficiary)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *PatientPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedOrganizationResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner != nil {
		for idx := range *p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner {
			rsc := (*p.IncludedPractitionerRoleResourcesReferencedByGeneralpractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByLink != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByLink {
			rsc := (*p.IncludedPatientResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByLink != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByLink {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*p.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAccountResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingPatient {
			rsc := (*p.RevIncludedAccountResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingSubject {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingPatient {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedDocumentManifestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingPatient {
			rsc := (*p.RevIncludedGoalResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*p.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEnrollmentRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEnrollmentRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedEnrollmentRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedConsentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingPatient {
			rsc := (*p.RevIncludedConsentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedResearchSubjectResourcesReferencingIndividual != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingIndividual {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingIndividual)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchSubjectResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedResearchSubjectResourcesReferencingPatient {
			rsc := (*p.RevIncludedResearchSubjectResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
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
	if p.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRelatedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRelatedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedRelatedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBodyStructureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBodyStructureResourcesReferencingPatient {
			rsc := (*p.RevIncludedBodyStructureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPatient {
			rsc := (*p.RevIncludedPersonResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingPatient {
			rsc := (*p.RevIncludedContractResourcesReferencingPatient)[idx]
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
	if p.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPatient {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCareTeamResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingPatient {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingSubject)[idx]
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
	if p.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPatient {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient {
			rsc := (*p.RevIncludedFamilyMemberHistoryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPatient {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPatient)[idx]
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
	if p.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*p.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPatient {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationEvaluationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationEvaluationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationEvaluationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingPatient {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedProvenanceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingPatient {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedTaskResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingPatient {
			rsc := (*p.RevIncludedTaskResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPatient {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingPatient {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAllergyIntoleranceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingPatient {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPatient {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingPatient {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*p.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPatient {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSubject {
			rsc := (*p.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedListResourcesReferencingPatient {
			rsc := (*p.RevIncludedListResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPatient {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDeviceResourcesReferencingPatient {
			rsc := (*p.RevIncludedDeviceResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPatient {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMediaResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingPatient {
			rsc := (*p.RevIncludedMediaResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMolecularSequenceResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMolecularSequenceResourcesReferencingPatient {
			rsc := (*p.RevIncludedMolecularSequenceResourcesReferencingPatient)[idx]
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
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingPatient {
			rsc := (*p.RevIncludedFlagResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAppointmentResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*p.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPatient {
			rsc := (*p.RevIncludedObservationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPatient)[idx]
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
	if p.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
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
	if p.RevIncludedCommunicationRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingPatient)[idx]
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
	if p.RevIncludedBasicResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingPatient {
			rsc := (*p.RevIncludedBasicResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPatient {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPatient {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingPatient {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingPatient)[idx]
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
	if p.RevIncludedAuditEventResourcesReferencingPatientPath1 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath1 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingPatientPath2 != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingPatientPath2 {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingPatientPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*p.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingPatient {
			rsc := (*p.RevIncludedConditionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedCompositionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingPatient {
			rsc := (*p.RevIncludedCompositionResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingPatient {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingLink {
			rsc := (*p.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*p.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingBeneficiary != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingBeneficiary {
			rsc := (*p.RevIncludedCoverageResourcesReferencingBeneficiary)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *p.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*p.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
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
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPatient {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingPatient {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingPatient {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingPatient)[idx]
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
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPatient != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPatient {
			rsc := (*p.RevIncludedClaimResourcesReferencingPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
