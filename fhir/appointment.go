package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Appointment is a FHIR Appointment resource
type Appointment struct {
	DomainResource        `bson:",inline"`

	Identifier            []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                            `bson:"status,omitempty" json:"status,omitempty"`
	CancelationReason     *CodeableConcept                  `bson:"cancelationReason,omitempty" json:"cancelationReason,omitempty"`
	ServiceCategory       []CodeableConcept                 `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept                 `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty             []CodeableConcept                 `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept                  `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	ReasonCode            []CodeableConcept                 `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference                       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Priority              *uint32                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Description           string                            `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInformation []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Start                 *FHIRDateTime                     `bson:"start,omitempty" json:"start,omitempty"`
	End                   *FHIRDateTime                     `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration       *uint32                           `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot                  []Reference                       `bson:"slot,omitempty" json:"slot,omitempty"`
	Created               *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Comment               string                            `bson:"comment,omitempty" json:"comment,omitempty"`
	PatientInstruction    string                            `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	BasedOn               []Reference                       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Participant           []AppointmentParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	RequestedPeriod       []Period                          `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Appointment) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Appointment"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "appointment" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type appointment Appointment

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Appointment) UnmarshalJSON(data []byte) (err error) {
	x2 := appointment{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Appointment(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Appointment) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Appointment"
	} else if x.ResourceType != "Appointment" {
		return fmt.Errorf("Expected resourceType to be Appointment, instead received %s", x.ResourceType)
	}
	return nil
}

// AppointmentParticipantComponent ...
type AppointmentParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Required        string            `bson:"required,omitempty" json:"required,omitempty"`
	Status          string            `bson:"status,omitempty" json:"status,omitempty"`
	Period          *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

// AppointmentPlus ...
type AppointmentPlus struct {
	Appointment                     `bson:",inline"`
	AppointmentPlusRelatedResources `bson:",inline"`
}

// AppointmentPlusRelatedResources ...
type AppointmentPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPractitioner                  *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPractitioner,omitempty"`
	IncludedSlotResourcesReferencedBySlot                                  *[]Slot                       `bson:"_includedSlotResourcesReferencedBySlot,omitempty"`
	IncludedPractitionerResourcesReferencedByActor                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByActor,omitempty"`
	IncludedDeviceResourcesReferencedByActor                               *[]Device                     `bson:"_includedDeviceResourcesReferencedByActor,omitempty"`
	IncludedPatientResourcesReferencedByActor                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByActor,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByActor                    *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByActor,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByActor                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByActor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByActor                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByActor,omitempty"`
	IncludedLocationResourcesReferencedByActor                             *[]Location                   `bson:"_includedLocationResourcesReferencedByActor,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedConditionResourcesReferencedByReasonreference                  *[]Condition                  `bson:"_includedConditionResourcesReferencedByReasonreference,omitempty"`
	IncludedObservationResourcesReferencedByReasonreference                *[]Observation                `bson:"_includedObservationResourcesReferencedByReasonreference,omitempty"`
	IncludedProcedureResourcesReferencedByReasonreference                  *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByReasonreference,omitempty"`
	IncludedImmunizationRecommendationResourcesReferencedByReasonreference *[]ImmunizationRecommendation `bson:"_includedImmunizationRecommendationResourcesReferencedByReasonreference,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
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
	RevIncludedImagingStudyResourcesReferencingBasedon                     *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingAppointment                    *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingAppointment,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingActivityreference               *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingAppointment          *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingAppointment,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
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
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

// GetIncludedPractitionerResourceReferencedByPractitioner ...
func (a *AppointmentPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPractitioner() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByPractitioner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByPractitioner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByPractitioner))
	} else if len(*a.IncludedPractitionerResourcesReferencedByPractitioner) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByPractitioner)[0]
	}
	return
}

// GetIncludedSlotResourcesReferencedBySlot ...
func (a *AppointmentPlusRelatedResources) GetIncludedSlotResourcesReferencedBySlot() (slots []Slot, err error) {
	if a.IncludedSlotResourcesReferencedBySlot == nil {
		err = errors.New("Included slots not requested")
	} else {
		slots = *a.IncludedSlotResourcesReferencedBySlot
	}
	return
}

// GetIncludedPractitionerResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActor() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByActor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByActor))
	} else if len(*a.IncludedPractitionerResourcesReferencedByActor) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedDeviceResourceReferencedByActor() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedByActor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedByActor))
	} else if len(*a.IncludedDeviceResourcesReferencedByActor) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedPatientResourceReferencedByActor() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByActor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByActor))
	} else if len(*a.IncludedPatientResourcesReferencedByActor) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedHealthcareServiceResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedHealthcareServiceResourceReferencedByActor() (healthcareService *HealthcareService, err error) {
	if a.IncludedHealthcareServiceResourcesReferencedByActor == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*a.IncludedHealthcareServiceResourcesReferencedByActor))
	} else if len(*a.IncludedHealthcareServiceResourcesReferencedByActor) == 1 {
		healthcareService = &(*a.IncludedHealthcareServiceResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByActor() (practitionerRole *PractitionerRole, err error) {
	if a.IncludedPractitionerRoleResourcesReferencedByActor == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*a.IncludedPractitionerRoleResourcesReferencedByActor))
	} else if len(*a.IncludedPractitionerRoleResourcesReferencedByActor) == 1 {
		practitionerRole = &(*a.IncludedPractitionerRoleResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedRelatedPersonResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByActor() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByActor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByActor))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByActor) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedLocationResourceReferencedByActor ...
func (a *AppointmentPlusRelatedResources) GetIncludedLocationResourceReferencedByActor() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedByActor == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedByActor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedByActor))
	} else if len(*a.IncludedLocationResourcesReferencedByActor) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedByActor)[0]
	}
	return
}

// GetIncludedServiceRequestResourcesReferencedByBasedon ...
func (a *AppointmentPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if a.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *a.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ...
func (a *AppointmentPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatient))
	} else if len(*a.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedConditionResourcesReferencedByReasonreference ...
func (a *AppointmentPlusRelatedResources) GetIncludedConditionResourcesReferencedByReasonreference() (conditions []Condition, err error) {
	if a.IncludedConditionResourcesReferencedByReasonreference == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *a.IncludedConditionResourcesReferencedByReasonreference
	}
	return
}

// GetIncludedObservationResourcesReferencedByReasonreference ...
func (a *AppointmentPlusRelatedResources) GetIncludedObservationResourcesReferencedByReasonreference() (observations []Observation, err error) {
	if a.IncludedObservationResourcesReferencedByReasonreference == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *a.IncludedObservationResourcesReferencedByReasonreference
	}
	return
}

// GetIncludedProcedureResourcesReferencedByReasonreference ...
func (a *AppointmentPlusRelatedResources) GetIncludedProcedureResourcesReferencedByReasonreference() (procedures []Procedure, err error) {
	if a.IncludedProcedureResourcesReferencedByReasonreference == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *a.IncludedProcedureResourcesReferencedByReasonreference
	}
	return
}

// GetIncludedImmunizationRecommendationResourcesReferencedByReasonreference ...
func (a *AppointmentPlusRelatedResources) GetIncludedImmunizationRecommendationResourcesReferencedByReasonreference() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference == nil {
		err = errors.New("Included immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference
	}
	return
}

// GetIncludedLocationResourceReferencedByLocation ...
func (a *AppointmentPlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if a.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*a.IncludedLocationResourcesReferencedByLocation))
	} else if len(*a.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*a.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *a.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if a.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *a.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if a.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *a.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *a.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if a.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *a.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if a.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *a.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if a.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *a.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingBasedon ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if a.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *a.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingAppointment ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingAppointment() (encounters []Encounter, err error) {
	if a.RevIncludedEncounterResourcesReferencingAppointment == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *a.RevIncludedEncounterResourcesReferencingAppointment
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if a.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *a.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if a.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *a.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if a.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *a.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingActivityreference ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if a.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *a.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *a.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingAppointment ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingAppointment() (appointmentResponses []AppointmentResponse, err error) {
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *a.RevIncludedAppointmentResponseResourcesReferencingAppointment
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if a.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *a.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if a.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *a.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *a.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if a.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *a.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if a.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *a.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (a *AppointmentPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*a.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedSlotResourcesReferencedBySlot != nil {
		for idx := range *a.IncludedSlotResourcesReferencedBySlot {
			rsc := (*a.IncludedSlotResourcesReferencedBySlot)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByActor {
			rsc := (*a.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByActor {
			rsc := (*a.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByActor {
			rsc := (*a.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *a.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*a.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*a.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*a.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*a.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference {
			rsc := (*a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (a *AppointmentPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*a.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEncounterResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedEncounterResourcesReferencingAppointment {
			rsc := (*a.RevIncludedEncounterResourcesReferencingAppointment)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *a.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*a.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedAppointmentResponseResourcesReferencingAppointment {
			rsc := (*a.RevIncludedAppointmentResponseResourcesReferencingAppointment)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (a *AppointmentPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByPractitioner != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByPractitioner {
			rsc := (*a.IncludedPractitionerResourcesReferencedByPractitioner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedSlotResourcesReferencedBySlot != nil {
		for idx := range *a.IncludedSlotResourcesReferencedBySlot {
			rsc := (*a.IncludedSlotResourcesReferencedBySlot)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPractitionerResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedDeviceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedDeviceResourcesReferencedByActor {
			rsc := (*a.IncludedDeviceResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByActor {
			rsc := (*a.IncludedPatientResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedHealthcareServiceResourcesReferencedByActor != nil {
		for idx := range *a.IncludedHealthcareServiceResourcesReferencedByActor {
			rsc := (*a.IncludedHealthcareServiceResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPractitionerRoleResourcesReferencedByActor != nil {
		for idx := range *a.IncludedPractitionerRoleResourcesReferencedByActor {
			rsc := (*a.IncludedPractitionerRoleResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByActor != nil {
		for idx := range *a.IncludedRelatedPersonResourcesReferencedByActor {
			rsc := (*a.IncludedRelatedPersonResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByActor != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByActor {
			rsc := (*a.IncludedLocationResourcesReferencedByActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *a.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*a.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *a.IncludedPatientResourcesReferencedByPatient {
			rsc := (*a.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*a.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*a.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*a.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference != nil {
		for idx := range *a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference {
			rsc := (*a.IncludedImmunizationRecommendationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *a.IncludedLocationResourcesReferencedByLocation {
			rsc := (*a.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*a.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*a.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *a.RevIncludedConsentResourcesReferencingData {
			rsc := (*a.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*a.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *a.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*a.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*a.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*a.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedContractResourcesReferencingSubject {
			rsc := (*a.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*a.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *a.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*a.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*a.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEncounterResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedEncounterResourcesReferencingAppointment {
			rsc := (*a.RevIncludedEncounterResourcesReferencingAppointment)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*a.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *a.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*a.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *a.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*a.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*a.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *a.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*a.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*a.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*a.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*a.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*a.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *a.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*a.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *a.RevIncludedListResourcesReferencingItem {
			rsc := (*a.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedAppointmentResponseResourcesReferencingAppointment != nil {
		for idx := range *a.RevIncludedAppointmentResponseResourcesReferencingAppointment {
			rsc := (*a.RevIncludedAppointmentResponseResourcesReferencingAppointment)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *a.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*a.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*a.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*a.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *a.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*a.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*a.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *a.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*a.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *a.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*a.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *a.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*a.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*a.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *a.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*a.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*a.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*a.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*a.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*a.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
