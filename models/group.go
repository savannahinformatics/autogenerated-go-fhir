package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Group ...
type Group struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active         *bool                          `bson:"active,omitempty" json:"active,omitempty"`
	Type           string                         `bson:"type,omitempty" json:"type,omitempty"`
	Actual         *bool                          `bson:"actual,omitempty" json:"actual,omitempty"`
	Code           *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Name           string                         `bson:"name,omitempty" json:"name,omitempty"`
	Quantity       *uint32                        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ManagingEntity *Reference                     `bson:"managingEntity,omitempty" json:"managingEntity,omitempty"`
	Characteristic []GroupCharacteristicComponent `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Member         []GroupMemberComponent         `bson:"member,omitempty" json:"member,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Group) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Group"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "group" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type group Group

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Group) UnmarshalJSON(data []byte) (err error) {
	x2 := group{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Group(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Group) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Group"
	} else if x.ResourceType != "Group" {
		return fmt.Errorf("Expected resourceType to be Group, instead received %s", x.ResourceType)
	}
	return nil
}

// GroupCharacteristicComponent ...
type GroupCharacteristicComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueReference       *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Exclude              *bool            `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period          `bson:"period,omitempty" json:"period,omitempty"`
}

// GroupMemberComponent ...
type GroupMemberComponent struct {
	BackboneElement `bson:",inline"`
	Entity          *Reference `bson:"entity,omitempty" json:"entity,omitempty"`
	Period          *Period    `bson:"period,omitempty" json:"period,omitempty"`
	Inactive        *bool      `bson:"inactive,omitempty" json:"inactive,omitempty"`
}

// GroupPlus ...
type GroupPlus struct {
	Group                     `bson:",inline"`
	GroupPlusRelatedResources `bson:",inline"`
}

// GroupPlusRelatedResources ...
type GroupPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByManagingentity                *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByManagingentity,omitempty"`
	IncludedOrganizationResourcesReferencedByManagingentity                *[]Organization               `bson:"_includedOrganizationResourcesReferencedByManagingentity,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByManagingentity            *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByManagingentity,omitempty"`
	IncludedRelatedPersonResourcesReferencedByManagingentity               *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByManagingentity,omitempty"`
	IncludedPractitionerResourcesReferencedByMember                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByMember,omitempty"`
	IncludedGroupResourcesReferencedByMember                               *[]Group                      `bson:"_includedGroupResourcesReferencedByMember,omitempty"`
	IncludedDeviceResourcesReferencedByMember                              *[]Device                     `bson:"_includedDeviceResourcesReferencedByMember,omitempty"`
	IncludedMedicationResourcesReferencedByMember                          *[]Medication                 `bson:"_includedMedicationResourcesReferencedByMember,omitempty"`
	IncludedPatientResourcesReferencedByMember                             *[]Patient                    `bson:"_includedPatientResourcesReferencedByMember,omitempty"`
	IncludedSubstanceResourcesReferencedByMember                           *[]Substance                  `bson:"_includedSubstanceResourcesReferencedByMember,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByMember                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByMember,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedInvoiceResourcesReferencingSubject                          *[]Invoice                    `bson:"_revIncludedInvoiceResourcesReferencingSubject,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject                 *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedGoalResourcesReferencingSubject                             *[]Goal                       `bson:"_revIncludedGoalResourcesReferencingSubject,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingActor                            *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingActor,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingSubject                *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingSubject                   *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingSubject                   *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingSubject,omitempty"`
	RevIncludedGroupResourcesReferencingMember                             *[]Group                      `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCareTeamResourcesReferencingSubject                         *[]CareTeam                   `bson:"_revIncludedCareTeamResourcesReferencingSubject,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingSubject                     *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingSubject,omitempty"`
	RevIncludedChargeItemResourcesReferencingSubject                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingSubject,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingSubject                        *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingSubject                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationResourcesReferencingRecipient                  *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingRecipient,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingSubject               *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceUseStatementResourcesReferencingPatient               *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingPatient,omitempty"`
	RevIncludedRequestGroupResourcesReferencingSubject                     *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingSubject                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                         *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingSubject                         *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingSubject                        *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                             *[]List                       `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingSubject                *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                             *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubject                     *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubject,omitempty"`
	RevIncludedGuidanceResponseResourcesReferencingSubject                 *[]GuidanceResponse           `bson:"_revIncludedGuidanceResponseResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingSubject         *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingSubject,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingSubject              *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSubject             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSubject,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingSubject               *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedConditionResourcesReferencingSubject                        *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSubject               *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

// GetIncludedPractitionerResourceReferencedByManagingentity ...
func (g *GroupPlusRelatedResources) GetIncludedPractitionerResourceReferencedByManagingentity() (practitioner *Practitioner, err error) {
	if g.IncludedPractitionerResourcesReferencedByManagingentity == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*g.IncludedPractitionerResourcesReferencedByManagingentity) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*g.IncludedPractitionerResourcesReferencedByManagingentity))
	} else if len(*g.IncludedPractitionerResourcesReferencedByManagingentity) == 1 {
		practitioner = &(*g.IncludedPractitionerResourcesReferencedByManagingentity)[0]
	}
	return
}

// GetIncludedOrganizationResourceReferencedByManagingentity ...
func (g *GroupPlusRelatedResources) GetIncludedOrganizationResourceReferencedByManagingentity() (organization *Organization, err error) {
	if g.IncludedOrganizationResourcesReferencedByManagingentity == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*g.IncludedOrganizationResourcesReferencedByManagingentity) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*g.IncludedOrganizationResourcesReferencedByManagingentity))
	} else if len(*g.IncludedOrganizationResourcesReferencedByManagingentity) == 1 {
		organization = &(*g.IncludedOrganizationResourcesReferencedByManagingentity)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByManagingentity ...
func (g *GroupPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByManagingentity() (practitionerRole *PractitionerRole, err error) {
	if g.IncludedPractitionerRoleResourcesReferencedByManagingentity == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*g.IncludedPractitionerRoleResourcesReferencedByManagingentity) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*g.IncludedPractitionerRoleResourcesReferencedByManagingentity))
	} else if len(*g.IncludedPractitionerRoleResourcesReferencedByManagingentity) == 1 {
		practitionerRole = &(*g.IncludedPractitionerRoleResourcesReferencedByManagingentity)[0]
	}
	return
}

// GetIncludedRelatedPersonResourceReferencedByManagingentity ...
func (g *GroupPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByManagingentity() (relatedPerson *RelatedPerson, err error) {
	if g.IncludedRelatedPersonResourcesReferencedByManagingentity == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*g.IncludedRelatedPersonResourcesReferencedByManagingentity) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*g.IncludedRelatedPersonResourcesReferencedByManagingentity))
	} else if len(*g.IncludedRelatedPersonResourcesReferencedByManagingentity) == 1 {
		relatedPerson = &(*g.IncludedRelatedPersonResourcesReferencedByManagingentity)[0]
	}
	return
}

// GetIncludedPractitionerResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedPractitionerResourceReferencedByMember() (practitioner *Practitioner, err error) {
	if g.IncludedPractitionerResourcesReferencedByMember == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*g.IncludedPractitionerResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*g.IncludedPractitionerResourcesReferencedByMember))
	} else if len(*g.IncludedPractitionerResourcesReferencedByMember) == 1 {
		practitioner = &(*g.IncludedPractitionerResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedGroupResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedGroupResourceReferencedByMember() (group *Group, err error) {
	if g.IncludedGroupResourcesReferencedByMember == nil {
		err = errors.New("Included groups not requested")
	} else if len(*g.IncludedGroupResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*g.IncludedGroupResourcesReferencedByMember))
	} else if len(*g.IncludedGroupResourcesReferencedByMember) == 1 {
		group = &(*g.IncludedGroupResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedDeviceResourceReferencedByMember() (device *Device, err error) {
	if g.IncludedDeviceResourcesReferencedByMember == nil {
		err = errors.New("Included devices not requested")
	} else if len(*g.IncludedDeviceResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*g.IncludedDeviceResourcesReferencedByMember))
	} else if len(*g.IncludedDeviceResourcesReferencedByMember) == 1 {
		device = &(*g.IncludedDeviceResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedMedicationResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedMedicationResourceReferencedByMember() (medication *Medication, err error) {
	if g.IncludedMedicationResourcesReferencedByMember == nil {
		err = errors.New("Included medications not requested")
	} else if len(*g.IncludedMedicationResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*g.IncludedMedicationResourcesReferencedByMember))
	} else if len(*g.IncludedMedicationResourcesReferencedByMember) == 1 {
		medication = &(*g.IncludedMedicationResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedPatientResourceReferencedByMember() (patient *Patient, err error) {
	if g.IncludedPatientResourcesReferencedByMember == nil {
		err = errors.New("Included patients not requested")
	} else if len(*g.IncludedPatientResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*g.IncludedPatientResourcesReferencedByMember))
	} else if len(*g.IncludedPatientResourcesReferencedByMember) == 1 {
		patient = &(*g.IncludedPatientResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedSubstanceResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedSubstanceResourceReferencedByMember() (substance *Substance, err error) {
	if g.IncludedSubstanceResourcesReferencedByMember == nil {
		err = errors.New("Included substances not requested")
	} else if len(*g.IncludedSubstanceResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*g.IncludedSubstanceResourcesReferencedByMember))
	} else if len(*g.IncludedSubstanceResourcesReferencedByMember) == 1 {
		substance = &(*g.IncludedSubstanceResourcesReferencedByMember)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByMember ...
func (g *GroupPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByMember() (practitionerRole *PractitionerRole, err error) {
	if g.IncludedPractitionerRoleResourcesReferencedByMember == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*g.IncludedPractitionerRoleResourcesReferencedByMember) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*g.IncludedPractitionerRoleResourcesReferencedByMember))
	} else if len(*g.IncludedPractitionerRoleResourcesReferencedByMember) == 1 {
		practitionerRole = &(*g.IncludedPractitionerRoleResourcesReferencedByMember)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (g *GroupPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if g.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *g.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedInvoiceResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingSubject() (invoices []Invoice, err error) {
	if g.RevIncludedInvoiceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *g.RevIncludedInvoiceResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if g.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *g.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if g.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *g.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if g.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *g.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (g *GroupPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if g.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *g.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (g *GroupPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedGoalResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedGoalResourcesReferencingSubject() (goals []Goal, err error) {
	if g.RevIncludedGoalResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded goals not requested")
	} else {
		goals = *g.RevIncludedGoalResourcesReferencingSubject
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (g *GroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedConsentResourcesReferencingActor ...
func (g *GroupPlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if g.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *g.RevIncludedConsentResourcesReferencingActor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *g.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (g *GroupPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if g.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *g.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (g *GroupPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingSubject() (serviceRequests []ServiceRequest, err error) {
	if g.RevIncludedServiceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *g.RevIncludedServiceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (g *GroupPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if g.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *g.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if g.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *g.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedRiskAssessmentResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingSubject() (riskAssessments []RiskAssessment, err error) {
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *g.RevIncludedRiskAssessmentResourcesReferencingSubject
	}
	return
}

// GetRevIncludedGroupResourcesReferencingMember ...
func (g *GroupPlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if g.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *g.RevIncludedGroupResourcesReferencingMember
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (g *GroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (g *GroupPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCareTeamResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedCareTeamResourcesReferencingSubject() (careTeams []CareTeam, err error) {
	if g.RevIncludedCareTeamResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded careTeams not requested")
	} else {
		careTeams = *g.RevIncludedCareTeamResourcesReferencingSubject
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (g *GroupPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if g.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *g.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingSubject() (imagingStudies []ImagingStudy, err error) {
	if g.RevIncludedImagingStudyResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *g.RevIncludedImagingStudyResourcesReferencingSubject
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingSubject() (chargeItems []ChargeItem, err error) {
	if g.RevIncludedChargeItemResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *g.RevIncludedChargeItemResourcesReferencingSubject
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (g *GroupPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingSubject() (encounters []Encounter, err error) {
	if g.RevIncludedEncounterResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *g.RevIncludedEncounterResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSubject() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingRecipient ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if g.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *g.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (g *GroupPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (g *GroupPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if g.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *g.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (g *GroupPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if g.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *g.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceUseStatementResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingSubject() (deviceUseStatements []DeviceUseStatement, err error) {
	if g.RevIncludedDeviceUseStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *g.RevIncludedDeviceUseStatementResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDeviceUseStatementResourcesReferencingPatient ...
func (g *GroupPlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingPatient() (deviceUseStatements []DeviceUseStatement, err error) {
	if g.RevIncludedDeviceUseStatementResourcesReferencingPatient == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *g.RevIncludedDeviceUseStatementResourcesReferencingPatient
	}
	return
}

// GetRevIncludedRequestGroupResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingSubject() (requestGroups []RequestGroup, err error) {
	if g.RevIncludedRequestGroupResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *g.RevIncludedRequestGroupResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if g.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *g.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (g *GroupPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *g.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (g *GroupPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (g *GroupPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if g.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *g.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (g *GroupPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if g.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *g.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (g *GroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (g *GroupPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (g *GroupPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if g.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *g.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedSpecimenResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if g.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *g.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingSubject() (carePlans []CarePlan, err error) {
	if g.RevIncludedCarePlanResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *g.RevIncludedCarePlanResourcesReferencingSubject
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingSubject() (procedures []Procedure, err error) {
	if g.RevIncludedProcedureResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *g.RevIncludedProcedureResourcesReferencingSubject
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (g *GroupPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedListResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMedicationRequestResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingSubject() (medicationRequests []MedicationRequest, err error) {
	if g.RevIncludedMedicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *g.RevIncludedMedicationRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMediaResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if g.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *g.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if g.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *g.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if g.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *g.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if g.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *g.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if g.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *g.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedFlagResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if g.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *g.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubject() (adverseEvents []AdverseEvent, err error) {
	if g.RevIncludedAdverseEventResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *g.RevIncludedAdverseEventResourcesReferencingSubject
	}
	return
}

// GetRevIncludedGuidanceResponseResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedGuidanceResponseResourcesReferencingSubject() (guidanceResponses []GuidanceResponse, err error) {
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded guidanceResponses not requested")
	} else {
		guidanceResponses = *g.RevIncludedGuidanceResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedObservationResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if g.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *g.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (g *GroupPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if g.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *g.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedMedicationAdministrationResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingSubject() (medicationAdministrations []MedicationAdministration, err error) {
	if g.RevIncludedMedicationAdministrationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *g.RevIncludedMedicationAdministrationResourcesReferencingSubject
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (g *GroupPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if g.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *g.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedMedicationStatementResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingSubject() (medicationStatements []MedicationStatement, err error) {
	if g.RevIncludedMedicationStatementResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *g.RevIncludedMedicationStatementResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSubject() (communicationRequests []CommunicationRequest, err error) {
	if g.RevIncludedCommunicationRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *g.RevIncludedCommunicationRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *g.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingRecipient ...
func (g *GroupPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *g.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if g.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *g.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingSubject() (medicationDispenses []MedicationDispense, err error) {
	if g.RevIncludedMedicationDispenseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *g.RevIncludedMedicationDispenseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *g.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if g.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *g.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if g.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *g.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if g.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *g.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if g.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *g.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (g *GroupPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if g.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *g.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (g *GroupPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if g.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *g.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (g *GroupPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if g.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *g.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedConditionResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedConditionResourcesReferencingSubject() (conditions []Condition, err error) {
	if g.RevIncludedConditionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *g.RevIncludedConditionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (g *GroupPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (g *GroupPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *g.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSubject ...
func (g *GroupPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSubject() (clinicalImpressions []ClinicalImpression, err error) {
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *g.RevIncludedClinicalImpressionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (g *GroupPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (g *GroupPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (g *GroupPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByManagingentity {
			rsc := (*g.IncludedPractitionerResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedOrganizationResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedOrganizationResourcesReferencedByManagingentity {
			rsc := (*g.IncludedOrganizationResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerRoleResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedPractitionerRoleResourcesReferencedByManagingentity {
			rsc := (*g.IncludedPractitionerRoleResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedRelatedPersonResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedRelatedPersonResourcesReferencedByManagingentity {
			rsc := (*g.IncludedRelatedPersonResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedGroupResourcesReferencedByMember != nil {
		for idx := range *g.IncludedGroupResourcesReferencedByMember {
			rsc := (*g.IncludedGroupResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedDeviceResourcesReferencedByMember {
			rsc := (*g.IncludedDeviceResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for idx := range *g.IncludedMedicationResourcesReferencedByMember {
			rsc := (*g.IncludedMedicationResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByMember {
			rsc := (*g.IncludedPatientResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedSubstanceResourcesReferencedByMember {
			rsc := (*g.IncludedSubstanceResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerRoleResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerRoleResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerRoleResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (g *GroupPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *g.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*g.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedInvoiceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedInvoiceResourcesReferencingSubject {
			rsc := (*g.RevIncludedInvoiceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*g.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingActor {
			rsc := (*g.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*g.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*g.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*g.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *g.RevIncludedGroupResourcesReferencingMember {
			rsc := (*g.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*g.RevIncludedCareTeamResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *g.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*g.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*g.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*g.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*g.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*g.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *g.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*g.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *g.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*g.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *g.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*g.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*g.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*g.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*g.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*g.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedListResourcesReferencingSubject {
			rsc := (*g.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*g.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*g.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*g.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*g.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*g.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*g.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*g.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*g.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*g.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*g.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*g.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (g *GroupPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.IncludedPractitionerResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByManagingentity {
			rsc := (*g.IncludedPractitionerResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedOrganizationResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedOrganizationResourcesReferencedByManagingentity {
			rsc := (*g.IncludedOrganizationResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerRoleResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedPractitionerRoleResourcesReferencedByManagingentity {
			rsc := (*g.IncludedPractitionerRoleResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedRelatedPersonResourcesReferencedByManagingentity != nil {
		for idx := range *g.IncludedRelatedPersonResourcesReferencedByManagingentity {
			rsc := (*g.IncludedRelatedPersonResourcesReferencedByManagingentity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedGroupResourcesReferencedByMember != nil {
		for idx := range *g.IncludedGroupResourcesReferencedByMember {
			rsc := (*g.IncludedGroupResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedDeviceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedDeviceResourcesReferencedByMember {
			rsc := (*g.IncludedDeviceResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedMedicationResourcesReferencedByMember != nil {
		for idx := range *g.IncludedMedicationResourcesReferencedByMember {
			rsc := (*g.IncludedMedicationResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPatientResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPatientResourcesReferencedByMember {
			rsc := (*g.IncludedPatientResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedSubstanceResourcesReferencedByMember != nil {
		for idx := range *g.IncludedSubstanceResourcesReferencedByMember {
			rsc := (*g.IncludedSubstanceResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.IncludedPractitionerRoleResourcesReferencedByMember != nil {
		for idx := range *g.IncludedPractitionerRoleResourcesReferencedByMember {
			rsc := (*g.IncludedPractitionerRoleResourcesReferencedByMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *g.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*g.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedInvoiceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedInvoiceResourcesReferencingSubject {
			rsc := (*g.RevIncludedInvoiceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*g.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*g.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGoalResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGoalResourcesReferencingSubject {
			rsc := (*g.RevIncludedGoalResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingData {
			rsc := (*g.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *g.RevIncludedConsentResourcesReferencingActor {
			rsc := (*g.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*g.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *g.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*g.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*g.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*g.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedContractResourcesReferencingSubject {
			rsc := (*g.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedRiskAssessmentResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRiskAssessmentResourcesReferencingSubject {
			rsc := (*g.RevIncludedRiskAssessmentResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *g.RevIncludedGroupResourcesReferencingMember {
			rsc := (*g.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*g.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCareTeamResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCareTeamResourcesReferencingSubject {
			rsc := (*g.RevIncludedCareTeamResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *g.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*g.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*g.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedChargeItemResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedChargeItemResourcesReferencingSubject {
			rsc := (*g.RevIncludedChargeItemResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEncounterResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedEncounterResourcesReferencingSubject {
			rsc := (*g.RevIncludedEncounterResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*g.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *g.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*g.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceUseStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceUseStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceUseStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceUseStatementResourcesReferencingPatient != nil {
		for idx := range *g.RevIncludedDeviceUseStatementResourcesReferencingPatient {
			rsc := (*g.RevIncludedDeviceUseStatementResourcesReferencingPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedRequestGroupResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedRequestGroupResourcesReferencingSubject {
			rsc := (*g.RevIncludedRequestGroupResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *g.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*g.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*g.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *g.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*g.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*g.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*g.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*g.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*g.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*g.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCarePlanResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCarePlanResourcesReferencingSubject {
			rsc := (*g.RevIncludedCarePlanResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedProcedureResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedProcedureResourcesReferencingSubject {
			rsc := (*g.RevIncludedProcedureResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *g.RevIncludedListResourcesReferencingItem {
			rsc := (*g.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedListResourcesReferencingSubject {
			rsc := (*g.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*g.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*g.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*g.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedAdverseEventResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedAdverseEventResourcesReferencingSubject {
			rsc := (*g.RevIncludedAdverseEventResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedGuidanceResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedGuidanceResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedGuidanceResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*g.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *g.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*g.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationAdministrationResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationAdministrationResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationAdministrationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*g.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*g.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationStatementResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationStatementResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationStatementResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingSubject {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *g.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*g.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*g.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedMedicationDispenseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedMedicationDispenseResourcesReferencingSubject {
			rsc := (*g.RevIncludedMedicationDispenseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*g.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *g.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*g.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *g.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*g.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*g.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedConditionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedConditionResourcesReferencingSubject {
			rsc := (*g.RevIncludedConditionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*g.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *g.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*g.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*g.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*g.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSubject != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSubject {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*g.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*g.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
