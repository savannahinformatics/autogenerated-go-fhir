package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type QuestionnaireResponse struct {
	DomainResource `bson:",inline"`
	Identifier     *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn        []Reference                          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf         []Reference                          `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Questionnaire  *Canonical                           `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status         string                               `bson:"status,omitempty" json:"status,omitempty"`
	Subject        *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter      *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Authored       *FHIRDateTime                        `bson:"authored,omitempty" json:"authored,omitempty"`
	Author         *Reference                           `bson:"author,omitempty" json:"author,omitempty"`
	Source         *Reference                           `bson:"source,omitempty" json:"source,omitempty"`
	Item           []QuestionnaireResponseItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "QuestionnaireResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "questionnaireResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaireResponse QuestionnaireResponse

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *QuestionnaireResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaireResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = QuestionnaireResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *QuestionnaireResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "QuestionnaireResponse"
	} else if x.ResourceType != "QuestionnaireResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be QuestionnaireResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type QuestionnaireResponseItemComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                                     `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition      string                                     `bson:"definition,omitempty" json:"definition,omitempty"`
	Text            string                                     `bson:"text,omitempty" json:"text,omitempty"`
	Answer          []QuestionnaireResponseItemAnswerComponent `bson:"answer,omitempty" json:"answer,omitempty"`
	Item            []QuestionnaireResponseItemComponent       `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswerComponent struct {
	BackboneElement `bson:",inline"`
	ValueBoolean    *bool                                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64                             `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32                               `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime                        `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueTime       *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string                               `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding                              `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference                           `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Item            []QuestionnaireResponseItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponsePlus struct {
	QuestionnaireResponse                     `bson:",inline"`
	QuestionnaireResponsePlusRelatedResources `bson:",inline"`
}

type QuestionnaireResponsePlusRelatedResources struct {
	IncludedQuestionnaireResourcesReferencedByQuestionnaire                *[]Questionnaire              `bson:"_includedQuestionnaireResourcesReferencedByQuestionnaire,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                              *[]Device                     `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                             *[]Patient                    `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByAuthor                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor                       *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedObservationResourcesReferencedByPartof                         *[]Observation                `bson:"_includedObservationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                           *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                             *[]Patient                    `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedPractitionerRoleResourcesReferencedBySource                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedBySource,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySource                       *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedBySource,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedConsentResourcesReferencingSourcereference                  *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingSourcereference,omitempty"`
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
	RevIncludedObservationResourcesReferencingDerivedfrom                  *[]Observation                `bson:"_revIncludedObservationResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedObservationResourcesReferencingHasmember                    *[]Observation                `bson:"_revIncludedObservationResourcesReferencingHasmember,omitempty"`
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
	RevIncludedClinicalImpressionResourcesReferencingInvestigation         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedQuestionnaireResourceReferencedByQuestionnaire() (questionnaire *Questionnaire, err error) {
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire == nil {
		err = errors.New("Included questionnaires not requested")
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) > 1 {
		err = fmt.Errorf("Expected 0 or 1 questionnaire, but found %d", len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire))
	} else if len(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire) == 1 {
		questionnaire = &(*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if q.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *q.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if q.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *q.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*q.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if q.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*q.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*q.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*q.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*q.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if q.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*q.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*q.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*q.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*q.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByAuthor() (practitionerRole *PractitionerRole, err error) {
	if q.IncludedPractitionerRoleResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*q.IncludedPractitionerRoleResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*q.IncludedPractitionerRoleResourcesReferencedByAuthor))
	} else if len(*q.IncludedPractitionerRoleResourcesReferencedByAuthor) == 1 {
		practitionerRole = &(*q.IncludedPractitionerRoleResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedByPatient))
	} else if len(*q.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedObservationResourcesReferencedByPartof() (observations []Observation, err error) {
	if q.IncludedObservationResourcesReferencedByPartof == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *q.IncludedObservationResourcesReferencedByPartof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if q.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *q.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if q.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*q.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*q.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*q.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*q.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if q.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*q.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*q.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*q.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if q.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*q.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*q.IncludedPatientResourcesReferencedBySource))
	} else if len(*q.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*q.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedBySource() (practitionerRole *PractitionerRole, err error) {
	if q.IncludedPractitionerRoleResourcesReferencedBySource == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*q.IncludedPractitionerRoleResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*q.IncludedPractitionerRoleResourcesReferencedBySource))
	} else if len(*q.IncludedPractitionerRoleResourcesReferencedBySource) == 1 {
		practitionerRole = &(*q.IncludedPractitionerRoleResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if q.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*q.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*q.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*q.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *q.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConsentResourcesReferencingSourcereference() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingSourcereference == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingSourcereference
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if q.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *q.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if q.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *q.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if q.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *q.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if q.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *q.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedObservationResourcesReferencingDerivedfrom() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedObservationResourcesReferencingHasmember() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingHasmember == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingHasmember
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *q.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if q.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *q.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*q.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*q.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *q.IncludedObservationResourcesReferencedByPartof {
			rsc := (*q.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *q.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*q.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*q.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnaireResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingSourcereference != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingSourcereference {
			rsc := (*q.RevIncludedConsentResourcesReferencingSourcereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*q.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*q.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*q.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*q.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*q.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingHasmember != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingHasmember {
			rsc := (*q.RevIncludedObservationResourcesReferencingHasmember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (q *QuestionnaireResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.IncludedQuestionnaireResourcesReferencedByQuestionnaire != nil {
		for idx := range *q.IncludedQuestionnaireResourcesReferencedByQuestionnaire {
			rsc := (*q.IncludedQuestionnaireResourcesReferencedByQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*q.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *q.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*q.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*q.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*q.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*q.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*q.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *q.IncludedPatientResourcesReferencedByPatient {
			rsc := (*q.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *q.IncludedObservationResourcesReferencedByPartof {
			rsc := (*q.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *q.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*q.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *q.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*q.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPatientResourcesReferencedBySource {
			rsc := (*q.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *q.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*q.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *q.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*q.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingSourcereference != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingSourcereference {
			rsc := (*q.RevIncludedConsentResourcesReferencingSourcereference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*q.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*q.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*q.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*q.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*q.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingHasmember != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingHasmember {
			rsc := (*q.RevIncludedObservationResourcesReferencingHasmember)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
