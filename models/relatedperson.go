package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RelatedPerson struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                                 `bson:"active,omitempty" json:"active,omitempty"`
	Patient        *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Relationship   []CodeableConcept                     `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name           []HumanName                           `bson:"name,omitempty" json:"name,omitempty"`
	Telecom        []ContactPoint                        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender         string                                `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate      *FHIRDateTime                         `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address        []Address                             `bson:"address,omitempty" json:"address,omitempty"`
	Photo          []Attachment                          `bson:"photo,omitempty" json:"photo,omitempty"`
	Period         *Period                               `bson:"period,omitempty" json:"period,omitempty"`
	Communication  []RelatedPersonCommunicationComponent `bson:"communication,omitempty" json:"communication,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *RelatedPerson) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RelatedPerson"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "relatedPerson" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type relatedPerson RelatedPerson

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *RelatedPerson) UnmarshalJSON(data []byte) (err error) {
	x2 := relatedPerson{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RelatedPerson(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RelatedPerson) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RelatedPerson"
	} else if x.ResourceType != "RelatedPerson" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RelatedPerson, instead received %s", x.ResourceType))
	}
	return nil
}

type RelatedPersonCommunicationComponent struct {
	BackboneElement `bson:",inline"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred       *bool            `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type RelatedPersonPlus struct {
	RelatedPerson                     `bson:",inline"`
	RelatedPersonPlusRelatedResources `bson:",inline"`
}

type RelatedPersonPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedInvoiceResourcesReferencingParticipant                      *[]Invoice                    `bson:"_revIncludedInvoiceResourcesReferencingParticipant,omitempty"`
	RevIncludedInvoiceResourcesReferencingRecipient                        *[]Invoice                    `bson:"_revIncludedInvoiceResourcesReferencingRecipient,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                  *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRecipient               *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRecipient,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                            *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedConsentResourcesReferencingConsentor                        *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingConsentor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingAuthor                 *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                  *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedPersonResourcesReferencingLink                              *[]Person                     `bson:"_revIncludedPersonResourcesReferencingLink,omitempty"`
	RevIncludedPersonResourcesReferencingRelatedperson                     *[]Person                     `bson:"_revIncludedPersonResourcesReferencingRelatedperson,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingSigner                          *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSigner,omitempty"`
	RevIncludedGroupResourcesReferencingManagingentity                     *[]Group                      `bson:"_revIncludedGroupResourcesReferencingManagingentity,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCareTeamResourcesReferencingParticipant                     *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingParticipant,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                   *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformeractor                *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingPerformeractor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingParticipant                    *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingParticipant,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingSender                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant                 *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                         *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingPayee               *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingPayee,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingRecorder              *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingRecorder,omitempty"`
	RevIncludedAllergyIntoleranceResourcesReferencingAsserter              *[]AllergyIntolerance         `bson:"_revIncludedAllergyIntoleranceResourcesReferencingAsserter,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                       *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                      *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester              *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendedperformer      *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingIntendedperformer,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                           *[]Media                      `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAdverseEventResourcesReferencingRecorder                    *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingRecorder,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject                     *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPerformer                    *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer       *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSource               *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingSource,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender              *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedBasicResourcesReferencingAuthor                             *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingAuthor,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer             *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                         *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingSource                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingAsserter                       *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingAsserter,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                       *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingAttester                     *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAttester,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedPatientResourcesReferencingLink                             *[]Patient                    `bson:"_revIncludedPatientResourcesReferencingLink,omitempty"`
	RevIncludedCoverageResourcesReferencingPayor                           *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingPayor,omitempty"`
	RevIncludedCoverageResourcesReferencingSubscriber                      *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingSubscriber,omitempty"`
	RevIncludedCoverageResourcesReferencingPolicyholder                    *[]Coverage                   `bson:"_revIncludedCoverageResourcesReferencingPolicyholder,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSource             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSource,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingPayee                              *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingPayee,omitempty"`
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if r.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*r.IncludedPatientResourcesReferencedByPatient))
	} else if len(*r.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*r.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if r.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *r.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingRecipient() (invoices []Invoice, err error) {
	if r.RevIncludedInvoiceResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *r.RevIncludedInvoiceResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRecipient() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedConsentResourcesReferencingConsentor() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingConsentor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingConsentor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if r.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *r.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if r.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *r.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if r.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *r.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if r.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *r.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if r.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *r.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPersonResourcesReferencingLink() (people []Person, err error) {
	if r.RevIncludedPersonResourcesReferencingLink == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *r.RevIncludedPersonResourcesReferencingLink
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPersonResourcesReferencingRelatedperson() (people []Person, err error) {
	if r.RevIncludedPersonResourcesReferencingRelatedperson == nil {
		err = errors.New("RevIncluded people not requested")
	} else {
		people = *r.RevIncludedPersonResourcesReferencingRelatedperson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedContractResourcesReferencingSigner() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSigner == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSigner
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedGroupResourcesReferencingManagingentity() (groups []Group, err error) {
	if r.RevIncludedGroupResourcesReferencingManagingentity == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *r.RevIncludedGroupResourcesReferencingManagingentity
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingParticipant() (careTeams []CareTeam, err error) {
	if r.RevIncludedCareTeamResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *r.RevIncludedCareTeamResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if r.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *r.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if r.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *r.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if r.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *r.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingParticipant() (encounters []Encounter, err error) {
	if r.RevIncludedEncounterResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *r.RevIncludedEncounterResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if r.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *r.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingPayee() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if r.RevIncludedExplanationOfBenefitResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *r.RevIncludedExplanationOfBenefitResourcesReferencingPayee
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingRecorder() (allergyIntolerances []AllergyIntolerance, err error) {
	if r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAllergyIntoleranceResourcesReferencingAsserter() (allergyIntolerances []AllergyIntolerance, err error) {
	if r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded allergyIntolerances not requested")
	} else {
		allergyIntolerances = *r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if r.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *r.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if r.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *r.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if r.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *r.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if r.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *r.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if r.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *r.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingRecorder() (adverseEvents []AdverseEvent, err error) {
	if r.RevIncludedAdverseEventResourcesReferencingRecorder == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *r.RevIncludedAdverseEventResourcesReferencingRecorder
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if r.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *r.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPerformer() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if r.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *r.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSource() (medicationStatements []MedicationStatement, err error) {
	if r.RevIncludedMedicationStatementResourcesReferencingSource == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *r.RevIncludedMedicationStatementResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedBasicResourcesReferencingAuthor() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if r.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *r.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if r.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *r.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedConditionResourcesReferencingAsserter() (conditions []Condition, err error) {
	if r.RevIncludedConditionResourcesReferencingAsserter == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *r.RevIncludedConditionResourcesReferencingAsserter
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAttester() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingAttester == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingAttester
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPatientResourcesReferencingLink() (patients []Patient, err error) {
	if r.RevIncludedPatientResourcesReferencingLink == nil {
		err = errors.New("RevIncluded patients not requested")
	} else {
		patients = *r.RevIncludedPatientResourcesReferencingLink
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPayor() (coverages []Coverage, err error) {
	if r.RevIncludedCoverageResourcesReferencingPayor == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *r.RevIncludedCoverageResourcesReferencingPayor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingSubscriber() (coverages []Coverage, err error) {
	if r.RevIncludedCoverageResourcesReferencingSubscriber == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *r.RevIncludedCoverageResourcesReferencingSubscriber
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedCoverageResourcesReferencingPolicyholder() (coverages []Coverage, err error) {
	if r.RevIncludedCoverageResourcesReferencingPolicyholder == nil {
		err = errors.New("RevIncluded coverages not requested")
	} else {
		coverages = *r.RevIncludedCoverageResourcesReferencingPolicyholder
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSource() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSource
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if r.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *r.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedClaimResourcesReferencingPayee() (claims []Claim, err error) {
	if r.RevIncludedClaimResourcesReferencingPayee == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *r.RevIncludedClaimResourcesReferencingPayee
	}
	return
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RelatedPersonPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*r.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*r.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingActor {
			rsc := (*r.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*r.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingLink {
			rsc := (*r.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingRelatedperson != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingRelatedperson {
			rsc := (*r.RevIncludedPersonResourcesReferencingRelatedperson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSigner {
			rsc := (*r.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *r.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*r.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*r.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*r.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *r.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*r.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *r.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*r.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*r.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*r.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*r.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*r.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *r.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*r.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *r.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*r.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*r.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*r.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*r.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*r.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPatientResourcesReferencingLink {
			rsc := (*r.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*r.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*r.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*r.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*r.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *r.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*r.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RelatedPersonPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *r.IncludedPatientResourcesReferencedByPatient {
			rsc := (*r.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*r.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedInvoiceResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedInvoiceResourcesReferencingRecipient {
			rsc := (*r.RevIncludedInvoiceResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingActor {
			rsc := (*r.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingConsentor != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingConsentor {
			rsc := (*r.RevIncludedConsentResourcesReferencingConsentor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingLink {
			rsc := (*r.RevIncludedPersonResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPersonResourcesReferencingRelatedperson != nil {
		for idx := range *r.RevIncludedPersonResourcesReferencingRelatedperson {
			rsc := (*r.RevIncludedPersonResourcesReferencingRelatedperson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSigner != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSigner {
			rsc := (*r.RevIncludedContractResourcesReferencingSigner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedGroupResourcesReferencingManagingentity != nil {
		for idx := range *r.RevIncludedGroupResourcesReferencingManagingentity {
			rsc := (*r.RevIncludedGroupResourcesReferencingManagingentity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCareTeamResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedCareTeamResourcesReferencingParticipant {
			rsc := (*r.RevIncludedCareTeamResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*r.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *r.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*r.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *r.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*r.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEncounterResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedEncounterResourcesReferencingParticipant {
			rsc := (*r.RevIncludedEncounterResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *r.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*r.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*r.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*r.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedExplanationOfBenefitResourcesReferencingPayee != nil {
		for idx := range *r.RevIncludedExplanationOfBenefitResourcesReferencingPayee {
			rsc := (*r.RevIncludedExplanationOfBenefitResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter != nil {
		for idx := range *r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter {
			rsc := (*r.RevIncludedAllergyIntoleranceResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*r.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*r.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*r.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *r.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*r.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*r.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingRecorder != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingRecorder {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingRecorder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*r.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedObservationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*r.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationStatementResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedMedicationStatementResourcesReferencingSource {
			rsc := (*r.RevIncludedMedicationStatementResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingAuthor {
			rsc := (*r.RevIncludedBasicResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *r.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*r.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingAsserter != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingAsserter {
			rsc := (*r.RevIncludedConditionResourcesReferencingAsserter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingAttester != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingAttester {
			rsc := (*r.RevIncludedCompositionResourcesReferencingAttester)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPatientResourcesReferencingLink != nil {
		for idx := range *r.RevIncludedPatientResourcesReferencingLink {
			rsc := (*r.RevIncludedPatientResourcesReferencingLink)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingPayor != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingPayor {
			rsc := (*r.RevIncludedCoverageResourcesReferencingPayor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingSubscriber != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingSubscriber {
			rsc := (*r.RevIncludedCoverageResourcesReferencingSubscriber)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCoverageResourcesReferencingPolicyholder != nil {
		for idx := range *r.RevIncludedCoverageResourcesReferencingPolicyholder {
			rsc := (*r.RevIncludedCoverageResourcesReferencingPolicyholder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSource {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *r.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*r.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClaimResourcesReferencingPayee != nil {
		for idx := range *r.RevIncludedClaimResourcesReferencingPayee {
			rsc := (*r.RevIncludedClaimResourcesReferencingPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
