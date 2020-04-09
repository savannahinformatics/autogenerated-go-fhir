package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MedicationStatement ...
type MedicationStatement struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn                   []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                    []Reference       `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                    string            `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason              []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category                  *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	MedicationCodeableConcept *CodeableConcept  `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference        `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference        `bson:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime         *FHIRDateTime     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period           `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	DateAsserted              *FHIRDateTime     `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	InformationSource         *Reference        `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	DerivedFrom               []Reference       `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	ReasonCode                []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                    []Dosage          `bson:"dosage,omitempty" json:"dosage,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *MedicationStatement) MarshalJSON() ([]byte, error) {
	x.ResourceType = "MedicationStatement"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "medicationStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationStatement MedicationStatement

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationStatement(x2)
		return x.checkResourceType()
	}
	return
}

func (x *MedicationStatement) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "MedicationStatement"
	} else if x.ResourceType != "MedicationStatement" {
		return fmt.Errorf("Expected resourceType to be MedicationStatement, instead received %s", x.ResourceType)
	}
	return nil
}

// MedicationStatementPlus ...
type MedicationStatementPlus struct {
	MedicationStatement                     `bson:",inline"`
	MedicationStatementPlusRelatedResources `bson:",inline"`
}

// MedicationStatementPlusRelatedResources ...
type MedicationStatementPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEpisodeOfCareResourcesReferencedByContext                      *[]EpisodeOfCare              `bson:"_includedEpisodeOfCareResourcesReferencedByContext,omitempty"`
	IncludedEncounterResourcesReferencedByContext                          *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByContext,omitempty"`
	IncludedMedicationResourcesReferencedByMedication                      *[]Medication                 `bson:"_includedMedicationResourcesReferencedByMedication,omitempty"`
	IncludedMedicationDispenseResourcesReferencedByPartof                  *[]MedicationDispense         `bson:"_includedMedicationDispenseResourcesReferencedByPartof,omitempty"`
	IncludedObservationResourcesReferencedByPartof                         *[]Observation                `bson:"_includedObservationResourcesReferencedByPartof,omitempty"`
	IncludedMedicationAdministrationResourcesReferencedByPartof            *[]MedicationAdministration   `bson:"_includedMedicationAdministrationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                           *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedMedicationStatementResourcesReferencedByPartof                 *[]MedicationStatement        `bson:"_includedMedicationStatementResourcesReferencedByPartof,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedOrganizationResourcesReferencedBySource                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedBySource,omitempty"`
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

// GetIncludedGroupResourceReferencedBySubject ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if m.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*m.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*m.IncludedGroupResourcesReferencedBySubject))
	} else if len(*m.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*m.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedBySubject ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedBySubject))
	} else if len(*m.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedByPatient))
	} else if len(*m.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedEpisodeOfCareResourceReferencedByContext ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedEpisodeOfCareResourceReferencedByContext() (episodeOfCare *EpisodeOfCare, err error) {
	if m.IncludedEpisodeOfCareResourcesReferencedByContext == nil {
		err = errors.New("Included episodeofcares not requested")
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 episodeOfCare, but found %d", len(*m.IncludedEpisodeOfCareResourcesReferencedByContext))
	} else if len(*m.IncludedEpisodeOfCareResourcesReferencedByContext) == 1 {
		episodeOfCare = &(*m.IncludedEpisodeOfCareResourcesReferencedByContext)[0]
	}
	return
}

// GetIncludedEncounterResourceReferencedByContext ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedEncounterResourceReferencedByContext() (encounter *Encounter, err error) {
	if m.IncludedEncounterResourcesReferencedByContext == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*m.IncludedEncounterResourcesReferencedByContext))
	} else if len(*m.IncludedEncounterResourcesReferencedByContext) == 1 {
		encounter = &(*m.IncludedEncounterResourcesReferencedByContext)[0]
	}
	return
}

// GetIncludedMedicationResourceReferencedByMedication ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedMedicationResourceReferencedByMedication() (medication *Medication, err error) {
	if m.IncludedMedicationResourcesReferencedByMedication == nil {
		err = errors.New("Included medications not requested")
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*m.IncludedMedicationResourcesReferencedByMedication))
	} else if len(*m.IncludedMedicationResourcesReferencedByMedication) == 1 {
		medication = &(*m.IncludedMedicationResourcesReferencedByMedication)[0]
	}
	return
}

// GetIncludedMedicationDispenseResourcesReferencedByPartof ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedMedicationDispenseResourcesReferencedByPartof() (medicationDispenses []MedicationDispense, err error) {
	if m.IncludedMedicationDispenseResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationDispenses not requested")
	} else {
		medicationDispenses = *m.IncludedMedicationDispenseResourcesReferencedByPartof
	}
	return
}

// GetIncludedObservationResourcesReferencedByPartof ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedObservationResourcesReferencedByPartof() (observations []Observation, err error) {
	if m.IncludedObservationResourcesReferencedByPartof == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *m.IncludedObservationResourcesReferencedByPartof
	}
	return
}

// GetIncludedMedicationAdministrationResourcesReferencedByPartof ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedMedicationAdministrationResourcesReferencedByPartof() (medicationAdministrations []MedicationAdministration, err error) {
	if m.IncludedMedicationAdministrationResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationAdministrations not requested")
	} else {
		medicationAdministrations = *m.IncludedMedicationAdministrationResourcesReferencedByPartof
	}
	return
}

// GetIncludedProcedureResourcesReferencedByPartof ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if m.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *m.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

// GetIncludedMedicationStatementResourcesReferencedByPartof ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedMedicationStatementResourcesReferencedByPartof() (medicationStatements []MedicationStatement, err error) {
	if m.IncludedMedicationStatementResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationStatements not requested")
	} else {
		medicationStatements = *m.IncludedMedicationStatementResourcesReferencedByPartof
	}
	return
}

// GetIncludedPractitionerResourceReferencedBySource ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if m.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*m.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*m.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*m.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*m.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

// GetIncludedOrganizationResourceReferencedBySource ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySource() (organization *Organization, err error) {
	if m.IncludedOrganizationResourcesReferencedBySource == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*m.IncludedOrganizationResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*m.IncludedOrganizationResourcesReferencedBySource))
	} else if len(*m.IncludedOrganizationResourcesReferencedBySource) == 1 {
		organization = &(*m.IncludedOrganizationResourcesReferencedBySource)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedBySource ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if m.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*m.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*m.IncludedPatientResourcesReferencedBySource))
	} else if len(*m.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*m.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedBySource ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedBySource() (practitionerRole *PractitionerRole, err error) {
	if m.IncludedPractitionerRoleResourcesReferencedBySource == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*m.IncludedPractitionerRoleResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*m.IncludedPractitionerRoleResourcesReferencedBySource))
	} else if len(*m.IncludedPractitionerRoleResourcesReferencedBySource) == 1 {
		practitionerRole = &(*m.IncludedPractitionerRoleResourcesReferencedBySource)[0]
	}
	return
}

// GetIncludedRelatedPersonResourceReferencedBySource ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if m.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*m.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*m.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*m.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*m.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *m.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if m.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *m.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if m.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *m.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *m.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if m.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *m.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if m.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *m.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if m.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *m.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if m.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *m.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if m.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *m.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if m.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *m.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *m.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingSubstance ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if m.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *m.RevIncludedAdverseEventResourcesReferencingSubstance
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if m.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *m.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingPartof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPartof() (observations []Observation, err error) {
	if m.RevIncludedObservationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *m.RevIncludedObservationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if m.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *m.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedMedicationStatementResourcesReferencingPartof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPartof() (medicationStatements []MedicationStatement, err error) {
	if m.RevIncludedMedicationStatementResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *m.RevIncludedMedicationStatementResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *m.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if m.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *m.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if m.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *m.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedGroupResourcesReferencedBySubject {
			rsc := (*m.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySubject {
			rsc := (*m.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationDispenseResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationDispenseResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationDispenseResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedObservationResourcesReferencedByPartof {
			rsc := (*m.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*m.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationStatementResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationStatementResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationStatementResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*m.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*m.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySource {
			rsc := (*m.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*m.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (m *MedicationStatementPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *m.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*m.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*m.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (m *MedicationStatementPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedGroupResourcesReferencedBySubject {
			rsc := (*m.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySubject {
			rsc := (*m.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *m.IncludedPatientResourcesReferencedByPatient {
			rsc := (*m.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedEpisodeOfCareResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEpisodeOfCareResourcesReferencedByContext {
			rsc := (*m.IncludedEpisodeOfCareResourcesReferencedByContext)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedEncounterResourcesReferencedByContext != nil {
		for idx := range *m.IncludedEncounterResourcesReferencedByContext {
			rsc := (*m.IncludedEncounterResourcesReferencedByContext)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationResourcesReferencedByMedication != nil {
		for idx := range *m.IncludedMedicationResourcesReferencedByMedication {
			rsc := (*m.IncludedMedicationResourcesReferencedByMedication)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationDispenseResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationDispenseResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationDispenseResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedObservationResourcesReferencedByPartof {
			rsc := (*m.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*m.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedMedicationStatementResourcesReferencedByPartof != nil {
		for idx := range *m.IncludedMedicationStatementResourcesReferencedByPartof {
			rsc := (*m.IncludedMedicationStatementResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*m.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *m.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*m.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPatientResourcesReferencedBySource {
			rsc := (*m.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *m.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*m.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *m.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*m.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*m.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*m.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *m.RevIncludedConsentResourcesReferencingData {
			rsc := (*m.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*m.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *m.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*m.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*m.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*m.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedContractResourcesReferencingSubject {
			rsc := (*m.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*m.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *m.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*m.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*m.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *m.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*m.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *m.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*m.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*m.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *m.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*m.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*m.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*m.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*m.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*m.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *m.RevIncludedListResourcesReferencingItem {
			rsc := (*m.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *m.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*m.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*m.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*m.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*m.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*m.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *m.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*m.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *m.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*m.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*m.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *m.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*m.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *m.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*m.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *m.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*m.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*m.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *m.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*m.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*m.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*m.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*m.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*m.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
