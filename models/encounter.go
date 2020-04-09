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

type Encounter struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status          string                             `bson:"status,omitempty" json:"status,omitempty"`
	StatusHistory   []EncounterStatusHistoryComponent  `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Class           *Coding                            `bson:"class,omitempty" json:"class,omitempty"`
	ClassHistory    []EncounterClassHistoryComponent   `bson:"classHistory,omitempty" json:"classHistory,omitempty"`
	Type            []CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	ServiceType     *CodeableConcept                   `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Priority        *CodeableConcept                   `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject         *Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	EpisodeOfCare   []Reference                        `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	BasedOn         []Reference                        `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Participant     []EncounterParticipantComponent    `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment     []Reference                        `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Period          *Period                            `bson:"period,omitempty" json:"period,omitempty"`
	Length          *Quantity                          `bson:"length,omitempty" json:"length,omitempty"`
	ReasonCode      []CodeableConcept                  `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference []Reference                        `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Diagnosis       []EncounterDiagnosisComponent      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account         []Reference                        `bson:"account,omitempty" json:"account,omitempty"`
	Hospitalization *EncounterHospitalizationComponent `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Location        []EncounterLocationComponent       `bson:"location,omitempty" json:"location,omitempty"`
	ServiceProvider *Reference                         `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	PartOf          *Reference                         `bson:"partOf,omitempty" json:"partOf,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Encounter) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Encounter"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "encounter" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type encounter Encounter

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Encounter) UnmarshalJSON(data []byte) (err error) {
	x2 := encounter{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Encounter(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Encounter) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Encounter"
	} else if x.ResourceType != "Encounter" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Encounter, instead received %s", x.ResourceType))
	}
	return nil
}

type EncounterStatusHistoryComponent struct {
	BackboneElement `bson:",inline"`
	Status          string  `bson:"status,omitempty" json:"status,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterClassHistoryComponent struct {
	BackboneElement `bson:",inline"`
	Class           *Coding `bson:"class,omitempty" json:"class,omitempty"`
	Period          *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Individual      *Reference        `bson:"individual,omitempty" json:"individual,omitempty"`
}

type EncounterDiagnosisComponent struct {
	BackboneElement `bson:",inline"`
	Condition       *Reference       `bson:"condition,omitempty" json:"condition,omitempty"`
	Use             *CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
	Rank            *uint32          `bson:"rank,omitempty" json:"rank,omitempty"`
}

type EncounterHospitalizationComponent struct {
	BackboneElement        `bson:",inline"`
	PreAdmissionIdentifier *Identifier       `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference        `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept  `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept  `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	DietPreference         []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialCourtesy        []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	SpecialArrangement     []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	Destination            *Reference        `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept  `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}

type EncounterLocationComponent struct {
	BackboneElement `bson:",inline"`
	Location        *Reference       `bson:"location,omitempty" json:"location,omitempty"`
	Status          string           `bson:"status,omitempty" json:"status,omitempty"`
	PhysicalType    *CodeableConcept `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

type EncounterPlus struct {
	Encounter                     `bson:",inline"`
	EncounterPlusRelatedResources `bson:",inline"`
}

type EncounterPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPractitioner                  *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare                *[]EpisodeOfCare              `bson:"_includedEpisodeOfCareResourcesReferencedByEpisodeofcare,omitempty"`
	IncludedConditionResourcesReferencedByDiagnosis                        *[]Condition                  `bson:"_includedConditionResourcesReferencedByDiagnosis,omitempty"`
	IncludedProcedureResourcesReferencedByDiagnosis                        *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByDiagnosis,omitempty"`
	IncludedAppointmentResourcesReferencedByAppointment                    *[]Appointment                `bson:"_includedAppointmentResourcesReferencedByAppointment,omitempty"`
	IncludedEncounterResourcesReferencedByPartof                           *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByPartof,omitempty"`
	IncludedPractitionerResourcesReferencedByParticipant                   *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByParticipant               *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant                  *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedConditionResourcesReferencedByReasonreference                  *[]Condition                  `bson:"_includedConditionResourcesReferencedByReasonreference,omitempty"`
	IncludedObservationResourcesReferencedByReasonreference                *[]Observation                `bson:"_includedObservationResourcesReferencedByReasonreference,omitempty"`
	IncludedProcedureResourcesReferencedByReasonreference                  *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByReasonreference,omitempty"`
	IncludedImmunizationRecommendationResourcesReferencedByReasonreference *[]ImmunizationRecommendation `bson:"_includedImmunizationRecommendationResourcesReferencedByReasonreference,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	IncludedOrganizationResourcesReferencedByServiceprovider               *[]Organization               `bson:"_includedOrganizationResourcesReferencedByServiceprovider,omitempty"`
	IncludedAccountResourcesReferencedByAccount                            *[]Account                    `bson:"_includedAccountResourcesReferencedByAccount,omitempty"`
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
	RevIncludedDocumentReferenceResourcesReferencingEncounter              *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingEncounter,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingEncounter                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingEncounter                 *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingEncounter,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCareTeamResourcesReferencingEncounter                       *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingEncounter,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingEncounter                   *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingEncounter,omitempty"`
	RevIncludedChargeItemResourcesReferencingContext                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingContext,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingPartof                         *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingEncounter                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedRequestGroupResourcesReferencingEncounter                   *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingEncounter,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingEncounter                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingEncounter                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingEncounter,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingEncounter           *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingEncounter,omitempty"`
	RevIncludedCarePlanResourcesReferencingEncounter                       *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingEncounter,omitempty"`
	RevIncludedProcedureResourcesReferencingEncounter                      *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingEncounter,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingEncounter                           *[]List                       `bson:"_revIncludedListResourcesReferencingEncounter,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingEncounter              *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedVisionPrescriptionResourcesReferencingEncounter             *[]VisionPrescription         `bson:"_revIncludedVisionPrescriptionResourcesReferencingEncounter,omitempty"`
	RevIncludedMediaResourcesReferencingEncounter                          *[]Media                      `bson:"_revIncludedMediaResourcesReferencingEncounter,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingEncounter                           *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingEncounter,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingEncounter                    *[]Observation                `bson:"_revIncludedObservationResourcesReferencingEncounter,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingContext         *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingContext,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingContext              *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingContext,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingEncounter           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingEncounter,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingContext               *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingContext,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingEncounter               *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingEncounter,omitempty"`
	RevIncludedNutritionOrderResourcesReferencingEncounter                 *[]NutritionOrder             `bson:"_revIncludedNutritionOrderResourcesReferencingEncounter,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingEncounter                      *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEncounter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEncounter                    *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEncounter,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingEncounter          *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingEncounter,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingEncounter             *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingEncounter,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingEncounter                          *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingEncounter,omitempty"`
}

func (e *EncounterPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPractitioner() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByPractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByPractitioner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByPractitioner))
	} else if len(*e.IncludedPractitionerResourcesReferencedByPractitioner) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByPractitioner)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if e.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*e.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*e.IncludedGroupResourcesReferencedBySubject))
	} else if len(*e.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*e.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedBySubject))
	} else if len(*e.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedEpisodeOfCareResourcesReferencedByEpisodeofcare() (episodeOfCares []EpisodeOfCare, err error) {
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare == nil {
		err = errors.New("Included episodeOfCares not requested")
	} else {
		episodeOfCares = *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedConditionResourceReferencedByDiagnosis() (condition *Condition, err error) {
	if e.IncludedConditionResourcesReferencedByDiagnosis == nil {
		err = errors.New("Included conditions not requested")
	} else if len(*e.IncludedConditionResourcesReferencedByDiagnosis) > 1 {
		err = fmt.Errorf("Expected 0 or 1 condition, but found %d", len(*e.IncludedConditionResourcesReferencedByDiagnosis))
	} else if len(*e.IncludedConditionResourcesReferencedByDiagnosis) == 1 {
		condition = &(*e.IncludedConditionResourcesReferencedByDiagnosis)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedProcedureResourceReferencedByDiagnosis() (procedure *Procedure, err error) {
	if e.IncludedProcedureResourcesReferencedByDiagnosis == nil {
		err = errors.New("Included procedures not requested")
	} else if len(*e.IncludedProcedureResourcesReferencedByDiagnosis) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedure, but found %d", len(*e.IncludedProcedureResourcesReferencedByDiagnosis))
	} else if len(*e.IncludedProcedureResourcesReferencedByDiagnosis) == 1 {
		procedure = &(*e.IncludedProcedureResourcesReferencedByDiagnosis)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedAppointmentResourcesReferencedByAppointment() (appointments []Appointment, err error) {
	if e.IncludedAppointmentResourcesReferencedByAppointment == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *e.IncludedAppointmentResourcesReferencedByAppointment
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedEncounterResourceReferencedByPartof() (encounter *Encounter, err error) {
	if e.IncludedEncounterResourcesReferencedByPartof == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*e.IncludedEncounterResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*e.IncludedEncounterResourcesReferencedByPartof))
	} else if len(*e.IncludedEncounterResourcesReferencedByPartof) == 1 {
		encounter = &(*e.IncludedEncounterResourcesReferencedByPartof)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParticipant() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByParticipant))
	} else if len(*e.IncludedPractitionerResourcesReferencedByParticipant) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByParticipant)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByParticipant() (practitionerRole *PractitionerRole, err error) {
	if e.IncludedPractitionerRoleResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*e.IncludedPractitionerRoleResourcesReferencedByParticipant))
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByParticipant) == 1 {
		practitionerRole = &(*e.IncludedPractitionerRoleResourcesReferencedByParticipant)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParticipant() (relatedPerson *RelatedPerson, err error) {
	if e.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*e.IncludedRelatedPersonResourcesReferencedByParticipant))
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByParticipant) == 1 {
		relatedPerson = &(*e.IncludedRelatedPersonResourcesReferencedByParticipant)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if e.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *e.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedConditionResourcesReferencedByReasonreference() (conditions []Condition, err error) {
	if e.IncludedConditionResourcesReferencedByReasonreference == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *e.IncludedConditionResourcesReferencedByReasonreference
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedObservationResourcesReferencedByReasonreference() (observations []Observation, err error) {
	if e.IncludedObservationResourcesReferencedByReasonreference == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *e.IncludedObservationResourcesReferencedByReasonreference
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedProcedureResourcesReferencedByReasonreference() (procedures []Procedure, err error) {
	if e.IncludedProcedureResourcesReferencedByReasonreference == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *e.IncludedProcedureResourcesReferencedByReasonreference
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedImmunizationRecommendationResourcesReferencedByReasonreference() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference == nil {
		err = errors.New("Included immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByLocation))
	} else if len(*e.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedOrganizationResourceReferencedByServiceprovider() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByServiceprovider == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByServiceprovider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByServiceprovider))
	} else if len(*e.IncludedOrganizationResourcesReferencedByServiceprovider) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByServiceprovider)[0]
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedAccountResourcesReferencedByAccount() (accounts []Account, err error) {
	if e.IncludedAccountResourcesReferencedByAccount == nil {
		err = errors.New("Included accounts not requested")
	} else {
		accounts = *e.IncludedAccountResourcesReferencedByAccount
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *e.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingEncounter() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingEncounter() (serviceRequests []ServiceRequest, err error) {
	if e.RevIncludedServiceRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *e.RevIncludedServiceRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if e.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *e.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingEncounter() (riskAssessments []RiskAssessment, err error) {
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *e.RevIncludedRiskAssessmentResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingEncounter() (careTeams []CareTeam, err error) {
	if e.RevIncludedCareTeamResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *e.RevIncludedCareTeamResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingEncounter() (imagingStudies []ImagingStudy, err error) {
	if e.RevIncludedImagingStudyResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *e.RevIncludedImagingStudyResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingContext() (chargeItems []ChargeItem, err error) {
	if e.RevIncludedChargeItemResourcesReferencingContext == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *e.RevIncludedChargeItemResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPartof() (encounters []Encounter, err error) {
	if e.RevIncludedEncounterResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *e.RevIncludedEncounterResourcesReferencingPartof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingEncounter() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingEncounter() (requestGroups []RequestGroup, err error) {
	if e.RevIncludedRequestGroupResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *e.RevIncludedRequestGroupResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingEncounter() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingEncounter() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingEncounter() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingEncounter() (carePlans []CarePlan, err error) {
	if e.RevIncludedCarePlanResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *e.RevIncludedCarePlanResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingEncounter() (procedures []Procedure, err error) {
	if e.RevIncludedProcedureResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *e.RevIncludedProcedureResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedListResourcesReferencingEncounter() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingEncounter() (medicationRequests []MedicationRequest, err error) {
	if e.RevIncludedMedicationRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *e.RevIncludedMedicationRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingEncounter() (visionPrescriptions []VisionPrescription, err error) {
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMediaResourcesReferencingEncounter() (media []Media, err error) {
	if e.RevIncludedMediaResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *e.RevIncludedMediaResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedFlagResourcesReferencingEncounter() (flags []Flag, err error) {
	if e.RevIncludedFlagResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *e.RevIncludedFlagResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedObservationResourcesReferencingEncounter() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingContext() (medicationAdministrations []MedicationAdministration, err error) {
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *e.RevIncludedMedicationAdministrationResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingContext() (medicationStatements []MedicationStatement, err error) {
	if e.RevIncludedMedicationStatementResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *e.RevIncludedMedicationStatementResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingEncounter() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingContext() (medicationDispenses []MedicationDispense, err error) {
	if e.RevIncludedMedicationDispenseResourcesReferencingContext == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *e.RevIncludedMedicationDispenseResourcesReferencingContext
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingEncounter() (diagnosticReports []DiagnosticReport, err error) {
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *e.RevIncludedDiagnosticReportResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingEncounter() (nutritionOrders []NutritionOrder, err error) {
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *e.RevIncludedNutritionOrderResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEncounter() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEncounter() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingEncounter() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingEncounter() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EncounterPlusRelatedResources) GetRevIncludedClaimResourcesReferencingEncounter() (claims []Claim, err error) {
	if e.RevIncludedClaimResourcesReferencingEncounter == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *e.RevIncludedClaimResourcesReferencingEncounter
	}
	return
}

func (e *EncounterPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedGroupResourcesReferencedBySubject {
			rsc := (*e.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedPatientResourcesReferencedBySubject {
			rsc := (*e.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for idx := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			rsc := (*e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedConditionResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedProcedureResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for idx := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			rsc := (*e.IncludedAppointmentResourcesReferencedByAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByPartof {
			rsc := (*e.IncludedEncounterResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *e.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*e.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*e.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*e.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*e.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference {
			rsc := (*e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByLocation {
			rsc := (*e.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByServiceprovider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByServiceprovider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAccountResourcesReferencedByAccount != nil {
		for idx := range *e.IncludedAccountResourcesReferencedByAccount {
			rsc := (*e.IncludedAccountResourcesReferencedByAccount)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedServiceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedServiceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*e.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRiskAssessmentResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCareTeamResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEncounter {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedChargeItemResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedChargeItemResourcesReferencingContext {
			rsc := (*e.RevIncludedChargeItemResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingPartof {
			rsc := (*e.RevIncludedEncounterResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*e.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*e.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*e.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingEncounter {
			rsc := (*e.RevIncludedTaskResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter {
			rsc := (*e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedListResourcesReferencingEncounter {
			rsc := (*e.RevIncludedListResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedVisionPrescriptionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMediaResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedMediaResourcesReferencingEncounter {
			rsc := (*e.RevIncludedMediaResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedFlagResourcesReferencingEncounter {
			rsc := (*e.RevIncludedFlagResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*e.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedObservationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationAdministrationResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationAdministrationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationStatementResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationStatementResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationStatementResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationDispenseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationDispenseResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationDispenseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			rsc := (*e.RevIncludedNutritionOrderResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedConditionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClaimResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClaimResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClaimResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EncounterPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedGroupResourcesReferencedBySubject {
			rsc := (*e.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *e.IncludedPatientResourcesReferencedBySubject {
			rsc := (*e.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare != nil {
		for idx := range *e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare {
			rsc := (*e.IncludedEpisodeOfCareResourcesReferencedByEpisodeofcare)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedConditionResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByDiagnosis != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByDiagnosis {
			rsc := (*e.IncludedProcedureResourcesReferencedByDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAppointmentResourcesReferencedByAppointment != nil {
		for idx := range *e.IncludedAppointmentResourcesReferencedByAppointment {
			rsc := (*e.IncludedAppointmentResourcesReferencedByAppointment)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByPartof != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByPartof {
			rsc := (*e.IncludedEncounterResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByParticipant {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByParticipant {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *e.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*e.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*e.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*e.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*e.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference != nil {
		for idx := range *e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference {
			rsc := (*e.IncludedImmunizationRecommendationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByLocation {
			rsc := (*e.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByServiceprovider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByServiceprovider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByServiceprovider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedAccountResourcesReferencedByAccount != nil {
		for idx := range *e.IncludedAccountResourcesReferencedByAccount {
			rsc := (*e.IncludedAccountResourcesReferencedByAccount)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedServiceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedServiceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedServiceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*e.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRiskAssessmentResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRiskAssessmentResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRiskAssessmentResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCareTeamResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCareTeamResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCareTeamResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEncounter {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedChargeItemResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedChargeItemResourcesReferencingContext {
			rsc := (*e.RevIncludedChargeItemResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEncounterResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedEncounterResourcesReferencingPartof {
			rsc := (*e.RevIncludedEncounterResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*e.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*e.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedRequestGroupResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedRequestGroupResourcesReferencingEncounter {
			rsc := (*e.RevIncludedRequestGroupResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*e.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingEncounter {
			rsc := (*e.RevIncludedTaskResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter {
			rsc := (*e.RevIncludedExplanationOfBenefitResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCarePlanResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCarePlanResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCarePlanResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProcedureResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedProcedureResourcesReferencingEncounter {
			rsc := (*e.RevIncludedProcedureResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedListResourcesReferencingEncounter {
			rsc := (*e.RevIncludedListResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedMedicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedMedicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVisionPrescriptionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedVisionPrescriptionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedVisionPrescriptionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMediaResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedMediaResourcesReferencingEncounter {
			rsc := (*e.RevIncludedMediaResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedFlagResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedFlagResourcesReferencingEncounter {
			rsc := (*e.RevIncludedFlagResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*e.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingEncounter {
			rsc := (*e.RevIncludedObservationResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationAdministrationResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationAdministrationResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationAdministrationResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationStatementResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationStatementResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationStatementResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMedicationDispenseResourcesReferencingContext != nil {
		for idx := range *e.RevIncludedMedicationDispenseResourcesReferencingContext {
			rsc := (*e.RevIncludedMedicationDispenseResourcesReferencingContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDiagnosticReportResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedDiagnosticReportResourcesReferencingEncounter {
			rsc := (*e.RevIncludedDiagnosticReportResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedNutritionOrderResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedNutritionOrderResourcesReferencingEncounter {
			rsc := (*e.RevIncludedNutritionOrderResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedConditionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClaimResourcesReferencingEncounter != nil {
		for idx := range *e.RevIncludedClaimResourcesReferencingEncounter {
			rsc := (*e.RevIncludedClaimResourcesReferencingEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
