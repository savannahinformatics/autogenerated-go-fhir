package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Practitioner ...
type Practitioner struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                                `bson:"active,omitempty" json:"active,omitempty"`
	Name           []HumanName                          `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint                       `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address        []Address                            `bson:"address,omitempty" json:"address,omitempty"`
	Gender         string                               `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime                        `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo          []Attachment                         `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification  []PractitionerQualificationComponent `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication  []CodeableConcept                    `bson:"communication,omitempty" json:"communication,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Practitioner) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Practitioner"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "practitioner" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type practitioner Practitioner

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Practitioner) UnmarshalJSON(data []byte) (err error) {
	x2 := practitioner{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Practitioner(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Practitioner) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Practitioner"
	} else if x.ResourceType != "Practitioner" {
		return fmt.Errorf("Expected resourceType to be Practitioner, instead received %s", x.ResourceType)
	}
	return nil
}

// PractitionerQualificationComponent ...
type PractitionerQualificationComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period          *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Issuer          *Reference       `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

// PractitionerPlus ...
type PractitionerPlus struct {
	Practitioner                     `bson:",inline"`
	PractitionerPlusRelatedResources `bson:",inline"`
}

// PractitionerPlusRelatedResources ...
type PractitionerPlusRelatedResources struct {
	RevIncludedAppointmentResourcesReferencingPractitioner                 *[]Appointment                 `bson:"_revIncludedAppointmentResourcesReferencingPractitioner,omitempty"`
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
	RevIncludedDocumentManifestResourcesReferencingSubject                 *[]DocumentManifest            `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
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
	RevIncludedDocumentReferenceResourcesReferencingSubject                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthenticator          *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthenticator,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                 *[]DocumentReference           `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor    *[]CoverageEligibilityResponse `bson:"_revIncludedCoverageEligibilityResponseResourcesReferencingRequestor,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingReporter                   *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingReporter,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport               `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingPractitioner            *[]PractitionerRole            `bson:"_revIncludedPractitionerRoleResourcesReferencingPractitioner,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest              `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                  *[]SupplyRequest               `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult          `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedPersonResourcesReferencingPractitioner                      *[]Person                      `bson:"_revIncludedPersonResourcesReferencingPractitioner,omitempty"`
	RevIncludedPersonResourcesReferencingLink                              *[]Person                      `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
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
	RevIncludedEncounterResourcesReferencingPractitioner                   *[]Encounter                   `bson:"_revIncludedEncounterResourcesReferencingPractitioner,omitempty"`
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
	RevIncludedEpisodeOfCareResourcesReferencingCaremanager                *[]EpisodeOfCare               `bson:"_revIncludedEpisodeOfCareResourcesReferencingCaremanager,omitempty"`
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
	RevIncludedFlagResourcesReferencingSubject                             *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                              *[]Flag                        `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse         `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingPractitioner         *[]AppointmentResponse         `bson:"_revIncludedAppointmentResponseResourcesReferencingPractitioner,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder                    *[]AdverseEvent                `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject                     *[]AdverseEvent                `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
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
	RevIncludedMedicationDispenseResourcesReferencingReceiver              *[]MedicationDispense          `bson:"_revIncludedMedicationDispenseResourcesReferencingReceiver,omitempty"`
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

// GetRevIncludedAppointmentResourcesReferencingPractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingPractitioner() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingPractitioner
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingActor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedAccountResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if p.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *p.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

// GetRevIncludedInvoiceResourcesReferencingParticipant ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if p.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *p.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRecipient ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedConsentResourcesReferencingActor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingActor
	}
	return
}

// GetRevIncludedConsentResourcesReferencingConsentor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthenticator() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedCoverageEligibilityResponseResourcesReferencingRequestor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCoverageEligibilityResponseResourcesReferencingRequestor() (coverageEligibilityResponses []CoverageEligibilityResponse, err error) {
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded coverageEligibilityResponses not requested")
	} else {
		coverageEligibilityResponses = *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingReporter ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingReporter() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingReporter
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedPractitionerRoleResourcesReferencingPractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingPractitioner() (practitionerRoles []PractitionerRole, err error) {
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if p.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *p.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedSupplyRequestResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if p.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *p.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if p.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *p.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedPersonResourcesReferencingPractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPersonResourcesReferencingPractitioner() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingPractitioner
	}
	return
}

// GetRevIncludedPersonResourcesReferencingLink ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if p.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *p.RevIncludedPersonResourcesReferencingLink
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedContractResourcesReferencingSigner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSigner
	}
	return
}

// GetRevIncludedRiskAssessmentResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *p.RevIncludedRiskAssessmentResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedGroupResourcesReferencingManagingentity ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedGroupResourcesReferencingManagingentity() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingManagingentity == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingManagingentity
	}
	return
}

// GetRevIncludedGroupResourcesReferencingMember ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if p.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *p.RevIncludedGroupResourcesReferencingMember
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingProvider ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingProvider() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingProvider
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCareTeamResourcesReferencingParticipant ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if p.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *p.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingInterpreter ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingInterpreter() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingInterpreter
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingReferrer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingReferrer() (imagingStudies []ImagingStudy, err error) {
	if p.RevIncludedImagingStudyResourcesReferencingReferrer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *p.RevIncludedImagingStudyResourcesReferencingReferrer
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingEnterer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingPerformeractor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingPractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingPractitioner() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingPractitioner
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingParticipant ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingParticipant() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingSender ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingRecipient ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingAuthor() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedRequestGroupResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingAuthor() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedRequestGroupResourcesReferencingParticipant ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if p.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *p.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingReceiver ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingReceiver() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingReceiver
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingAuthor() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingSender ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingSender() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingSender == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingSender
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingResponsible ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingResponsible() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingResponsible
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingEnterer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingEnterer() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingEnterer
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingAgent ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingOwner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

// GetRevIncludedTaskResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingCareteam ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingCareteam() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingPayee ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingProvider ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingProvider() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingEnterer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingEnterer() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer
	}
	return
}

// GetRevIncludedResearchStudyResourcesReferencingPrincipalinvestigator ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingPrincipalinvestigator() (researchStudies []ResearchStudy, err error) {
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator
	}
	return
}

// GetRevIncludedSpecimenResourcesReferencingCollector ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingCollector() (specimen []Specimen, err error) {
	if p.RevIncludedSpecimenResourcesReferencingCollector == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *p.RevIncludedSpecimenResourcesReferencingCollector
	}
	return
}

// GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

// GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if p.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *p.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedEpisodeOfCareResourcesReferencingCaremanager ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEpisodeOfCareResourcesReferencingCaremanager() (episodeOfCares []EpisodeOfCare, err error) {
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager == nil {
		err = errors.New("RevIncluded episodeOfCares not requested")
	} else {
		episodeOfCares = *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedListResourcesReferencingSource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingSource
	}
	return
}

// GetRevIncludedImmunizationResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingPerformer() (immunizations []Immunization, err error) {
	if p.RevIncludedImmunizationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *p.RevIncludedImmunizationResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedMedicationRequestResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

// GetRevIncludedVisionPrescriptionResourcesReferencingPrescriber ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedVisionPrescriptionResourcesReferencingPrescriber() (visionPrescriptions []VisionPrescription, err error) {
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber == nil {
		err = errors.New("RevIncluded visionPrescriptions not requested")
	} else {
		visionPrescriptions = *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber
	}
	return
}

// GetRevIncludedMediaResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMediaResourcesReferencingOperator ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if p.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *p.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedFlagResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

// GetRevIncludedFlagResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingActor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingPractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingPractitioner() (appointmentResponses []AppointmentResponse, err error) {
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingRecorder ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedMedicationAdministrationResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedMedicationStatementResourcesReferencingSource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if p.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *p.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingRequester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingSender ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingRecipient ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedBasicResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedClaimResponseResourcesReferencingRequestor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequestor() (claimResponses []ClaimResponse, err error) {
	if p.RevIncludedClaimResponseResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *p.RevIncludedClaimResponseResourcesReferencingRequestor
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingReceiver ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingReceiver() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingReceiver
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingResponsibleparty ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingResponsibleparty() (medicationDispenses []MedicationDispense, err error) {
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingPerformer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingPerformer() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingResultsinterpreter ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResultsinterpreter() (diagnosticReports []DiagnosticReport, err error) {
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter
	}
	return
}

// GetRevIncludedNutritionOrderResourcesReferencingProvider ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedNutritionOrderResourcesReferencingProvider() (nutritionOrders []NutritionOrder, err error) {
	if p.RevIncludedNutritionOrderResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded nutritionOrders not requested")
	} else {
		nutritionOrders = *p.RevIncludedNutritionOrderResourcesReferencingProvider
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingAgent ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingSource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedPaymentReconciliationResourcesReferencingRequestor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPaymentReconciliationResourcesReferencingRequestor() (paymentReconciliations []PaymentReconciliation, err error) {
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor == nil {
		err = errors.New("RevIncluded paymentReconciliations not requested")
	} else {
		paymentReconciliations = *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedConditionResourcesReferencingAsserter ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingAttester ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingAuthor() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedPatientResourcesReferencingGeneralpractitioner ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPatientResourcesReferencingGeneralpractitioner() (patients []Patient, err error) {
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *p.RevIncludedPatientResourcesReferencingGeneralpractitioner
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSource ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

// GetRevIncludedCoverageEligibilityRequestResourcesReferencingProvider ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingProvider() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider
	}
	return
}

// GetRevIncludedCoverageEligibilityRequestResourcesReferencingEnterer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingEnterer() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer
	}
	return
}

// GetRevIncludedScheduleResourcesReferencingActor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if p.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *p.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

// GetRevIncludedSupplyDeliveryResourcesReferencingReceiver ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingReceiver() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver
	}
	return
}

// GetRevIncludedSupplyDeliveryResourcesReferencingSupplier ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedSupplyDeliveryResourcesReferencingSupplier() (supplyDeliveries []SupplyDelivery, err error) {
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyDeliveries not requested")
	} else {
		supplyDeliveries = *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingAssessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAssessor() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingAssessor
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedClaimResourcesReferencingCareteam ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingCareteam() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingCareteam == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingCareteam
	}
	return
}

// GetRevIncludedClaimResourcesReferencingPayee ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

// GetRevIncludedClaimResourcesReferencingProvider ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingProvider() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingProvider == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingProvider
	}
	return
}

// GetRevIncludedClaimResourcesReferencingEnterer ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedClaimResourcesReferencingEnterer() (claims []Claim, err error) {
	if p.RevIncludedClaimResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *p.RevIncludedClaimResourcesReferencingEnterer
	}
	return
}

// GetIncludedResources ...
func (p *PractitionerPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ...
func (p *PractitionerPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActor {
			rsc := (*p.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPractitionerRoleResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPersonResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*p.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingInterpreter {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingReferrer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingReferrer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (p *PractitionerPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*p.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*p.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingActor {
			rsc := (*p.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*p.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthenticator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedCoverageEligibilityResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPractitionerRoleResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPractitionerRoleResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPractitionerRoleResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedPersonResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *p.RevIncludedPersonResourcesReferencingLink {
			rsc := (*p.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSigner {
			rsc := (*p.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*p.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*p.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *p.RevIncludedGroupResourcesReferencingMember {
			rsc := (*p.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingProvider {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*p.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingInterpreter != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingInterpreter {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImagingStudyResourcesReferencingReferrer != nil {
		for idx := range *p.RevIncludedImagingStudyResourcesReferencingReferrer {
			rsc := (*p.RevIncludedImagingStudyResourcesReferencingReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedEncounterResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*p.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingAuthor {
			rsc := (*p.RevIncludedLinkageResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *p.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*p.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingAuthor {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingSender {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingResponsible != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingResponsible {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingResponsible)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingEnterer {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*p.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*p.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingCareteam)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingProvider {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer {
			rsc := (*p.RevIncludedExplanationOfBenefitResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator != nil {
		for idx := range *p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator {
			rsc := (*p.RevIncludedResearchStudyResourcesReferencingPrincipalinvestigator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSpecimenResourcesReferencingCollector != nil {
		for idx := range *p.RevIncludedSpecimenResourcesReferencingCollector {
			rsc := (*p.RevIncludedSpecimenResourcesReferencingCollector)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*p.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*p.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager != nil {
		for idx := range *p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager {
			rsc := (*p.RevIncludedEpisodeOfCareResourcesReferencingCaremanager)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedListResourcesReferencingSource {
			rsc := (*p.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedImmunizationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedImmunizationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedImmunizationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*p.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber != nil {
		for idx := range *p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber {
			rsc := (*p.RevIncludedVisionPrescriptionResourcesReferencingPrescriber)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*p.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *p.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*p.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*p.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAppointmentResponseResourcesReferencingPractitioner != nil {
		for idx := range *p.RevIncludedAppointmentResponseResourcesReferencingPractitioner {
			rsc := (*p.RevIncludedAppointmentResponseResourcesReferencingPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*p.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*p.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResponseResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedClaimResponseResourcesReferencingRequestor {
			rsc := (*p.RevIncludedClaimResponseResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingReceiver {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty != nil {
		for idx := range *p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty {
			rsc := (*p.RevIncludedMedicationDispenseResourcesReferencingResponsibleparty)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingPerformer != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingPerformer {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter != nil {
		for idx := range *p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter {
			rsc := (*p.RevIncludedDiagnosticReportResourcesReferencingResultsinterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedNutritionOrderResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedNutritionOrderResourcesReferencingProvider {
			rsc := (*p.RevIncludedNutritionOrderResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPaymentReconciliationResourcesReferencingRequestor != nil {
		for idx := range *p.RevIncludedPaymentReconciliationResourcesReferencingRequestor {
			rsc := (*p.RevIncludedPaymentReconciliationResourcesReferencingRequestor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*p.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*p.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPatientResourcesReferencingGeneralpractitioner != nil {
		for idx := range *p.RevIncludedPatientResourcesReferencingGeneralpractitioner {
			rsc := (*p.RevIncludedPatientResourcesReferencingGeneralpractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer {
			rsc := (*p.RevIncludedCoverageEligibilityRequestResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *p.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*p.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingReceiver != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingReceiver {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingReceiver)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedSupplyDeliveryResourcesReferencingSupplier != nil {
		for idx := range *p.RevIncludedSupplyDeliveryResourcesReferencingSupplier {
			rsc := (*p.RevIncludedSupplyDeliveryResourcesReferencingSupplier)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingAssessor != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingAssessor {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingAssessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingCareteam != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingCareteam {
			rsc := (*p.RevIncludedClaimResourcesReferencingCareteam)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*p.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingProvider != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingProvider {
			rsc := (*p.RevIncludedClaimResourcesReferencingProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if p.RevIncludedClaimResourcesReferencingEnterer != nil {
		for idx := range *p.RevIncludedClaimResourcesReferencingEnterer {
			rsc := (*p.RevIncludedClaimResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
