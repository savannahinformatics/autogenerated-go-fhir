package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Specimen struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	AccessionIdentifier *Identifier                   `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	Status              string                        `bson:"status,omitempty" json:"status,omitempty"`
	Type                *CodeableConcept              `bson:"type,omitempty" json:"type,omitempty"`
	Subject             *Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	ReceivedTime        *FHIRDateTime                 `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	Parent              []Reference                   `bson:"parent,omitempty" json:"parent,omitempty"`
	Request             []Reference                   `bson:"request,omitempty" json:"request,omitempty"`
	Collection          *SpecimenCollectionComponent  `bson:"collection,omitempty" json:"collection,omitempty"`
	Processing          []SpecimenProcessingComponent `bson:"processing,omitempty" json:"processing,omitempty"`
	Container           []SpecimenContainerComponent  `bson:"container,omitempty" json:"container,omitempty"`
	Condition           []CodeableConcept             `bson:"condition,omitempty" json:"condition,omitempty"`
	Note                []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Specimen) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Specimen"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "specimen" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type specimen Specimen

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Specimen) UnmarshalJSON(data []byte) (err error) {
	x2 := specimen{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Specimen(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Specimen) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Specimen"
	} else if x.ResourceType != "Specimen" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Specimen, instead received %s", x.ResourceType))
	}
	return nil
}

type SpecimenCollectionComponent struct {
	BackboneElement              `bson:",inline"`
	Collector                    *Reference       `bson:"collector,omitempty" json:"collector,omitempty"`
	CollectedDateTime            *FHIRDateTime    `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod              *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Duration                     *Quantity        `bson:"duration,omitempty" json:"duration,omitempty"`
	Quantity                     *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Method                       *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	BodySite                     *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `bson:"fastingStatusCodeableConcept,omitempty" json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Quantity        `bson:"fastingStatusDuration,omitempty" json:"fastingStatusDuration,omitempty"`
}

type SpecimenProcessingComponent struct {
	BackboneElement `bson:",inline"`
	Description     string           `bson:"description,omitempty" json:"description,omitempty"`
	Procedure       *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive        []Reference      `bson:"additive,omitempty" json:"additive,omitempty"`
	TimeDateTime    *FHIRDateTime    `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod      *Period          `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}

type SpecimenContainerComponent struct {
	BackboneElement         `bson:",inline"`
	Identifier              []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description             string           `bson:"description,omitempty" json:"description,omitempty"`
	Type                    *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity                *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

type SpecimenPlus struct {
	Specimen                     `bson:",inline"`
	SpecimenPlusRelatedResources `bson:",inline"`
}

type SpecimenPlusRelatedResources struct {
	IncludedSpecimenResourcesReferencedByParent                            *[]Specimen                   `bson:"_includedSpecimenResourcesReferencedByParent,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                             *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedSubstanceResourcesReferencedBySubject                          *[]Substance                  `bson:"_includedSubstanceResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                           *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByCollector                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByCollector,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByCollector                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByCollector,omitempty"`
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
	RevIncludedServiceRequestResourcesReferencingSpecimen                  *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingSpecimen,omitempty"`
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
	RevIncludedSpecimenResourcesReferencingParent                          *[]Specimen                   `bson:"_revIncludedSpecimenResourcesReferencingParent,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedMediaResourcesReferencingSubject                            *[]Media                      `bson:"_revIncludedMediaResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingSpecimen                     *[]Observation                `bson:"_revIncludedObservationResourcesReferencingSpecimen,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingSpecimen                *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingSpecimen,omitempty"`
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

func (s *SpecimenPlusRelatedResources) GetIncludedSpecimenResourcesReferencedByParent() (specimen []Specimen, err error) {
	if s.IncludedSpecimenResourcesReferencedByParent == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *s.IncludedSpecimenResourcesReferencedByParent
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if s.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*s.IncludedGroupResourcesReferencedBySubject))
	} else if len(*s.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*s.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if s.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*s.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*s.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*s.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedBySubject))
	} else if len(*s.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedSubstanceResourceReferencedBySubject() (substance *Substance, err error) {
	if s.IncludedSubstanceResourcesReferencedBySubject == nil {
		err = errors.New("Included substances not requested")
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*s.IncludedSubstanceResourcesReferencedBySubject))
	} else if len(*s.IncludedSubstanceResourcesReferencedBySubject) == 1 {
		substance = &(*s.IncludedSubstanceResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if s.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*s.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*s.IncludedLocationResourcesReferencedBySubject))
	} else if len(*s.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*s.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if s.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*s.IncludedPatientResourcesReferencedByPatient))
	} else if len(*s.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*s.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCollector() (practitioner *Practitioner, err error) {
	if s.IncludedPractitionerResourcesReferencedByCollector == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*s.IncludedPractitionerResourcesReferencedByCollector) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*s.IncludedPractitionerResourcesReferencedByCollector))
	} else if len(*s.IncludedPractitionerResourcesReferencedByCollector) == 1 {
		practitioner = &(*s.IncludedPractitionerResourcesReferencedByCollector)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByCollector() (practitionerRole *PractitionerRole, err error) {
	if s.IncludedPractitionerRoleResourcesReferencedByCollector == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*s.IncludedPractitionerRoleResourcesReferencedByCollector) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*s.IncludedPractitionerRoleResourcesReferencedByCollector))
	} else if len(*s.IncludedPractitionerRoleResourcesReferencedByCollector) == 1 {
		practitionerRole = &(*s.IncludedPractitionerRoleResourcesReferencedByCollector)[0]
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *s.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if s.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *s.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if s.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *s.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *s.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingSpecimen() (serviceRequests []ServiceRequest, err error) {
	if s.RevIncludedServiceRequestResourcesReferencingSpecimen == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *s.RevIncludedServiceRequestResourcesReferencingSpecimen
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if s.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *s.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if s.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *s.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if s.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *s.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if s.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *s.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if s.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *s.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if s.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *s.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedSpecimenResourcesReferencingParent() (specimen []Specimen, err error) {
	if s.RevIncludedSpecimenResourcesReferencingParent == nil {
		err = errors.New("RevIncluded specimen not requested")
	} else {
		specimen = *s.RevIncludedSpecimenResourcesReferencingParent
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedMediaResourcesReferencingSubject() (media []Media, err error) {
	if s.RevIncludedMediaResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded media not requested")
	} else {
		media = *s.RevIncludedMediaResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *s.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedObservationResourcesReferencingSpecimen() (observations []Observation, err error) {
	if s.RevIncludedObservationResourcesReferencingSpecimen == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *s.RevIncludedObservationResourcesReferencingSpecimen
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if s.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *s.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *s.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingSpecimen() (diagnosticReports []DiagnosticReport, err error) {
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if s.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *s.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if s.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *s.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (s *SpecimenPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSpecimenResourcesReferencedByParent != nil {
		for idx := range *s.IncludedSpecimenResourcesReferencedByParent {
			rsc := (*s.IncludedSpecimenResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedGroupResourcesReferencedBySubject {
			rsc := (*s.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*s.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedPatientResourcesReferencedBySubject {
			rsc := (*s.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedSubstanceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubject {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedLocationResourcesReferencedBySubject {
			rsc := (*s.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPatient {
			rsc := (*s.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerResourcesReferencedByCollector != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByCollector {
			rsc := (*s.IncludedPractitionerResourcesReferencedByCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByCollector != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByCollector {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SpecimenPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingParent != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingParent {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*s.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedObservationResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedDiagnosticReportResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (s *SpecimenPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.IncludedSpecimenResourcesReferencedByParent != nil {
		for idx := range *s.IncludedSpecimenResourcesReferencedByParent {
			rsc := (*s.IncludedSpecimenResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedGroupResourcesReferencedBySubject {
			rsc := (*s.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*s.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedPatientResourcesReferencedBySubject {
			rsc := (*s.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedSubstanceResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedSubstanceResourcesReferencedBySubject {
			rsc := (*s.IncludedSubstanceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *s.IncludedLocationResourcesReferencedBySubject {
			rsc := (*s.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *s.IncludedPatientResourcesReferencedByPatient {
			rsc := (*s.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerResourcesReferencedByCollector != nil {
		for idx := range *s.IncludedPractitionerResourcesReferencedByCollector {
			rsc := (*s.IncludedPractitionerResourcesReferencedByCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.IncludedPractitionerRoleResourcesReferencedByCollector != nil {
		for idx := range *s.IncludedPractitionerRoleResourcesReferencedByCollector {
			rsc := (*s.IncludedPractitionerRoleResourcesReferencedByCollector)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*s.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*s.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *s.RevIncludedConsentResourcesReferencingData {
			rsc := (*s.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*s.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *s.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*s.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*s.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedServiceRequestResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedServiceRequestResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedServiceRequestResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*s.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedContractResourcesReferencingSubject {
			rsc := (*s.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*s.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *s.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*s.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*s.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *s.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*s.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *s.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*s.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*s.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *s.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*s.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*s.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*s.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*s.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*s.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedSpecimenResourcesReferencingParent != nil {
		for idx := range *s.RevIncludedSpecimenResourcesReferencingParent {
			rsc := (*s.RevIncludedSpecimenResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *s.RevIncludedListResourcesReferencingItem {
			rsc := (*s.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedMediaResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedMediaResourcesReferencingSubject {
			rsc := (*s.RevIncludedMediaResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*s.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedObservationResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedObservationResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedObservationResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*s.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*s.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *s.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*s.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*s.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDiagnosticReportResourcesReferencingSpecimen != nil {
		for idx := range *s.RevIncludedDiagnosticReportResourcesReferencingSpecimen {
			rsc := (*s.RevIncludedDiagnosticReportResourcesReferencingSpecimen)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *s.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*s.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *s.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*s.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *s.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*s.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*s.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *s.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*s.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*s.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*s.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*s.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*s.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
