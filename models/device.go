package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Device ...
type Device struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition         *Reference                      `bson:"definition,omitempty" json:"definition,omitempty"`
	UdiCarrier         []DeviceUdiCarrierComponent     `bson:"udiCarrier,omitempty" json:"udiCarrier,omitempty"`
	Status             string                          `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason       []CodeableConcept               `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	DistinctIdentifier string                          `bson:"distinctIdentifier,omitempty" json:"distinctIdentifier,omitempty"`
	Manufacturer       string                          `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	ManufactureDate    *FHIRDateTime                   `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	ExpirationDate     *FHIRDateTime                   `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	LotNumber          string                          `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	SerialNumber       string                          `bson:"serialNumber,omitempty" json:"serialNumber,omitempty"`
	DeviceName         []DeviceDeviceNameComponent     `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	ModelNumber        string                          `bson:"modelNumber,omitempty" json:"modelNumber,omitempty"`
	PartNumber         string                          `bson:"partNumber,omitempty" json:"partNumber,omitempty"`
	Type               *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Specialization     []DeviceSpecializationComponent `bson:"specialization,omitempty" json:"specialization,omitempty"`
	Version            []DeviceVersionComponent        `bson:"version,omitempty" json:"version,omitempty"`
	Property           []DevicePropertyComponent       `bson:"property,omitempty" json:"property,omitempty"`
	Patient            *Reference                      `bson:"patient,omitempty" json:"patient,omitempty"`
	Owner              *Reference                      `bson:"owner,omitempty" json:"owner,omitempty"`
	Contact            []ContactPoint                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Location           *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	URL                string                          `bson:"url,omitempty" json:"url,omitempty"`
	Note               []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	Safety             []CodeableConcept               `bson:"safety,omitempty" json:"safety,omitempty"`
	Parent             *Reference                      `bson:"parent,omitempty" json:"parent,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Device) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Device"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "device" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type device Device

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Device) UnmarshalJSON(data []byte) (err error) {
	x2 := device{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Device(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Device) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Device"
	} else if x.ResourceType != "Device" {
		return fmt.Errorf("Expected resourceType to be Device, instead received %s", x.ResourceType)
	}
	return nil
}

// DeviceUdiCarrierComponent ...
type DeviceUdiCarrierComponent struct {
	BackboneElement  `bson:",inline"`
	DeviceIdentifier string `bson:"deviceIdentifier,omitempty" json:"deviceIdentifier,omitempty"`
	Issuer           string `bson:"issuer,omitempty" json:"issuer,omitempty"`
	Jurisdiction     string `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	CarrierAIDC      string `bson:"carrierAIDC,omitempty" json:"carrierAIDC,omitempty"`
	CarrierHRF       string `bson:"carrierHRF,omitempty" json:"carrierHRF,omitempty"`
	EntryType        string `bson:"entryType,omitempty" json:"entryType,omitempty"`
}

// DeviceDeviceNameComponent ...
type DeviceDeviceNameComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
}

// DeviceSpecializationComponent ...
type DeviceSpecializationComponent struct {
	BackboneElement `bson:",inline"`
	SystemType      *CodeableConcept `bson:"systemType,omitempty" json:"systemType,omitempty"`
	Version         string           `bson:"version,omitempty" json:"version,omitempty"`
}

// DeviceVersionComponent ...
type DeviceVersionComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Component       *Identifier      `bson:"component,omitempty" json:"component,omitempty"`
	Value           string           `bson:"value,omitempty" json:"value,omitempty"`
}

// DevicePropertyComponent ...
type DevicePropertyComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	ValueQuantity   []Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCode       []CodeableConcept `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}

// DevicePlus ...
type DevicePlus struct {
	Device                     `bson:",inline"`
	DevicePlusRelatedResources `bson:",inline"`
}

// DevicePlusRelatedResources ...
type DevicePlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
	RevIncludedAppointmentResourcesReferencingActor                        *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingActor,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedAccountResourcesReferencingSubject                          *[]Account                    `bson:"_revIncludedAccountResourcesReferencingSubject,omitempty"`
	RevIncludedInvoiceResourcesReferencingParticipant                      *[]Invoice                    `bson:"_revIncludedInvoiceResourcesReferencingParticipant,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingSubject                 *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingSubject,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingAuthor                  *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingAuthor,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
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
	RevIncludedDocumentReferenceResourcesReferencingAuthor                 *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingAuthor,omitempty"`
	RevIncludedMeasureReportResourcesReferencingSubject                    *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingSubject,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingRequester                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedServiceRequestResourcesReferencingPerformer                 *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedServiceRequestResourcesReferencingSubject                   *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedSupplyRequestResourcesReferencingRequester                  *[]SupplyRequest              `bson:"_revIncludedSupplyRequestResourcesReferencingRequester,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedRiskAssessmentResourcesReferencingPerformer                 *[]RiskAssessment             `bson:"_revIncludedRiskAssessmentResourcesReferencingPerformer,omitempty"`
	RevIncludedGroupResourcesReferencingMember                             *[]Group                      `bson:"_revIncludedGroupResourcesReferencingMember,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingPerformer                   *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingPerformer,omitempty"`
	RevIncludedImagingStudyResourcesReferencingSubject                     *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingSubject,omitempty"`
	RevIncludedChargeItemResourcesReferencingEnterer                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingEnterer,omitempty"`
	RevIncludedChargeItemResourcesReferencingPerformeractor                *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingPerformeractor,omitempty"`
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
	RevIncludedDeviceUseStatementResourcesReferencingDevice                *[]DeviceUseStatement         `bson:"_revIncludedDeviceUseStatementResourcesReferencingDevice,omitempty"`
	RevIncludedRequestGroupResourcesReferencingAuthor                      *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingAuthor,omitempty"`
	RevIncludedRequestGroupResourcesReferencingParticipant                 *[]RequestGroup               `bson:"_revIncludedRequestGroupResourcesReferencingParticipant,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingRequester                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingRequester,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPerformer                  *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPerformer,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingSubject                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingSubject,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingDevice                     *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingDevice,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingTarget                     *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingTarget,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingAgent                         *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingAgent,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingOwner                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingOwner,omitempty"`
	RevIncludedTaskResourcesReferencingRequester                           *[]Task                       `bson:"_revIncludedTaskResourcesReferencingRequester,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingDetailudi           *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingDetailudi,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi        *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingProcedureudi,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi        *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingSubdetailudi,omitempty"`
	RevIncludedExplanationOfBenefitResourcesReferencingItemudi             *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingItemudi,omitempty"`
	RevIncludedSpecimenResourcesReferencingSubject                         *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingSubject,omitempty"`
	RevIncludedCarePlanResourcesReferencingPerformer                       *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPerformer,omitempty"`
	RevIncludedProcedureResourcesReferencingPerformer                      *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPerformer,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedListResourcesReferencingSubject                             *[]List                       `bson:"_revIncludedListResourcesReferencingSubject,omitempty"`
	RevIncludedListResourcesReferencingSource                              *[]List                       `bson:"_revIncludedListResourcesReferencingSource,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingRequester              *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedMedicationRequestResourcesReferencingIntendedperformer      *[]MedicationRequest          `bson:"_revIncludedMedicationRequestResourcesReferencingIntendedperformer,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedMediaResourcesReferencingOperator                           *[]Media                      `bson:"_revIncludedMediaResourcesReferencingOperator,omitempty"`
	RevIncludedMediaResourcesReferencingDevice                             *[]Media                      `bson:"_revIncludedMediaResourcesReferencingDevice,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedDeviceMetricResourcesReferencingParent                      *[]DeviceMetric               `bson:"_revIncludedDeviceMetricResourcesReferencingParent,omitempty"`
	RevIncludedDeviceMetricResourcesReferencingSource                      *[]DeviceMetric               `bson:"_revIncludedDeviceMetricResourcesReferencingSource,omitempty"`
	RevIncludedFlagResourcesReferencingAuthor                              *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingAuthor,omitempty"`
	RevIncludedAppointmentResponseResourcesReferencingActor                *[]AppointmentResponse        `bson:"_revIncludedAppointmentResponseResourcesReferencingActor,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubstance                   *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubstance,omitempty"`
	RevIncludedObservationResourcesReferencingSubject                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSubject,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingDevice                       *[]Observation                `bson:"_revIncludedObservationResourcesReferencingDevice,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingPerformer       *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingPerformer,omitempty"`
	RevIncludedMedicationAdministrationResourcesReferencingDevice          *[]MedicationAdministration   `bson:"_revIncludedMedicationAdministrationResourcesReferencingDevice,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRequester           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRequester,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingSender              *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingSender,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingRecipient           *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingRecipient,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedMedicationDispenseResourcesReferencingPerformer             *[]MedicationDispense         `bson:"_revIncludedMedicationDispenseResourcesReferencingPerformer,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSubject                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingAgent                         *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingAgent,omitempty"`
	RevIncludedAuditEventResourcesReferencingSource                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingSource,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingAuthor                       *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingAuthor,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingAuthor                     *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingAuthor,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingAuthor             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingAuthor,omitempty"`
	RevIncludedScheduleResourcesReferencingActor                           *[]Schedule                   `bson:"_revIncludedScheduleResourcesReferencingActor,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedClaimResourcesReferencingDetailudi                          *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingDetailudi,omitempty"`
	RevIncludedClaimResourcesReferencingProcedureudi                       *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingProcedureudi,omitempty"`
	RevIncludedClaimResourcesReferencingSubdetailudi                       *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingSubdetailudi,omitempty"`
	RevIncludedClaimResourcesReferencingItemudi                            *[]Claim                      `bson:"_revIncludedClaimResourcesReferencingItemudi,omitempty"`
}

// GetIncludedPatientResourceReferencedByPatient ...
func (d *DevicePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedOrganizationResourceReferencedByOrganization ...
func (d *DevicePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*d.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*d.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*d.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*d.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

// GetIncludedLocationResourceReferencedByLocation ...
func (d *DevicePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedByLocation))
	} else if len(*d.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingActor ...
func (d *DevicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingActor() (appointments []Appointment, err error) {
	if d.RevIncludedAppointmentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *d.RevIncludedAppointmentResourcesReferencingActor
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (d *DevicePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *d.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedAccountResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedAccountResourcesReferencingSubject() (accounts []Account, err error) {
	if d.RevIncludedAccountResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded accounts not requested")
	} else {
		accounts = *d.RevIncludedAccountResourcesReferencingSubject
	}
	return
}

// GetRevIncludedInvoiceResourcesReferencingParticipant ...
func (d *DevicePlusRelatedResources) GetRevIncludedInvoiceResourcesReferencingParticipant() (invoices []Invoice, err error) {
	if d.RevIncludedInvoiceResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded invoices not requested")
	} else {
		invoices = *d.RevIncludedInvoiceResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (d *DevicePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingSubject() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingAuthor() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedConsentResourcesReferencingActor ...
func (d *DevicePlusRelatedResources) GetRevIncludedConsentResourcesReferencingActor() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingActor == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingActor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingSubject() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingAuthor() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingSubject() (measureReports []MeasureReport, err error) {
	if d.RevIncludedMeasureReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *d.RevIncludedMeasureReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (d *DevicePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingRequester() (serviceRequests []ServiceRequest, err error) {
	if d.RevIncludedServiceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *d.RevIncludedServiceRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingPerformer() (serviceRequests []ServiceRequest, err error) {
	if d.RevIncludedServiceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *d.RevIncludedServiceRequestResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedServiceRequestResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingSubject() (serviceRequests []ServiceRequest, err error) {
	if d.RevIncludedServiceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *d.RevIncludedServiceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedSupplyRequestResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedSupplyRequestResourcesReferencingRequester() (supplyRequests []SupplyRequest, err error) {
	if d.RevIncludedSupplyRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded supplyRequests not requested")
	} else {
		supplyRequests = *d.RevIncludedSupplyRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (d *DevicePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if d.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *d.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedRiskAssessmentResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedRiskAssessmentResourcesReferencingPerformer() (riskAssessments []RiskAssessment, err error) {
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded riskAssessments not requested")
	} else {
		riskAssessments = *d.RevIncludedRiskAssessmentResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedGroupResourcesReferencingMember ...
func (d *DevicePlusRelatedResources) GetRevIncludedGroupResourcesReferencingMember() (groups []Group, err error) {
	if d.RevIncludedGroupResourcesReferencingMember == nil {
		err = errors.New("RevIncluded groups not requested")
	} else {
		groups = *d.RevIncludedGroupResourcesReferencingMember
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (d *DevicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (d *DevicePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (d *DevicePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingPerformer() (imagingStudies []ImagingStudy, err error) {
	if d.RevIncludedImagingStudyResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *d.RevIncludedImagingStudyResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedImagingStudyResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingSubject() (imagingStudies []ImagingStudy, err error) {
	if d.RevIncludedImagingStudyResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *d.RevIncludedImagingStudyResourcesReferencingSubject
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingEnterer ...
func (d *DevicePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingEnterer() (chargeItems []ChargeItem, err error) {
	if d.RevIncludedChargeItemResourcesReferencingEnterer == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *d.RevIncludedChargeItemResourcesReferencingEnterer
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingPerformeractor ...
func (d *DevicePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingPerformeractor() (chargeItems []ChargeItem, err error) {
	if d.RevIncludedChargeItemResourcesReferencingPerformeractor == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *d.RevIncludedChargeItemResourcesReferencingPerformeractor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (d *DevicePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingSender ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingSender() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingSender
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingRecipient ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingRecipient() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (d *DevicePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (d *DevicePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (d *DevicePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceUseStatementResourcesReferencingDevice ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceUseStatementResourcesReferencingDevice() (deviceUseStatements []DeviceUseStatement, err error) {
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded deviceUseStatements not requested")
	} else {
		deviceUseStatements = *d.RevIncludedDeviceUseStatementResourcesReferencingDevice
	}
	return
}

// GetRevIncludedRequestGroupResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingAuthor() (requestGroups []RequestGroup, err error) {
	if d.RevIncludedRequestGroupResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *d.RevIncludedRequestGroupResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedRequestGroupResourcesReferencingParticipant ...
func (d *DevicePlusRelatedResources) GetRevIncludedRequestGroupResourcesReferencingParticipant() (requestGroups []RequestGroup, err error) {
	if d.RevIncludedRequestGroupResourcesReferencingParticipant == nil {
		err = errors.New("RevIncluded requestGroups not requested")
	} else {
		requestGroups = *d.RevIncludedRequestGroupResourcesReferencingParticipant
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingRequester() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPerformer() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingSubject() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingDevice ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingDevice() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingDevice
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (d *DevicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingTarget ...
func (d *DevicePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingTarget() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingTarget
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (d *DevicePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingAgent ...
func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingAgent() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingAgent
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (d *DevicePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingOwner ...
func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingOwner() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingOwner == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingOwner
	}
	return
}

// GetRevIncludedTaskResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingRequester() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingRequester
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (d *DevicePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingDetailudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingDetailudi() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingProcedureudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingProcedureudi() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi
	}
	return
}

// GetRevIncludedExplanationOfBenefitResourcesReferencingItemudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingItemudi() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi
	}
	return
}

// GetRevIncludedSpecimenResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingSubject() (specimen []Specimen, err error) {
	if d.RevIncludedSpecimenResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *d.RevIncludedSpecimenResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPerformer() (carePlans []CarePlan, err error) {
	if d.RevIncludedCarePlanResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *d.RevIncludedCarePlanResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPerformer() (procedures []Procedure, err error) {
	if d.RevIncludedProcedureResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *d.RevIncludedProcedureResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedListResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingSubject() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingSubject
	}
	return
}

// GetRevIncludedListResourcesReferencingSource ...
func (d *DevicePlusRelatedResources) GetRevIncludedListResourcesReferencingSource() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingSource == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingSource
	}
	return
}

// GetRevIncludedMedicationRequestResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingRequester() (medicationRequests []MedicationRequest, err error) {
	if d.RevIncludedMedicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *d.RevIncludedMedicationRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedMedicationRequestResourcesReferencingIntendedperformer() (medicationRequests []MedicationRequest, err error) {
	if d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer == nil {
		err = errors.New("RevIncluded medicationRequests not requested")
	} else {
		medicationRequests = *d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer
	}
	return
}

// GetRevIncludedMediaResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if d.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *d.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMediaResourcesReferencingOperator ...
func (d *DevicePlusRelatedResources) GetRevIncludedMediaResourcesReferencingOperator() (media []Media, err error) {
	if d.RevIncludedMediaResourcesReferencingOperator == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *d.RevIncludedMediaResourcesReferencingOperator
	}
	return
}

// GetRevIncludedMediaResourcesReferencingDevice ...
func (d *DevicePlusRelatedResources) GetRevIncludedMediaResourcesReferencingDevice() (media []Media, err error) {
	if d.RevIncludedMediaResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *d.RevIncludedMediaResourcesReferencingDevice
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDeviceMetricResourcesReferencingParent ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceMetricResourcesReferencingParent() (deviceMetrics []DeviceMetric, err error) {
	if d.RevIncludedDeviceMetricResourcesReferencingParent == nil {
		err = errors.New("RevIncluded deviceMetrics not requested")
	} else {
		deviceMetrics = *d.RevIncludedDeviceMetricResourcesReferencingParent
	}
	return
}

// GetRevIncludedDeviceMetricResourcesReferencingSource ...
func (d *DevicePlusRelatedResources) GetRevIncludedDeviceMetricResourcesReferencingSource() (deviceMetrics []DeviceMetric, err error) {
	if d.RevIncludedDeviceMetricResourcesReferencingSource == nil {
		err = errors.New("RevIncluded deviceMetrics not requested")
	} else {
		deviceMetrics = *d.RevIncludedDeviceMetricResourcesReferencingSource
	}
	return
}

// GetRevIncludedFlagResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedFlagResourcesReferencingAuthor() (flags []Flag, err error) {
	if d.RevIncludedFlagResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *d.RevIncludedFlagResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedAppointmentResponseResourcesReferencingActor ...
func (d *DevicePlusRelatedResources) GetRevIncludedAppointmentResponseResourcesReferencingActor() (appointmentResponses []AppointmentResponse, err error) {
	if d.RevIncludedAppointmentResponseResourcesReferencingActor == nil {
		err = errors.New("RevIncluded appointmentResponses not requested")
	} else {
		appointmentResponses = *d.RevIncludedAppointmentResponseResourcesReferencingActor
	}
	return
}

// GetRevIncludedAdverseEventResourcesReferencingSubstance ...
func (d *DevicePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if d.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *d.RevIncludedAdverseEventResourcesReferencingSubstance
	}
	return
}

// GetRevIncludedObservationResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingSubject() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingSubject
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (d *DevicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingDevice ...
func (d *DevicePlusRelatedResources) GetRevIncludedObservationResourcesReferencingDevice() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingDevice
	}
	return
}

// GetRevIncludedMedicationAdministrationResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingPerformer() (medicationAdministrations []MedicationAdministration, err error) {
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedMedicationAdministrationResourcesReferencingDevice ...
func (d *DevicePlusRelatedResources) GetRevIncludedMedicationAdministrationResourcesReferencingDevice() (medicationAdministrations []MedicationAdministration, err error) {
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice == nil {
		err = errors.New("RevIncluded medicationAdministrations not requested")
	} else {
		medicationAdministrations = *d.RevIncludedMedicationAdministrationResourcesReferencingDevice
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (d *DevicePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingRequester ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRequester() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingRequester
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingSender ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingSender() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingSender == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingSender
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingRecipient ...
func (d *DevicePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingRecipient() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingRecipient
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedMedicationDispenseResourcesReferencingPerformer ...
func (d *DevicePlusRelatedResources) GetRevIncludedMedicationDispenseResourcesReferencingPerformer() (medicationDispenses []MedicationDispense, err error) {
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer == nil {
		err = errors.New("RevIncluded medicationDispenses not requested")
	} else {
		medicationDispenses = *d.RevIncludedMedicationDispenseResourcesReferencingPerformer
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSubject() (diagnosticReports []DiagnosticReport, err error) {
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *d.RevIncludedDiagnosticReportResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (d *DevicePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingAgent ...
func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingAgent() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingAgent == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingAgent
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingSource ...
func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingSource() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingSource == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingSource
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (d *DevicePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (d *DevicePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if d.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *d.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingAuthor() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (d *DevicePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingAuthor() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (d *DevicePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (d *DevicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor ...
func (d *DevicePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingAuthor() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor
	}
	return
}

// GetRevIncludedScheduleResourcesReferencingActor ...
func (d *DevicePlusRelatedResources) GetRevIncludedScheduleResourcesReferencingActor() (schedules []Schedule, err error) {
	if d.RevIncludedScheduleResourcesReferencingActor == nil {
		err = errors.New("RevIncluded schedules not requested")
	} else {
		schedules = *d.RevIncludedScheduleResourcesReferencingActor
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (d *DevicePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (d *DevicePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedClaimResourcesReferencingDetailudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedClaimResourcesReferencingDetailudi() (claims []Claim, err error) {
	if d.RevIncludedClaimResourcesReferencingDetailudi == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *d.RevIncludedClaimResourcesReferencingDetailudi
	}
	return
}

// GetRevIncludedClaimResourcesReferencingProcedureudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedClaimResourcesReferencingProcedureudi() (claims []Claim, err error) {
	if d.RevIncludedClaimResourcesReferencingProcedureudi == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *d.RevIncludedClaimResourcesReferencingProcedureudi
	}
	return
}

// GetRevIncludedClaimResourcesReferencingSubdetailudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedClaimResourcesReferencingSubdetailudi() (claims []Claim, err error) {
	if d.RevIncludedClaimResourcesReferencingSubdetailudi == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *d.RevIncludedClaimResourcesReferencingSubdetailudi
	}
	return
}

// GetRevIncludedClaimResourcesReferencingItemudi ...
func (d *DevicePlusRelatedResources) GetRevIncludedClaimResourcesReferencingItemudi() (claims []Claim, err error) {
	if d.RevIncludedClaimResourcesReferencingItemudi == nil {
		err = errors.New("RevIncluded claims not requested")
	} else {
		claims = *d.RevIncludedClaimResourcesReferencingItemudi
	}
	return
}

// GetIncludedResources ...
func (d *DevicePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*d.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *d.IncludedLocationResourcesReferencedByLocation {
			rsc := (*d.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (d *DevicePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*d.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActor {
			rsc := (*d.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*d.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*d.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSubject {
			rsc := (*d.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSource {
			rsc := (*d.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*d.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingDevice {
			rsc := (*d.RevIncludedMediaResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingParent != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingParent {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingParent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*d.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*d.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingDevice {
			rsc := (*d.RevIncludedObservationResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*d.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingDetailudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingDetailudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingDetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingProcedureudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingProcedureudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingProcedureudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingSubdetailudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingSubdetailudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingSubdetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingItemudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingItemudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingItemudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (d *DevicePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *d.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*d.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *d.IncludedLocationResourcesReferencedByLocation {
			rsc := (*d.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAccountResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedAccountResourcesReferencingSubject {
			rsc := (*d.RevIncludedAccountResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedInvoiceResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedInvoiceResourcesReferencingParticipant {
			rsc := (*d.RevIncludedInvoiceResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingActor {
			rsc := (*d.RevIncludedConsentResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingSubject {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedServiceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedServiceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedServiceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedSupplyRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedSupplyRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedSupplyRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRiskAssessmentResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedRiskAssessmentResourcesReferencingPerformer {
			rsc := (*d.RevIncludedRiskAssessmentResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedGroupResourcesReferencingMember != nil {
		for idx := range *d.RevIncludedGroupResourcesReferencingMember {
			rsc := (*d.RevIncludedGroupResourcesReferencingMember)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingPerformer {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedImagingStudyResourcesReferencingSubject {
			rsc := (*d.RevIncludedImagingStudyResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingEnterer != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingEnterer {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingEnterer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedChargeItemResourcesReferencingPerformeractor != nil {
		for idx := range *d.RevIncludedChargeItemResourcesReferencingPerformeractor {
			rsc := (*d.RevIncludedChargeItemResourcesReferencingPerformeractor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceUseStatementResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceUseStatementResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceUseStatementResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingAuthor {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedRequestGroupResourcesReferencingParticipant != nil {
		for idx := range *d.RevIncludedRequestGroupResourcesReferencingParticipant {
			rsc := (*d.RevIncludedRequestGroupResourcesReferencingParticipant)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPerformer {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingSubject {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingDevice {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingTarget {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingAgent {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingOwner != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingOwner {
			rsc := (*d.RevIncludedTaskResourcesReferencingOwner)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingRequester {
			rsc := (*d.RevIncludedTaskResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingDetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingProcedureudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingSubdetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi != nil {
		for idx := range *d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi {
			rsc := (*d.RevIncludedExplanationOfBenefitResourcesReferencingItemudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedSpecimenResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedSpecimenResourcesReferencingSubject {
			rsc := (*d.RevIncludedSpecimenResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedCarePlanResourcesReferencingPerformer {
			rsc := (*d.RevIncludedCarePlanResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedProcedureResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedProcedureResourcesReferencingPerformer {
			rsc := (*d.RevIncludedProcedureResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSubject {
			rsc := (*d.RevIncludedListResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedListResourcesReferencingSource {
			rsc := (*d.RevIncludedListResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer != nil {
		for idx := range *d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer {
			rsc := (*d.RevIncludedMedicationRequestResourcesReferencingIntendedperformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*d.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingOperator != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingOperator {
			rsc := (*d.RevIncludedMediaResourcesReferencingOperator)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMediaResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMediaResourcesReferencingDevice {
			rsc := (*d.RevIncludedMediaResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingParent != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingParent {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingParent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDeviceMetricResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedDeviceMetricResourcesReferencingSource {
			rsc := (*d.RevIncludedDeviceMetricResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedFlagResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedFlagResourcesReferencingAuthor {
			rsc := (*d.RevIncludedFlagResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAppointmentResponseResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedAppointmentResponseResourcesReferencingActor {
			rsc := (*d.RevIncludedAppointmentResponseResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *d.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*d.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingSubject {
			rsc := (*d.RevIncludedObservationResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingDevice {
			rsc := (*d.RevIncludedObservationResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationAdministrationResourcesReferencingDevice != nil {
		for idx := range *d.RevIncludedMedicationAdministrationResourcesReferencingDevice {
			rsc := (*d.RevIncludedMedicationAdministrationResourcesReferencingDevice)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRequester != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRequester {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRequester)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingSender != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingSender {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingSender)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingRecipient != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingRecipient {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingRecipient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedMedicationDispenseResourcesReferencingPerformer != nil {
		for idx := range *d.RevIncludedMedicationDispenseResourcesReferencingPerformer {
			rsc := (*d.RevIncludedMedicationDispenseResourcesReferencingPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedDiagnosticReportResourcesReferencingSubject {
			rsc := (*d.RevIncludedDiagnosticReportResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingAgent != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingAgent {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingAgent)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingSource {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingAuthor {
			rsc := (*d.RevIncludedCompositionResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingAuthor {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingAuthor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedScheduleResourcesReferencingActor != nil {
		for idx := range *d.RevIncludedScheduleResourcesReferencingActor {
			rsc := (*d.RevIncludedScheduleResourcesReferencingActor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingDetailudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingDetailudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingDetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingProcedureudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingProcedureudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingProcedureudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingSubdetailudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingSubdetailudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingSubdetailudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if d.RevIncludedClaimResourcesReferencingItemudi != nil {
		for idx := range *d.RevIncludedClaimResourcesReferencingItemudi {
			rsc := (*d.RevIncludedClaimResourcesReferencingItemudi)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
