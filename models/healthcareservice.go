package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type HealthcareService struct {
	DomainResource         `bson:",inline"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                                     `bson:"active,omitempty" json:"active,omitempty"`
	ProvidedBy             *Reference                                `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	Category               []CodeableConcept                         `bson:"category,omitempty" json:"category,omitempty"`
	Type                   []CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Specialty              []CodeableConcept                         `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                               `bson:"location,omitempty" json:"location,omitempty"`
	Name                   string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Comment                string                                    `bson:"comment,omitempty" json:"comment,omitempty"`
	ExtraDetails           string                                    `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	Photo                  *Attachment                               `bson:"photo,omitempty" json:"photo,omitempty"`
	Telecom                []ContactPoint                            `bson:"telecom,omitempty" json:"telecom,omitempty"`
	CoverageArea           []Reference                               `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                         `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	Eligibility            []HealthcareServiceEligibilityComponent   `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	Program                []CodeableConcept                         `bson:"program,omitempty" json:"program,omitempty"`
	Characteristic         []CodeableConcept                         `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Communication          []CodeableConcept                         `bson:"communication,omitempty" json:"communication,omitempty"`
	ReferralMethod         []CodeableConcept                         `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	AppointmentRequired    *bool                                     `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTimeComponent `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailableComponent  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions string                                    `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                               `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *HealthcareService) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "HealthcareService"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "healthcareService" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type healthcareService HealthcareService

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *HealthcareService) UnmarshalJSON(data []byte) (err error) {
	x2 := healthcareService{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = HealthcareService(x2)
		return x.checkResourceType()
	}
	return
}

func (x *HealthcareService) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "HealthcareService"
	} else if x.ResourceType != "HealthcareService" {
		return errors.New(fmt.Sprintf("Expected resourceType to be HealthcareService, instead received %s", x.ResourceType))
	}
	return nil
}

type HealthcareServiceEligibilityComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Comment         string           `bson:"comment,omitempty" json:"comment,omitempty"`
}

type HealthcareServiceAvailableTimeComponent struct {
	BackboneElement    `bson:",inline"`
	DaysOfWeek         []string      `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool         `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *FHIRDateTime `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *FHIRDateTime `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}

type HealthcareServiceNotAvailableComponent struct {
	BackboneElement `bson:",inline"`
	Description     string  `bson:"description,omitempty" json:"description,omitempty"`
	During          *Period `bson:"during,omitempty" json:"during,omitempty"`
}

type HealthcareServicePlus struct {
	HealthcareService                     `bson:",inline"`
	HealthcareServicePlusRelatedResources `bson:",inline"`
}

type HealthcareServicePlusRelatedResources struct {
	IncludedEndpointResourcesReferencedByEndpoint                          *[]Endpoint                   `bson:"_includedEndpointResourcesReferencedByEndpoint,omitempty"`
	IncludedLocationResourcesReferencedByCoveragearea                      *[]Location                   `bson:"_includedLocationResourcesReferencedByCoveragearea,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
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
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedPractitionerRoleResourcesReferencingService                 *[]PractitionerRole           `bson:"_revIncludedPractitionerRoleResourcesReferencingService,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingSupplier                   *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingSupplier,omitempty"`
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
	RevIncludedDeviceRequestResourcesReferencingPerformer                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                       *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender              *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedOrganizationAffiliationResourcesReferencingService          *[]OrganizationAffiliation    `bson:"_revIncludedOrganizationAffiliationResourcesReferencingService,omitempty"`
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
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpoint() (endpoints []Endpoint, err error) {
	if h.IncludedEndpointResourcesReferencedByEndpoint == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *h.IncludedEndpointResourcesReferencedByEndpoint
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedLocationResourcesReferencedByCoveragearea() (locations []Location, err error) {
	if h.IncludedLocationResourcesReferencedByCoveragearea == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *h.IncludedLocationResourcesReferencedByCoveragearea
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if h.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*h.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*h.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*h.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*h.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedLocationResourcesReferencedByLocation() (locations []Location, err error) {
	if h.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else {
		locations = *h.IncludedLocationResourcesReferencedByLocation
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if h.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *h.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if h.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *h.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if h.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *h.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if h.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *h.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if h.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *h.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if h.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *h.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if h.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *h.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if h.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *h.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *h.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if h.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *h.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *h.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if h.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *h.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingService() (practitionerRoles []PractitionerRole, err error) {
	if h.RevIncludedPractitionerRoleResourcesReferencingService == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *h.RevIncludedPractitionerRoleResourcesReferencingService
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if h.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *h.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingSupplier() (supplyRequests []SupplyRequest, err error) {
	if h.RevIncludedSupplyRequestResourcesReferencingSupplier == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *h.RevIncludedSupplyRequestResourcesReferencingSupplier
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if h.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *h.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if h.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *h.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *h.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *h.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if h.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *h.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if h.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *h.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if h.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *h.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if h.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *h.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if h.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *h.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if h.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *h.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if h.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *h.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if h.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *h.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if h.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *h.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if h.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *h.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if h.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *h.RevIncludedListResourcesReferencingItem
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if h.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *h.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if h.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *h.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if h.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *h.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if h.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *h.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if h.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *h.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if h.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *h.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if h.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *h.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *h.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if h.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *h.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if h.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *h.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if h.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *h.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingService() (organizationAffiliations []OrganizationAffiliation, err error) {
	if h.RevIncludedOrganizationAffiliationResourcesReferencingService == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *h.RevIncludedOrganizationAffiliationResourcesReferencingService
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if h.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *h.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if h.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *h.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if h.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *h.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if h.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *h.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if h.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *h.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if h.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *h.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if h.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *h.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if h.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *h.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if h.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *h.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *h.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if h.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *h.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *h.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*h.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByCoveragearea != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByCoveragearea {
			rsc := (*h.IncludedLocationResourcesReferencedByCoveragearea)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *h.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*h.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByLocation {
			rsc := (*h.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (h *HealthcareServicePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingData {
			rsc := (*h.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*h.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPractitionerRoleResourcesReferencingService != nil {
		for idx := range *h.RevIncludedPractitionerRoleResourcesReferencingService {
			rsc := (*h.RevIncludedPractitionerRoleResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *h.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*h.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*h.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingSubject {
			rsc := (*h.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *h.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*h.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*h.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *h.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*h.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *h.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*h.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*h.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*h.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*h.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*h.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*h.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*h.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*h.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrganizationAffiliationResourcesReferencingService != nil {
		for idx := range *h.RevIncludedOrganizationAffiliationResourcesReferencingService {
			rsc := (*h.RevIncludedOrganizationAffiliationResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *h.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*h.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*h.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*h.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *h.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*h.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*h.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*h.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (h *HealthcareServicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if h.IncludedEndpointResourcesReferencedByEndpoint != nil {
		for idx := range *h.IncludedEndpointResourcesReferencedByEndpoint {
			rsc := (*h.IncludedEndpointResourcesReferencedByEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByCoveragearea != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByCoveragearea {
			rsc := (*h.IncludedLocationResourcesReferencedByCoveragearea)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *h.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*h.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *h.IncludedLocationResourcesReferencedByLocation {
			rsc := (*h.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *h.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*h.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*h.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*h.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *h.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*h.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *h.RevIncludedConsentResourcesReferencingData {
			rsc := (*h.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*h.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *h.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*h.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*h.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPractitionerRoleResourcesReferencingService != nil {
		for idx := range *h.RevIncludedPractitionerRoleResourcesReferencingService {
			rsc := (*h.RevIncludedPractitionerRoleResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedSupplyRequestResourcesReferencingSupplier != nil {
		for idx := range *h.RevIncludedSupplyRequestResourcesReferencingSupplier {
			rsc := (*h.RevIncludedSupplyRequestResourcesReferencingSupplier)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*h.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedContractResourcesReferencingSubject {
			rsc := (*h.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *h.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*h.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *h.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*h.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*h.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*h.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *h.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*h.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *h.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*h.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*h.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *h.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*h.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *h.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*h.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*h.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*h.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*h.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*h.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *h.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*h.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *h.RevIncludedListResourcesReferencingItem {
			rsc := (*h.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*h.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*h.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *h.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*h.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*h.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*h.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *h.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*h.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*h.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedOrganizationAffiliationResourcesReferencingService != nil {
		for idx := range *h.RevIncludedOrganizationAffiliationResourcesReferencingService {
			rsc := (*h.RevIncludedOrganizationAffiliationResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *h.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*h.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *h.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*h.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *h.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*h.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*h.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *h.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*h.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *h.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*h.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *h.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*h.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *h.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*h.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*h.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*h.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
