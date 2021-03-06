package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Location ...
type Location struct {
	DomainResource         `bson:",inline"`
	Identifier             []Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                 string                              `bson:"status,omitempty" json:"status,omitempty"`
	OperationalStatus      *Coding                             `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Name                   string                              `bson:"name,omitempty" json:"name,omitempty"`
	Alias                  []string                            `bson:"alias,omitempty" json:"alias,omitempty"`
	Description            string                              `bson:"description,omitempty" json:"description,omitempty"`
	Mode                   string                              `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                   []CodeableConcept                   `bson:"type,omitempty" json:"type,omitempty"`
	Telecom                []ContactPoint                      `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address                *Address                            `bson:"address,omitempty" json:"address,omitempty"`
	PhysicalType           *CodeableConcept                    `bson:"physicalType,omitempty" json:"physicalType,omitempty"`
	Position               *LocationPositionComponent          `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization   *Reference                          `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf                 *Reference                          `bson:"partOf,omitempty" json:"partOf,omitempty"`
	HoursOfOperation       []LocationHoursOfOperationComponent `bson:"hoursOfOperation,omitempty" json:"hoursOfOperation,omitempty"`
	AvailabilityExceptions string                              `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                         `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Location) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Location"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "location" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type location Location

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Location) UnmarshalJSON(data []byte) (err error) {
	x2 := location{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Location(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Location) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Location"
	} else if x.ResourceType != "Location" {
		return fmt.Errorf("Expected resourceType to be Location, instead received %s", x.ResourceType)
	}
	return nil
}

// LocationPositionComponent ...
type LocationPositionComponent struct {
	BackboneElement `bson:",inline"`
	Longitude       *float64 `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude        *float64 `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Altitude        *float64 `bson:"altitude,omitempty" json:"altitude,omitempty"`
}

// LocationHoursOfOperationComponent ...
type LocationHoursOfOperationComponent struct {
	BackboneElement `bson:",inline"`
	DaysOfWeek      []string      `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay          *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	OpeningTime     *FHIRDateTime `bson:"openingTime,omitempty" json:"openingTime,omitempty"`
	ClosingTime     *FHIRDateTime `bson:"closingTime,omitempty" json:"closingTime,omitempty"`
}

// LocationPlus ...
type LocationPlus struct {
	Location                     `bson:",inline"`
	LocationPlusRelatedResources `bson:",inline"`
}

// LocationPlusRelatedResources ...
type LocationPlusRelatedResources struct {
	IncludedLocationResourcesReferencedByPartof                            *[]Location                   `bson:"_includedLocationResourcesReferencedByPartof,omitempty"`
	IncludedEndpointResourcesReferencedByEndpoint                          *[]Endpoint                   `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedAppointmentResourcesReferencingLocation                     *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingLocation,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                          *[]Account                    `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
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
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingReporter                   *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingReporter,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingLocation                *[]PractitionerRole           `bson:"_revIncludedPractitionerRoleResourcesReferencingLocation,omitempty"`
	RevIncludedServiceRequestResourcesReferencingSubject                   *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSubject                    *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingSubject,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedContractResourcesReferencingDomain                          *[]Contract                   `bson:"_revIncludedContractResourcesReferencingDomain,omitempty"`
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
	RevIncludedEncounterResourcesReferencingLocation                       *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingLocation,omitempty"`
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
	RevIncludedDeviceRequestResourcesReferencingSubject                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingLocation                      *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingLocation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingFacility            *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingFacility,omitempty"`
	RevIncludedResearchStudyResourcesReferencingSite                       *[]ResearchStudy              `bson:"_revIncludedResearchStudyResourcesReferencingSite,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                         *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedProcedureResourcesReferencingLocation                       *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingLocation,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                             *[]List                       `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedImmunizationResourcesReferencingLocation                    *[]Immunization               `bson:"_revIncludedImmunizationResourcesReferencingLocation,omitempty"`
	RevIncludedDeviceResourcesReferencingLocation                          *[]Device                     `bson:"_revIncludedDeviceResourcesReferencingLocation,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                             *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingLocation             *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingLocation,omitempty"`
	RevIncludedAdverseEventResourcesReferencingLocation                    *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingLocation,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingDestination           *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingDestination,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedOrganizationAffiliationResourcesReferencingLocation         *[]OrganizationAffiliation    `bson:"_revIncludedOrganizationAffiliationResourcesReferencingLocation,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingCoveragearea           *[]HealthcareService          `bson:"_revIncludedHealthcareServiceResourcesReferencingCoveragearea,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingLocation               *[]HealthcareService          `bson:"_revIncludedHealthcareServiceResourcesReferencingLocation,omitempty"`
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
	RevIncludedCoverageEligibilityRequestResourcesReferencingFacility      *[]CoverageEligibilityRequest `bson:"_revIncludedCoverageEligibilityRequestResourcesReferencingFacility,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingFacility                           *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingFacility,omitempty"`
	RevIncludedLocationResourcesReferencingPartof                          *[]Location                   `bson:"_revIncludedLocationResourcesReferencingPartof,omitempty"`
}

// GetIncludedLocationResourceReferencedByPartof ...
func (l *LocationPlusRelatedResources) GetIncludedLocationResourceReferencedByPartof() (location *Location, err error) {
	if l.IncludedLocationResourcesReferencedByPartof == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedLocationResourcesReferencedByPartof) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedLocationResourcesReferencedByPartof))
	} else if len(*l.IncludedLocationResourcesReferencedByPartof) == 1 {
		location = &(*l.IncludedLocationResourcesReferencedByPartof)[0]
	}
	return
}

// GetIncludedEndpointResourcesReferencedByEndpoint ...
func (l *LocationPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if l.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *l.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

// GetIncludedOrganizationResourceReferencedByOrganization ...
func (l *LocationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if l.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*l.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*l.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*l.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*l.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingActor ...
func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingLocation() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingLocation
	}
	return
}

// GetRevIncludedAccountResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if l.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *l.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (l *LocationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (l *LocationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (l *LocationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (l *LocationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if l.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *l.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (l *LocationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if l.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *l.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingReporter ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingReporter() (measureReports []MeasureReport, err error) {
	if l.RevIncludedMeasureReportResourcesReferencingReporter == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *l.RevIncludedMeasureReportResourcesReferencingReporter
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (l *LocationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedPractitionerRoleResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingLocation() (practitionerRoles []PractitionerRole, err error) {
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *l.RevIncludedPractitionerRoleResourcesReferencingLocation
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingSubject() (serviceRequests []ServiceRequest, err error) {
	if l.RevIncludedServiceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *l.RevIncludedServiceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedSupplyRequestResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSubject() (supplyRequests []SupplyRequest, err error) {
	if l.RevIncludedSupplyRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *l.RevIncludedSupplyRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (l *LocationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if l.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *l.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedContractResourcesReferencingDomain ...
func (l *LocationPlusRelatedResources) GetRevIncludedContractResourcesReferencingDomain() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingDomain == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingDomain
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (l *LocationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (l *LocationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (l *LocationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if l.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *l.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedEncounterResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingLocation() (encounters []Encounter, err error) {
	if l.RevIncludedEncounterResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *l.RevIncludedEncounterResourcesReferencingLocation
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (l *LocationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (l *LocationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (l *LocationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (l *LocationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if l.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *l.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (l *LocationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if l.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *l.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if l.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *l.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (l *LocationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *l.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (l *LocationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (l *LocationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (l *LocationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingLocation() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingLocation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (l *LocationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (l *LocationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingFacility ...
func (l *LocationPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingFacility() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility
	}
	return
}

// GetRevIncludedResearchStudyResourcesReferencingSite ...
func (l *LocationPlusRelatedResources) GetRevIncludedResearchStudyResourcesReferencingSite() (researchStudies []ResearchStudy, err error) {
	if l.RevIncludedResearchStudyResourcesReferencingSite == nil {
		err = errors.New("RevIncluded researchStudies not requested")
	} else {
		researchStudies = *l.RevIncludedResearchStudyResourcesReferencingSite
	}
	return
}

// GetRevIncludedSpecimenResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if l.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *l.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingLocation() (procedures []Procedure, err error) {
	if l.RevIncludedProcedureResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *l.RevIncludedProcedureResourcesReferencingLocation
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (l *LocationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedListResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingSubject
	}
	return
}

// GetRevIncludedImmunizationResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedImmunizationResourcesReferencingLocation() (immunizations []Immunization, err error) {
	if l.RevIncludedImmunizationResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded immunizations not requested")
	} else {
		immunizations = *l.RevIncludedImmunizationResourcesReferencingLocation
	}
	return
}

// GetRevIncludedDeviceResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedDeviceResourcesReferencingLocation() (devices []Device, err error) {
	if l.RevIncludedDeviceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded devices not requested")
	} else {
		devices = *l.RevIncludedDeviceResourcesReferencingLocation
	}
	return
}

// GetRevIncludedMediaResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if l.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *l.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedFlagResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if l.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *l.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingActor ...
func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if l.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *l.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingLocation() (appointmentResponses []AppointmentResponse, err error) {
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *l.RevIncludedAppointmentResponseResourcesReferencingLocation
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingLocation() (adverseEvents []AdverseEvent, err error) {
	if l.RevIncludedAdverseEventResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *l.RevIncludedAdverseEventResourcesReferencingLocation
	}
	return
}

// GetRevIncludedObservationResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if l.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *l.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (l *LocationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if l.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *l.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (l *LocationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (l *LocationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *l.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingDestination ...
func (l *LocationPlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingDestination() (medicationDispenses []MedicationDispense, err error) {
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *l.RevIncludedMedicationDispenseResourcesReferencingDestination
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *l.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedOrganizationAffiliationResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingLocation() (organizationAffiliations []OrganizationAffiliation, err error) {
	if l.RevIncludedOrganizationAffiliationResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *l.RevIncludedOrganizationAffiliationResourcesReferencingLocation
	}
	return
}

// GetRevIncludedHealthcareServiceResourcesReferencingCoveragearea ...
func (l *LocationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingCoveragearea() (healthcareServices []HealthcareService, err error) {
	if l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea
	}
	return
}

// GetRevIncludedHealthcareServiceResourcesReferencingLocation ...
func (l *LocationPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingLocation() (healthcareServices []HealthcareService, err error) {
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *l.RevIncludedHealthcareServiceResourcesReferencingLocation
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (l *LocationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (l *LocationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (l *LocationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if l.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *l.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (l *LocationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (l *LocationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (l *LocationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCoverageEligibilityRequestResourcesReferencingFacility ...
func (l *LocationPlusRelatedResources) GetRevIncludedCoverageEligibilityRequestResourcesReferencingFacility() (coverageEligibilityRequests []CoverageEligibilityRequest, err error) {
	if l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded coverageEligibilityRequests not requested")
	} else {
		coverageEligibilityRequests = *l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility
	}
	return
}

// GetRevIncludedScheduleResourcesReferencingActor ...
func (l *LocationPlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if l.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *l.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (l *LocationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (l *LocationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedClaimResourcesReferencingFacility ...
func (l *LocationPlusRelatedResources) GetRevIncludedClaimResourcesReferencingFacility() (claims []Claim, err error) {
	if l.RevIncludedClaimResourcesReferencingFacility == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *l.RevIncludedClaimResourcesReferencingFacility
	}
	return
}

// GetRevIncludedLocationResourcesReferencingPartof ...
func (l *LocationPlusRelatedResources) GetRevIncludedLocationResourcesReferencingPartof() (locations []Location, err error) {
	if l.RevIncludedLocationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *l.RevIncludedLocationResourcesReferencingPartof
	}
	return
}

// GetIncludedResources ...
func (l *LocationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedLocationResourcesReferencedByPartof != nil {
		for idx := range *l.IncludedLocationResourcesReferencedByPartof {
			rsc := (*l.IncludedLocationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *l.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*l.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*l.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (l *LocationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*l.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*l.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerRoleResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerRoleResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*l.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingDomain != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingDomain {
			rsc := (*l.RevIncludedContractResourcesReferencingDomain)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*l.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*l.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*l.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility {
			rsc := (*l.RevIncludedExplanationOfBenefitResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchStudyResourcesReferencingSite != nil {
		for idx := range *l.RevIncludedResearchStudyResourcesReferencingSite {
			rsc := (*l.RevIncludedResearchStudyResourcesReferencingSite)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*l.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProcedureResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProcedureResourcesReferencingLocation {
			rsc := (*l.RevIncludedProcedureResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedListResourcesReferencingSubject {
			rsc := (*l.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImmunizationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedImmunizationResourcesReferencingLocation {
			rsc := (*l.RevIncludedImmunizationResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedDeviceResourcesReferencingLocation {
			rsc := (*l.RevIncludedDeviceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*l.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAdverseEventResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAdverseEventResourcesReferencingLocation {
			rsc := (*l.RevIncludedAdverseEventResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*l.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*l.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*l.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination != nil {
		for idx := range *l.RevIncludedMedicationDispenseResourcesReferencingDestination {
			rsc := (*l.RevIncludedMedicationDispenseResourcesReferencingDestination)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedOrganizationAffiliationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedOrganizationAffiliationResourcesReferencingLocation {
			rsc := (*l.RevIncludedOrganizationAffiliationResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingLocation {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *l.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*l.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility {
			rsc := (*l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*l.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedClaimResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedClaimResourcesReferencingFacility {
			rsc := (*l.RevIncludedClaimResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLocationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedLocationResourcesReferencingPartof {
			rsc := (*l.RevIncludedLocationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (l *LocationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedLocationResourcesReferencedByPartof != nil {
		for idx := range *l.IncludedLocationResourcesReferencedByPartof {
			rsc := (*l.IncludedLocationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *l.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*l.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *l.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*l.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*l.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*l.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingReporter != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingReporter {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingReporter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPractitionerRoleResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedPractitionerRoleResourcesReferencingLocation {
			rsc := (*l.RevIncludedPractitionerRoleResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedSupplyRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedSupplyRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedSupplyRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*l.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingDomain != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingDomain {
			rsc := (*l.RevIncludedContractResourcesReferencingDomain)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEncounterResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedEncounterResourcesReferencingLocation {
			rsc := (*l.RevIncludedEncounterResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*l.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*l.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*l.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingLocation {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedExplanationOfBenefitResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedExplanationOfBenefitResourcesReferencingFacility {
			rsc := (*l.RevIncludedExplanationOfBenefitResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedResearchStudyResourcesReferencingSite != nil {
		for idx := range *l.RevIncludedResearchStudyResourcesReferencingSite {
			rsc := (*l.RevIncludedResearchStudyResourcesReferencingSite)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*l.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedProcedureResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedProcedureResourcesReferencingLocation {
			rsc := (*l.RevIncludedProcedureResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedListResourcesReferencingSubject {
			rsc := (*l.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedImmunizationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedImmunizationResourcesReferencingLocation {
			rsc := (*l.RevIncludedImmunizationResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDeviceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedDeviceResourcesReferencingLocation {
			rsc := (*l.RevIncludedDeviceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*l.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*l.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAppointmentResponseResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAppointmentResponseResourcesReferencingLocation {
			rsc := (*l.RevIncludedAppointmentResponseResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAdverseEventResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedAdverseEventResourcesReferencingLocation {
			rsc := (*l.RevIncludedAdverseEventResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*l.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*l.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*l.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedMedicationDispenseResourcesReferencingDestination != nil {
		for idx := range *l.RevIncludedMedicationDispenseResourcesReferencingDestination {
			rsc := (*l.RevIncludedMedicationDispenseResourcesReferencingDestination)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*l.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedOrganizationAffiliationResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedOrganizationAffiliationResourcesReferencingLocation {
			rsc := (*l.RevIncludedOrganizationAffiliationResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingCoveragearea)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedHealthcareServiceResourcesReferencingLocation != nil {
		for idx := range *l.RevIncludedHealthcareServiceResourcesReferencingLocation {
			rsc := (*l.RevIncludedHealthcareServiceResourcesReferencingLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *l.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*l.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility {
			rsc := (*l.RevIncludedCoverageEligibilityRequestResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *l.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*l.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedClaimResourcesReferencingFacility != nil {
		for idx := range *l.RevIncludedClaimResourcesReferencingFacility {
			rsc := (*l.RevIncludedClaimResourcesReferencingFacility)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if l.RevIncludedLocationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedLocationResourcesReferencingPartof {
			rsc := (*l.RevIncludedLocationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
