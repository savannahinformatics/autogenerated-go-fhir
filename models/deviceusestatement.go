package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DeviceUseStatement struct {
	DomainResource  `bson:",inline"`
	Identifier      []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn         []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status          string            `bson:"status,omitempty" json:"status,omitempty"`
	Subject         *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	DerivedFrom     []Reference       `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	TimingTiming    *Timing           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingPeriod    *Period           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDateTime  *FHIRDateTime     `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	RecordedOn      *FHIRDateTime     `bson:"recordedOn,omitempty" json:"recordedOn,omitempty"`
	Source          *Reference        `bson:"source,omitempty" json:"source,omitempty"`
	Device          *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	ReasonCode      []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite        *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note            []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceUseStatement) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DeviceUseStatement"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "deviceUseStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceUseStatement DeviceUseStatement

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceUseStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceUseStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceUseStatement(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DeviceUseStatement) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DeviceUseStatement"
	} else if x.ResourceType != "DeviceUseStatement" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DeviceUseStatement, instead received %s", x.ResourceType))
	}
	return nil
}

type DeviceUseStatementPlus struct {
	DeviceUseStatement                     `bson:",inline"`
	DeviceUseStatementPlusRelatedResources `bson:",inline"`
}

type DeviceUseStatementPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedByPatient                              *[]Group                      `bson:"_includedGroupResourcesReferencedByPatient,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
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

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedGroupResourceReferencedByPatient() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedByPatient == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedByPatient))
	} else if len(*d.IncludedGroupResourcesReferencedByPatient) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedDeviceResourceReferencedByDevice() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByDevice == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByDevice) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByDevice))
	} else if len(*d.IncludedDeviceResourcesReferencedByDevice) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByDevice)[0]
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *d.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if d.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *d.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if d.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *d.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedGroupResourcesReferencedByPatient {
			rsc := (*d.IncludedGroupResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*d.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DeviceUseStatementPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DeviceUseStatementPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedGroupResourcesReferencedBySubject {
			rsc := (*d.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *d.IncludedPatientResourcesReferencedBySubject {
			rsc := (*d.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedGroupResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedGroupResourcesReferencedByPatient {
			rsc := (*d.IncludedGroupResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *d.IncludedPatientResourcesReferencedByPatient {
			rsc := (*d.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.IncludedDeviceResourcesReferencedByDevice != nil {
		for idx := range *d.IncludedDeviceResourcesReferencedByDevice {
			rsc := (*d.IncludedDeviceResourcesReferencedByDevice)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
