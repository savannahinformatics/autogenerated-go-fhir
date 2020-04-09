package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MedicationAdministration struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier                                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Instantiates              []string                                     `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	PartOf                    []Reference                                  `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                    string                                       `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason              []CodeableConcept                            `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category                  *CodeableConcept                             `bson:"category,omitempty" json:"category,omitempty"`
	MedicationCodeableConcept *CodeableConcept                             `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                                   `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   *Reference                                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                                   `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation     []Reference                                  `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	EffectiveDateTime         *FHIRDateTime                                `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                                      `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Performer                 []MedicationAdministrationPerformerComponent `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode                []CodeableConcept                            `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []Reference                                  `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Request                   *Reference                                   `bson:"request,omitempty" json:"request,omitempty"`
	Device                    []Reference                                  `bson:"device,omitempty" json:"device,omitempty"`
	Note                      []Annotation                                 `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosageComponent     `bson:"dosage,omitempty" json:"dosage,omitempty"`
	EventHistory              []Reference                                  `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationAdministration) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "MedicationAdministration"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "medicationAdministration" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationAdministration MedicationAdministration

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationAdministration) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationAdministration{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationAdministration(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationAdministration) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationAdministration"
	} else if x.ResourceType != "MedicationAdministration" {
		return errors.New(fmt.Sprintf("Expected resourceType to be MedicationAdministration, instead received %s", x.ResourceType))
	}
	return nil
}

type MedicationAdministrationPerformerComponent struct {
	BackboneElement `bson:",inline"`
	Function        *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

type MedicationAdministrationDosageComponent struct {
	BackboneElement    `bson:",inline"`
	Text               string           `bson:"text,omitempty" json:"text,omitempty"`
	Site               *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	Route              *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method             *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Dose               *Quantity        `bson:"dose,omitempty" json:"dose,omitempty"`
	RateRatio          *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateSimpleQuantity *Quantity        `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
}

type MedicationAdministrationPlus struct {
	MedicationAdministration                     `bson:",inline"`
	MedicationAdministrationPlusRelatedResources `bson:",inline"`
}

type MedicationAdministrationPlusRelatedResources struct {
	IncludedMedicationRequestResourcesReferencedByRequest                  *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByRequest,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedMedicationResourcesReferencedByMedication                      *[]Medication                 `bson:"_includedMedicationResourcesReferencedByMedication,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext                      *[]EpisodeOfCare              `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                          *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedDeviceResourcesReferencedByDevice                              *[]Device                     `bson:"_includedDeviceResourcesReferencedByDevice,omitempty"`
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
	RevIncludedChargeItemResourcesReferencingService                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingService,omitempty"`
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
	RevIncludedProcedureResourcesReferencingPartof                         *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPartof,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubstance                   *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubstance,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPartof                       *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPartof,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedMedicationStatementResourcesReferencingPartof               *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingPartof,omitempty"`
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
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedMedicationRequestResourceReferencedByRequest() (medicationRequest *MedicationRequest, err error) {
	if m.IncludedMedicationRequestResourcesReferencedByRequest == nil {
		err = errors.New("Included medicationrequests not requested")
	} else if len(*m.IncludedMedicationRequestResourcesReferencedByRequest) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationRequest, but found %d", len(*m.IncludedMedicationRequestResourcesReferencedByRequest))
	} else if len(*m.IncludedMedicationRequestResourcesReferencedByRequest) == 1 {
		medicationRequest = &(*m.IncludedMedicationRequestResourcesReferencedByRequest)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*m.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedDeviceResourceReferencedByPerformer() (device *Device, err error) {
	if m.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else if len(*m.IncludedDeviceResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*m.IncludedDeviceResourcesReferencedByPerformer))
	} else if len(*m.IncludedDeviceResourcesReferencedByPerformer) == 1 {
		device = &(*m.IncludedDeviceResourcesReferencedByPerformer)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedPatientResourceReferencedByPerformer() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPerformer))
	} else if len(*m.IncludedPatientResourcesReferencedByPerformer) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPerformer)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPerformer() (practitionerRole *PractitionerRole, err error) {
	if m.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*m.IncludedPractitionerRoleResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*m.IncludedPractitionerRoleResourcesReferencedByPerformer))
	} else if len(*m.IncludedPractitionerRoleResourcesReferencedByPerformer) == 1 {
		practitionerRole = &(*m.IncludedPractitionerRoleResourcesReferencedByPerformer)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPerformer() (relatedPerson *RelatedPerson, err error) {
	if m.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*m.IncludedRelatedPersonResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*m.IncludedRelatedPersonResourcesReferencedByPerformer))
	} else if len(*m.IncludedRelatedPersonResourcesReferencedByPerformer) == 1 {
		relatedPerson = &(*m.IncludedRelatedPersonResourcesReferencedByPerformer)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if m.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*m.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*m.IncludedGroupResourcesReferencedBySubject))
	} else if len(*m.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*m.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedBySubject))
	} else if len(*m.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedMedicationResourceReferencedByMedication() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByMedication == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByMedication))
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByMedication)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPatient))
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if m.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*m.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*m.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if m.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*m.IncludedEncounterResourcesReferencedByContext))
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*m.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedDeviceResourcesReferencedByDevice() (devices []Device, err error) {
	if m.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *m.IncludedDeviceResourcesReferencedByDevice
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *m.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if m.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *m.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if m.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *m.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPartof() (procedures []Procedure, err error) {
	if m.RevIncludedProcedureResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *m.RevIncludedProcedureResourcesReferencingPartof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if m.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *m.RevIncludedAdverseEventResourcesReferencingSubstance
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if m.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *m.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPartof() (observations []Observation, err error) {
	if m.RevIncludedObservationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *m.RevIncludedObservationResourcesReferencingPartof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPartof() (medicationStatements []MedicationStatement, err error) {
	if m.RevIncludedMedicationStatementResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *m.RevIncludedMedicationStatementResourcesReferencingPartof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *m.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if m.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *m.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedMedicationRequestResourcesReferencedByRequest != nil {
		for idx := range *m.IncludedMedicationRequestResourcesReferencedByRequest {
			rsc := (*m.IncludedMedicationRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*m.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*m.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*m.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*m.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedGroupResourcesReferencedBySubject {
			rsc := (*m.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySubject {
			rsc := (*m.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*m.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationAdministrationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *m.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*m.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*m.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *m.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*m.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*m.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (m *MedicationAdministrationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedMedicationRequestResourcesReferencedByRequest != nil {
		for idx := range *m.IncludedMedicationRequestResourcesReferencedByRequest {
			rsc := (*m.IncludedMedicationRequestResourcesReferencedByRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*m.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*m.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*m.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*m.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedGroupResourcesReferencedBySubject {
			rsc := (*m.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySubject {
			rsc := (*m.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *m.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*m.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *m.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*m.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*m.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *m.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*m.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*m.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
