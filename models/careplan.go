package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CarePlan ... // TODO Write proper comment
type CarePlan struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []Canonical                 `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                    `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference                 `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf                []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                      `bson:"status,omitempty" json:"status,omitempty"`
	Intent                string                      `bson:"intent,omitempty" json:"intent,omitempty"`
	Category              []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Title                 string                      `bson:"title,omitempty" json:"title,omitempty"`
	Description           string                      `bson:"description,omitempty" json:"description,omitempty"`
	Subject               *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Period                *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Created               *FHIRDateTime               `bson:"created,omitempty" json:"created,omitempty"`
	Author                *Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Contributor           []Reference                 `bson:"contributor,omitempty" json:"contributor,omitempty"`
	CareTeam              []Reference                 `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Addresses             []Reference                 `bson:"addresses,omitempty" json:"addresses,omitempty"`
	SupportingInfo        []Reference                 `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Goal                  []Reference                 `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity              []CarePlanActivityComponent `bson:"activity,omitempty" json:"activity,omitempty"`
	Note                  []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *CarePlan) MarshalJSON() ([]byte, error) {
	x.ResourceType = "CarePlan"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "carePlan" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type carePlan CarePlan

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *CarePlan) UnmarshalJSON(data []byte) (err error) {
	x2 := carePlan{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CarePlan(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CarePlan) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CarePlan"
	} else if x.ResourceType != "CarePlan" {
		return fmt.Errorf("Expected resourceType to be CarePlan, instead received %s", x.ResourceType)
	}
	return nil
}

// CarePlanActivityComponent ... // TODO Write proper comment
type CarePlanActivityComponent struct {
	BackboneElement        `bson:",inline"`
	OutcomeCodeableConcept []CodeableConcept                `bson:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference                      `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference              *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

// CarePlanActivityDetailComponent ... // TODO Write proper comment
type CarePlanActivityDetailComponent struct {
	BackboneElement        `bson:",inline"`
	Kind                   string            `bson:"kind,omitempty" json:"kind,omitempty"`
	InstantiatesCanonical  []Canonical       `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string          `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Code                   *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference       `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                 string            `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason           *CodeableConcept  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	DoNotPerform           *bool             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	ScheduledTiming        *Timing           `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period           `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        string            `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	Performer              []Reference       `bson:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference        `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity         `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description            string            `bson:"description,omitempty" json:"description,omitempty"`
}

// CarePlanPlus ... // TODO Write proper comment
type CarePlanPlus struct {
	CarePlan                     `bson:",inline"`
	CarePlanPlusRelatedResources `bson:",inline"`
}

// CarePlanPlusRelatedResources ... // TODO Write proper comment
type CarePlanPlusRelatedResources struct {
	IncludedCareTeamResourcesReferencedByCareteam                          *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByCareteam,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                         *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByPerformer                *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGoalResourcesReferencedByGoal                                  *[]Goal                       `bson:"_includedGoalResourcesReferencedByGoal,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedCarePlanResourcesReferencedByReplaces                          *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByReplaces,omitempty"`
	IncludedQuestionnaireResourcesReferencedByInstantiatescanonical        *[]Questionnaire              `bson:"_includedQuestionnaireResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedMeasureResourcesReferencedByInstantiatescanonical              *[]Measure                    `bson:"_includedMeasureResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical       *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical  *[]OperationDefinition        `bson:"_includedOperationDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical   *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedCarePlanResourcesReferencedByPartof                            *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedAppointmentResourcesReferencedByActivityreference              *[]Appointment                `bson:"_includedAppointmentResourcesReferencedByActivityreference,omitempty"`
	IncludedMedicationRequestResourcesReferencedByActivityreference        *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedTaskResourcesReferencedByActivityreference                     *[]Task                       `bson:"_includedTaskResourcesReferencedByActivityreference,omitempty"`
	IncludedNutritionOrderResourcesReferencedByActivityreference           *[]NutritionOrder             `bson:"_includedNutritionOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedRequestGroupResourcesReferencedByActivityreference             *[]RequestGroup               `bson:"_includedRequestGroupResourcesReferencedByActivityreference,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByActivityreference       *[]VisionPrescription         `bson:"_includedVisionPrescriptionResourcesReferencedByActivityreference,omitempty"`
	IncludedDeviceRequestResourcesReferencedByActivityreference            *[]DeviceRequest              `bson:"_includedDeviceRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedServiceRequestResourcesReferencedByActivityreference           *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByActivityreference     *[]CommunicationRequest       `bson:"_includedCommunicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedConditionResourcesReferencedByCondition                        *[]Condition                  `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingReplaces                        *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingReplaces,omitempty"`
	RevIncludedCarePlanResourcesReferencingPartof                          *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPartof,omitempty"`
	RevIncludedCarePlanResourcesReferencingBasedon                         *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingBasedon,omitempty"`
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

// GetIncludedCareTeamResourcesReferencedByCareteam ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByCareteam() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByCareteam == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByCareteam
	}
	return
}

// GetIncludedPractitionerResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

// GetIncludedOrganizationResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

// GetIncludedCareTeamResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByPerformer() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByPerformer
	}
	return
}

// GetIncludedDeviceResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedDeviceResourcesReferencedByPerformer() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByPerformer
	}
	return
}

// GetIncludedPatientResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

// GetIncludedHealthcareServiceResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByPerformer() (healthcareServices []HealthcareService, err error) {
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *c.IncludedHealthcareServiceResourcesReferencedByPerformer
	}
	return
}

// GetIncludedPractitionerRoleResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByPerformer() (practitionerRoles []PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *c.IncludedPractitionerRoleResourcesReferencedByPerformer
	}
	return
}

// GetIncludedRelatedPersonResourcesReferencedByPerformer ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

// GetIncludedGoalResourcesReferencedByGoal ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedGoalResourcesReferencedByGoal() (goals []Goal, err error) {
	if c.IncludedGoalResourcesReferencedByGoal == nil {
		err = errors.New("Included goals not requested")
	} else {
		goals = *c.IncludedGoalResourcesReferencedByGoal
	}
	return
}

// GetIncludedGroupResourceReferencedBySubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedBySubject))
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedBySubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedCarePlanResourcesReferencedByReplaces ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByReplaces() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByReplaces == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByReplaces
	}
	return
}

// GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical() (questionnaires []Questionnaire, err error) {
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included questionnaires not requested")
	} else {
		questionnaires = *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedMeasureResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedMeasureResourcesReferencedByInstantiatescanonical() (measures []Measure, err error) {
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included measures not requested")
	} else {
		measures = *c.IncludedMeasureResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical() (operationDefinitions []OperationDefinition, err error) {
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included operationDefinitions not requested")
	} else {
		operationDefinitions = *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedCarePlanResourcesReferencedByPartof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByPartof() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByPartof == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByPartof
	}
	return
}

// GetIncludedEncounterResourceReferencedByEncounter ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

// GetIncludedAppointmentResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedAppointmentResourceReferencedByActivityreference() (appointment *Appointment, err error) {
	if c.IncludedAppointmentResourcesReferencedByActivityreference == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*c.IncludedAppointmentResourcesReferencedByActivityreference))
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) == 1 {
		appointment = &(*c.IncludedAppointmentResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedMedicationRequestResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedMedicationRequestResourceReferencedByActivityreference() (medicationRequest *MedicationRequest, err error) {
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included medicationrequests not requested")
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationRequest, but found %d", len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) == 1 {
		medicationRequest = &(*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedTaskResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedTaskResourceReferencedByActivityreference() (task *Task, err error) {
	if c.IncludedTaskResourcesReferencedByActivityreference == nil {
		err = errors.New("Included tasks not requested")
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 task, but found %d", len(*c.IncludedTaskResourcesReferencedByActivityreference))
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) == 1 {
		task = &(*c.IncludedTaskResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedNutritionOrderResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedNutritionOrderResourceReferencedByActivityreference() (nutritionOrder *NutritionOrder, err error) {
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included nutritionorders not requested")
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 nutritionOrder, but found %d", len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) == 1 {
		nutritionOrder = &(*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedRequestGroupResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedRequestGroupResourceReferencedByActivityreference() (requestGroup *RequestGroup, err error) {
	if c.IncludedRequestGroupResourcesReferencedByActivityreference == nil {
		err = errors.New("Included requestgroups not requested")
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 requestGroup, but found %d", len(*c.IncludedRequestGroupResourcesReferencedByActivityreference))
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) == 1 {
		requestGroup = &(*c.IncludedRequestGroupResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedVisionPrescriptionResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedVisionPrescriptionResourceReferencedByActivityreference() (visionPrescription *VisionPrescription, err error) {
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference == nil {
		err = errors.New("Included visionprescriptions not requested")
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 visionPrescription, but found %d", len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference))
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) == 1 {
		visionPrescription = &(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedDeviceRequestResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedDeviceRequestResourceReferencedByActivityreference() (deviceRequest *DeviceRequest, err error) {
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included devicerequests not requested")
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceRequest, but found %d", len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) == 1 {
		deviceRequest = &(*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedServiceRequestResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedServiceRequestResourceReferencedByActivityreference() (serviceRequest *ServiceRequest, err error) {
	if c.IncludedServiceRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included servicerequests not requested")
	} else if len(*c.IncludedServiceRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 serviceRequest, but found %d", len(*c.IncludedServiceRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedServiceRequestResourcesReferencedByActivityreference) == 1 {
		serviceRequest = &(*c.IncludedServiceRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedCommunicationRequestResourceReferencedByActivityreference ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCommunicationRequestResourceReferencedByActivityreference() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) == 1 {
		communicationRequest = &(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

// GetIncludedConditionResourcesReferencedByCondition ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if c.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResourcesReferencedByCondition
	}
	return
}

// GetIncludedCarePlanResourcesReferencedByBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *c.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingBasedon() (serviceRequests []ServiceRequest, err error) {
	if c.RevIncludedServiceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *c.RevIncludedServiceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if c.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *c.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if c.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *c.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingReplaces ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingReplaces() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingReplaces
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingPartof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPartof() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingBasedon() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingBasedon() (procedures []Procedure, err error) {
	if c.RevIncludedProcedureResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *c.RevIncludedProcedureResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedMediaResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedMediaResourcesReferencingBasedon() (media []Media, err error) {
	if c.RevIncludedMediaResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *c.RevIncludedMediaResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *c.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*c.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*c.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedServiceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedServiceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedServiceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMediaResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedMediaResourcesReferencingBasedon {
			rsc := (*c.RevIncludedMediaResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ... // TODO Write proper comment
func (c *CarePlanPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*c.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*c.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedServiceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedServiceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedServiceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMediaResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedMediaResourcesReferencingBasedon {
			rsc := (*c.RevIncludedMediaResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
