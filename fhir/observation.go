package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Observation ...
type Observation struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn              []Reference                          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf               []Reference                          `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status               string                               `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept                    `bson:"category,omitempty" json:"category,omitempty"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Focus                []Reference                          `bson:"focus,omitempty" json:"focus,omitempty"`
	Encounter            *Reference                           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime    *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	EffectiveTiming      *Timing                              `bson:"effectiveTiming,omitempty" json:"effectiveTiming,omitempty"`
	EffectiveInstant     *FHIRDateTime                        `bson:"effectiveInstant,omitempty" json:"effectiveInstant,omitempty"`
	Issued               *FHIRDateTime                        `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                          `bson:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger         *int32                               `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept                    `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Note                 []Annotation                         `bson:"note,omitempty" json:"note,omitempty"`
	BodySite             *CodeableConcept                     `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method               *CodeableConcept                     `bson:"method,omitempty" json:"method,omitempty"`
	Specimen             *Reference                           `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device               *Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	HasMember            []Reference                          `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	DerivedFrom          []Reference                          `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Component            []ObservationComponentComponent      `bson:"component,omitempty" json:"component,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Observation) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Observation"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "observation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observation Observation

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Observation) UnmarshalJSON(data []byte) (err error) {
	x2 := observation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Observation(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Observation) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Observation"
	} else if x.ResourceType != "Observation" {
		return fmt.Errorf("Expected resourceType to be Observation, instead received %s", x.ResourceType)
	}
	return nil
}

// ObservationReferenceRangeComponent ...
type ObservationReferenceRangeComponent struct {
	BackboneElement `bson:",inline"`
	Low             *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High            *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo       []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age             *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text            string            `bson:"text,omitempty" json:"text,omitempty"`
}

// ObservationComponentComponent ...
type ObservationComponentComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept                     `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity                            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept                     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string                               `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger         *int32                               `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange           *Range                               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueTime            *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       []CodeableConcept                    `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange       []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

// ObservationPlus ...
type ObservationPlus struct {
	Observation                     `bson:",inline"`
	ObservationPlusRelatedResources `bson:",inline"`
}

// ObservationPlusRelatedResources ...
type ObservationPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                             *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                           *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedMediaResourcesReferencedByDerivedfrom                          *[]Media                      `bson:"_includedMediaResourcesReferencedByDerivedfrom,omitempty"`
	IncludedObservationResourcesReferencedByDerivedfrom                    *[]Observation                `bson:"_includedObservationResourcesReferencedByDerivedfrom,omitempty"`
	IncludedImagingStudyResourcesReferencedByDerivedfrom                   *[]ImagingStudy               `bson:"_includedImagingStudyResourcesReferencedByDerivedfrom,omitempty"`
	IncludedMolecularSequenceResourcesReferencedByDerivedfrom              *[]MolecularSequence          `bson:"_includedMolecularSequenceResourcesReferencedByDerivedfrom,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom          *[]QuestionnaireResponse      `bson:"_includedQuestionnaireResponseResourcesReferencedByDerivedfrom,omitempty"`
	IncludedDocumentReferenceResourcesReferencedByDerivedfrom              *[]DocumentReference          `bson:"_includedDocumentReferenceResourcesReferencedByDerivedfrom,omitempty"`
	IncludedImmunizationResourcesReferencedByPartof                        *[]Immunization               `bson:"_includedImmunizationResourcesReferencedByPartof,omitempty"`
	IncludedMedicationDispenseResourcesReferencedByPartof                  *[]MedicationDispense         `bson:"_includedMedicationDispenseResourcesReferencedByPartof,omitempty"`
	IncludedMedicationAdministrationResourcesReferencedByPartof            *[]MedicationAdministration   `bson:"_includedMedicationAdministrationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                           *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedImagingStudyResourcesReferencedByPartof                        *[]ImagingStudy               `bson:"_includedImagingStudyResourcesReferencedByPartof,omitempty"`
	IncludedMedicationStatementResourcesReferencedByPartof                 *[]MedicationStatement        `bson:"_includedMedicationStatementResourcesReferencedByPartof,omitempty"`
	IncludedObservationResourcesReferencedByHasmember                      *[]Observation                `bson:"_includedObservationResourcesReferencedByHasmember,omitempty"`
	IncludedMolecularSequenceResourcesReferencedByHasmember                *[]MolecularSequence          `bson:"_includedMolecularSequenceResourcesReferencedByHasmember,omitempty"`
	IncludedQuestionnaireResponseResourcesReferencedByHasmember            *[]QuestionnaireResponse      `bson:"_includedQuestionnaireResponseResourcesReferencedByHasmember,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedMedicationRequestResourcesReferencedByBasedon                  *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByBasedon,omitempty"`
	IncludedNutritionOrderResourcesReferencedByBasedon                     *[]NutritionOrder             `bson:"_includedNutritionOrderResourcesReferencedByBasedon,omitempty"`
	IncludedDeviceRequestResourcesReferencedByBasedon                      *[]DeviceRequest              `bson:"_includedDeviceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedImmunizationRecommendationResourcesReferencedByBasedon         *[]ImmunizationRecommendation `bson:"_includedImmunizationRecommendationResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimen                          *[]Specimen                   `bson:"_includedSpecimenResourcesReferencedBySpecimen,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                         *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedDeviceResourcesReferencedByDevice                              *[]Device                     `bson:"_includedDeviceResourcesReferencedByDevice,omitempty"`
	IncludedDeviceMetricResourcesReferencedByDevice                        *[]DeviceMetric               `bson:"_includedDeviceMetricResourcesReferencedByDevice,omitempty"`
	RevIncludedAppointmentResourcesReferencingReasonreference              *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingReasonreference,omitempty"`
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
	RevIncludedEncounterResourcesReferencingReasonreference                *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingReasonreference,omitempty"`
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
	RevIncludedProcedureResourcesReferencingReasonreference                *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingReasonreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedImmunizationResourcesReferencingReaction                    *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingReaction,omitempty"`
	RevIncludedImmunizationResourcesReferencingReasonreference             *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingReasonreference,omitempty"`
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
	RevIncludedMedicationStatementResourcesReferencingPartof               *[]MedicationStatement        `bson:"_revIncludedMedicationStatementResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingResult                  *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingResult,omitempty"`
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
	RevIncludedQuestionnaireResponseResourcesReferencingPartof             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingPartof,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingFindingref            *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingFindingref,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingInvestigation         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingInvestigation,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

// GetIncludedGroupResourceReferencedBySubject ...
func (o *ObservationPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if o.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*o.IncludedGroupResourcesReferencedBySubject))
	} else if len(*o.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*o.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedBySubject ...
func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*o.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedBySubject ...
func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedBySubject))
	} else if len(*o.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedLocationResourceReferencedBySubject ...
func (o *ObservationPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if o.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*o.IncludedLocationResourcesReferencedBySubject))
	} else if len(*o.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*o.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedMediaResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedMediaResourcesReferencedByDerivedfrom() (media []Media, err error) {
	if o.IncludedMediaResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included media not requested")
	} else {
		media = *o.IncludedMediaResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedObservationResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedObservationResourcesReferencedByDerivedfrom() (observations []Observation, err error) {
	if o.IncludedObservationResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *o.IncludedObservationResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedImagingStudyResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedImagingStudyResourcesReferencedByDerivedfrom() (imagingStudies []ImagingStudy, err error) {
	if o.IncludedImagingStudyResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included imagingStudies not requested")
	} else {
		imagingStudies = *o.IncludedImagingStudyResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedMolecularSequenceResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedMolecularSequenceResourcesReferencedByDerivedfrom() (molecularSequences []MolecularSequence, err error) {
	if o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included molecularSequences not requested")
	} else {
		molecularSequences = *o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedQuestionnaireResponseResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedQuestionnaireResponseResourcesReferencedByDerivedfrom() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedDocumentReferenceResourcesReferencedByDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetIncludedDocumentReferenceResourcesReferencedByDerivedfrom() (documentReferences []DocumentReference, err error) {
	if o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom == nil {
		err = errors.New("Included documentReferences not requested")
	} else {
		documentReferences = *o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom
	}
	return
}

// GetIncludedImmunizationResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedImmunizationResourcesReferencedByPartof() (immunizations []Immunization, err error) {
	if o.IncludedImmunizationResourcesReferencedByPartof == nil {
		err = errors.New("Included immunizations not requested")
	} else {
		immunizations = *o.IncludedImmunizationResourcesReferencedByPartof
	}
	return
}

// GetIncludedMedicationDispenseResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedMedicationDispenseResourcesReferencedByPartof() (medicationDispenses []MedicationDispense, err error) {
	if o.IncludedMedicationDispenseResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationDispenses not requested")
	} else {
		medicationDispenses = *o.IncludedMedicationDispenseResourcesReferencedByPartof
	}
	return
}

// GetIncludedMedicationAdministrationResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedMedicationAdministrationResourcesReferencedByPartof() (medicationAdministrations []MedicationAdministration, err error) {
	if o.IncludedMedicationAdministrationResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationAdministrations not requested")
	} else {
		medicationAdministrations = *o.IncludedMedicationAdministrationResourcesReferencedByPartof
	}
	return
}

// GetIncludedProcedureResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if o.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *o.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

// GetIncludedImagingStudyResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedImagingStudyResourcesReferencedByPartof() (imagingStudies []ImagingStudy, err error) {
	if o.IncludedImagingStudyResourcesReferencedByPartof == nil {
		err = errors.New("Included imagingStudies not requested")
	} else {
		imagingStudies = *o.IncludedImagingStudyResourcesReferencedByPartof
	}
	return
}

// GetIncludedMedicationStatementResourcesReferencedByPartof ...
func (o *ObservationPlusRelatedResources) GetIncludedMedicationStatementResourcesReferencedByPartof() (medicationStatements []MedicationStatement, err error) {
	if o.IncludedMedicationStatementResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationStatements not requested")
	} else {
		medicationStatements = *o.IncludedMedicationStatementResourcesReferencedByPartof
	}
	return
}

// GetIncludedObservationResourcesReferencedByHasmember ...
func (o *ObservationPlusRelatedResources) GetIncludedObservationResourcesReferencedByHasmember() (observations []Observation, err error) {
	if o.IncludedObservationResourcesReferencedByHasmember == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *o.IncludedObservationResourcesReferencedByHasmember
	}
	return
}

// GetIncludedMolecularSequenceResourcesReferencedByHasmember ...
func (o *ObservationPlusRelatedResources) GetIncludedMolecularSequenceResourcesReferencedByHasmember() (molecularSequences []MolecularSequence, err error) {
	if o.IncludedMolecularSequenceResourcesReferencedByHasmember == nil {
		err = errors.New("Included molecularSequences not requested")
	} else {
		molecularSequences = *o.IncludedMolecularSequenceResourcesReferencedByHasmember
	}
	return
}

// GetIncludedQuestionnaireResponseResourcesReferencedByHasmember ...
func (o *ObservationPlusRelatedResources) GetIncludedQuestionnaireResponseResourcesReferencedByHasmember() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.IncludedQuestionnaireResponseResourcesReferencedByHasmember == nil {
		err = errors.New("Included questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.IncludedQuestionnaireResponseResourcesReferencedByHasmember
	}
	return
}

// GetIncludedCarePlanResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if o.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *o.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

// GetIncludedMedicationRequestResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedMedicationRequestResourcesReferencedByBasedon() (medicationRequests []MedicationRequest, err error) {
	if o.IncludedMedicationRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included medicationRequests not requested")
	} else {
		medicationRequests = *o.IncludedMedicationRequestResourcesReferencedByBasedon
	}
	return
}

// GetIncludedNutritionOrderResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedNutritionOrderResourcesReferencedByBasedon() (nutritionOrders []NutritionOrder, err error) {
	if o.IncludedNutritionOrderResourcesReferencedByBasedon == nil {
		err = errors.New("Included nutritionOrders not requested")
	} else {
		nutritionOrders = *o.IncludedNutritionOrderResourcesReferencedByBasedon
	}
	return
}

// GetIncludedDeviceRequestResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedDeviceRequestResourcesReferencedByBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.IncludedDeviceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included deviceRequests not requested")
	} else {
		deviceRequests = *o.IncludedDeviceRequestResourcesReferencedByBasedon
	}
	return
}

// GetIncludedServiceRequestResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if o.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *o.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

// GetIncludedImmunizationRecommendationResourcesReferencedByBasedon ...
func (o *ObservationPlusRelatedResources) GetIncludedImmunizationRecommendationResourcesReferencedByBasedon() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon == nil {
		err = errors.New("Included immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ...
func (o *ObservationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*o.IncludedPatientResourcesReferencedByPatient))
	} else if len(*o.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*o.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedSpecimenResourceReferencedBySpecimen ...
func (o *ObservationPlusRelatedResources) GetIncludedSpecimenResourceReferencedBySpecimen() (specimen *Specimen, err error) {
	if o.IncludedSpecimenResourcesReferencedBySpecimen == nil {
		err = errors.New("Included specimen not requested")
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) > 1 {
		err = fmt.Errorf("Expected 0 or 1 specimen, but found %d", len(*o.IncludedSpecimenResourcesReferencedBySpecimen))
	} else if len(*o.IncludedSpecimenResourcesReferencedBySpecimen) == 1 {
		specimen = &(*o.IncludedSpecimenResourcesReferencedBySpecimen)[0]
	}
	return
}

// GetIncludedPractitionerResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if o.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *o.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

// GetIncludedOrganizationResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if o.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *o.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

// GetIncludedCareTeamResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByPerformer() (careTeams []CareTeam, err error) {
	if o.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *o.IncludedCareTeamResourcesReferencedByPerformer
	}
	return
}

// GetIncludedPatientResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if o.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *o.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

// GetIncludedPractitionerRoleResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByPerformer() (practitionerRoles []PractitionerRole, err error) {
	if o.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *o.IncludedPractitionerRoleResourcesReferencedByPerformer
	}
	return
}

// GetIncludedRelatedPersonResourcesReferencedByPerformer ...
func (o *ObservationPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if o.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *o.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

// GetIncludedEncounterResourceReferencedByEncounter ...
func (o *ObservationPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if o.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*o.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*o.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*o.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedByDevice ...
func (o *ObservationPlusRelatedResources) GetIncludedDeviceResourceReferencedByDevice() (device *Device, err error) {
	if o.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*o.IncludedDeviceResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceResourcesReferencedByDevice) == 1 {
		device = &(*o.IncludedDeviceResourcesReferencedByDevice)[0]
	}
	return
}

// GetIncludedDeviceMetricResourceReferencedByDevice ...
func (o *ObservationPlusRelatedResources) GetIncludedDeviceMetricResourceReferencedByDevice() (deviceMetric *DeviceMetric, err error) {
	if o.IncludedDeviceMetricResourcesReferencedByDevice == nil {
		err = errors.New("Included devicemetrics not requested")
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceMetric, but found %d", len(*o.IncludedDeviceMetricResourcesReferencedByDevice))
	} else if len(*o.IncludedDeviceMetricResourcesReferencedByDevice) == 1 {
		deviceMetric = &(*o.IncludedDeviceMetricResourcesReferencedByDevice)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingReasonreference ...
func (o *ObservationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingReasonreference() (appointments []Appointment, err error) {
	if o.RevIncludedAppointmentResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *o.RevIncludedAppointmentResourcesReferencingReasonreference
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (o *ObservationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *o.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if o.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *o.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (o *ObservationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if o.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *o.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *o.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (o *ObservationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if o.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *o.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (o *ObservationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if o.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *o.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (o *ObservationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if o.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *o.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingService ...
func (o *ObservationPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if o.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *o.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingReasonreference ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingReasonreference() (encounters []Encounter, err error) {
	if o.RevIncludedEncounterResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *o.RevIncludedEncounterResourcesReferencingReasonreference
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if o.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *o.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if o.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *o.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (o *ObservationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (o *ObservationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (o *ObservationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if o.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *o.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingPartof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPartof() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingPartof
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingReasonreference ...
func (o *ObservationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingReasonreference() (procedures []Procedure, err error) {
	if o.RevIncludedProcedureResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *o.RevIncludedProcedureResourcesReferencingReasonreference
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (o *ObservationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedImmunizationResourcesReferencingReaction ...
func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingReaction() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingReaction == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingReaction
	}
	return
}

// GetRevIncludedImmunizationResourcesReferencingReasonreference ...
func (o *ObservationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingReasonreference() (immunizations []Immunization, err error) {
	if o.RevIncludedImmunizationResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *o.RevIncludedImmunizationResourcesReferencingReasonreference
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *o.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (o *ObservationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingDerivedfrom() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedObservationResourcesReferencingHasmember ...
func (o *ObservationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingHasmember() (observations []Observation, err error) {
	if o.RevIncludedObservationResourcesReferencingHasmember == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *o.RevIncludedObservationResourcesReferencingHasmember
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (o *ObservationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if o.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *o.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedMedicationStatementResourcesReferencingPartof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedMedicationStatementResourcesReferencingPartof() (medicationStatements []MedicationStatement, err error) {
	if o.RevIncludedMedicationStatementResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationStatements not requested")
	} else {
		medicationStatements = *o.RevIncludedMedicationStatementResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (o *ObservationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *o.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (o *ObservationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingResult ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingResult() (diagnosticReports []DiagnosticReport, err error) {
	if o.RevIncludedDiagnosticReportResourcesReferencingResult == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *o.RevIncludedDiagnosticReportResourcesReferencingResult
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (o *ObservationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if o.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *o.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (o *ObservationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (o *ObservationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if o.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *o.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (o *ObservationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (o *ObservationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (o *ObservationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingPartof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingPartof() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingPartof
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingFindingref ...
func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingFindingref() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingFindingref
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingInvestigation ...
func (o *ObservationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (o *ObservationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (o *ObservationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedLocationResourcesReferencedBySubject {
			rsc := (*o.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMediaResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedMediaResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedMediaResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedObservationResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImagingStudyResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedImagingStudyResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedImagingStudyResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImmunizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedImmunizationResourcesReferencedByPartof {
			rsc := (*o.IncludedImmunizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationDispenseResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationDispenseResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationDispenseResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*o.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImagingStudyResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedImagingStudyResourcesReferencedByPartof {
			rsc := (*o.IncludedImagingStudyResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationStatementResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationStatementResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationStatementResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByHasmember {
			rsc := (*o.IncludedObservationResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMolecularSequenceResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedMolecularSequenceResourcesReferencedByHasmember {
			rsc := (*o.IncludedMolecularSequenceResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByHasmember {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*o.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedNutritionOrderResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedNutritionOrderResourcesReferencedByBasedon {
			rsc := (*o.IncludedNutritionOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedDeviceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedDeviceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon {
			rsc := (*o.IncludedImmunizationRecommendationResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*o.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*o.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*o.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*o.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*o.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceMetricResourcesReferencedByDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (o *ObservationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedProcedureResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReaction)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingHasmember != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingHasmember {
			rsc := (*o.RevIncludedObservationResourcesReferencingHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResult)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingPartof {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingFindingref {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingFindingref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (o *ObservationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedGroupResourcesReferencedBySubject {
			rsc := (*o.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*o.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedPatientResourcesReferencedBySubject {
			rsc := (*o.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *o.IncludedLocationResourcesReferencedBySubject {
			rsc := (*o.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMediaResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedMediaResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedMediaResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedObservationResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImagingStudyResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedImagingStudyResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedImagingStudyResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedMolecularSequenceResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom != nil {
		for idx := range *o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom {
			rsc := (*o.IncludedDocumentReferenceResourcesReferencedByDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImmunizationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedImmunizationResourcesReferencedByPartof {
			rsc := (*o.IncludedImmunizationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationDispenseResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationDispenseResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationDispenseResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*o.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImagingStudyResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedImagingStudyResourcesReferencedByPartof {
			rsc := (*o.IncludedImagingStudyResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationStatementResourcesReferencedByPartof != nil {
		for idx := range *o.IncludedMedicationStatementResourcesReferencedByPartof {
			rsc := (*o.IncludedMedicationStatementResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedObservationResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedObservationResourcesReferencedByHasmember {
			rsc := (*o.IncludedObservationResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMolecularSequenceResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedMolecularSequenceResourcesReferencedByHasmember {
			rsc := (*o.IncludedMolecularSequenceResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedQuestionnaireResponseResourcesReferencedByHasmember != nil {
		for idx := range *o.IncludedQuestionnaireResponseResourcesReferencedByHasmember {
			rsc := (*o.IncludedQuestionnaireResponseResourcesReferencedByHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*o.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedMedicationRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedMedicationRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedMedicationRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedNutritionOrderResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedNutritionOrderResourcesReferencedByBasedon {
			rsc := (*o.IncludedNutritionOrderResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedDeviceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedDeviceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*o.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedImmunizationRecommendationResourcesReferencedByBasedon != nil {
		for idx := range *o.IncludedImmunizationRecommendationResourcesReferencedByBasedon {
			rsc := (*o.IncludedImmunizationRecommendationResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPatient {
			rsc := (*o.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedSpecimenResourcesReferencedBySpecimen != nil {
		for idx := range *o.IncludedSpecimenResourcesReferencedBySpecimen {
			rsc := (*o.IncludedSpecimenResourcesReferencedBySpecimen)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*o.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*o.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*o.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*o.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *o.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*o.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *o.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*o.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.IncludedDeviceMetricResourcesReferencedByDevice != nil {
		for idx := range *o.IncludedDeviceMetricResourcesReferencedByDevice {
			rsc := (*o.IncludedDeviceMetricResourcesReferencedByDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*o.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*o.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *o.RevIncludedConsentResourcesReferencingData {
			rsc := (*o.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*o.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *o.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*o.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*o.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*o.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedContractResourcesReferencingSubject {
			rsc := (*o.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*o.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *o.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*o.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *o.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*o.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*o.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *o.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*o.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *o.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*o.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*o.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *o.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*o.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*o.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*o.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*o.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*o.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*o.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedProcedureResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedProcedureResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedProcedureResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *o.RevIncludedListResourcesReferencingItem {
			rsc := (*o.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReaction != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReaction {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReaction)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedImmunizationResourcesReferencingReasonreference != nil {
		for idx := range *o.RevIncludedImmunizationResourcesReferencingReasonreference {
			rsc := (*o.RevIncludedImmunizationResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*o.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedObservationResourcesReferencingHasmember != nil {
		for idx := range *o.RevIncludedObservationResourcesReferencingHasmember {
			rsc := (*o.RevIncludedObservationResourcesReferencingHasmember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*o.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*o.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedMedicationStatementResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedMedicationStatementResourcesReferencingPartof {
			rsc := (*o.RevIncludedMedicationStatementResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *o.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*o.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*o.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDiagnosticReportResourcesReferencingResult != nil {
		for idx := range *o.RevIncludedDiagnosticReportResourcesReferencingResult {
			rsc := (*o.RevIncludedDiagnosticReportResourcesReferencingResult)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *o.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*o.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *o.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*o.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *o.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*o.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*o.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *o.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*o.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*o.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingPartof != nil {
		for idx := range *o.RevIncludedQuestionnaireResponseResourcesReferencingPartof {
			rsc := (*o.RevIncludedQuestionnaireResponseResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingFindingref != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingFindingref {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingFindingref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *o.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*o.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*o.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
