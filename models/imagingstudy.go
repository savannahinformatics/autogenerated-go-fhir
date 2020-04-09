package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ImagingStudy ...
type ImagingStudy struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             string                        `bson:"status,omitempty" json:"status,omitempty"`
	Modality           []Coding                      `bson:"modality,omitempty" json:"modality,omitempty"`
	Subject            *Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter          *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Started            *FHIRDateTime                 `bson:"started,omitempty" json:"started,omitempty"`
	BasedOn            []Reference                   `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Referrer           *Reference                    `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Interpreter        []Reference                   `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Endpoint           []Reference                   `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	NumberOfSeries     *uint32                       `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances  *uint32                       `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	ProcedureReference *Reference                    `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept             `bson:"procedureCode,omitempty" json:"procedureCode,omitempty"`
	Location           *Reference                    `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []CodeableConcept             `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note               []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
	Description        string                        `bson:"description,omitempty" json:"description,omitempty"`
	Series             []ImagingStudySeriesComponent `bson:"series,omitempty" json:"series,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *ImagingStudy) MarshalJSON() ([]byte, error) {
	x.ResourceType = "ImagingStudy"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "imagingStudy" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type imagingStudy ImagingStudy

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *ImagingStudy) UnmarshalJSON(data []byte) (err error) {
	x2 := imagingStudy{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImagingStudy(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImagingStudy) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImagingStudy"
	} else if x.ResourceType != "ImagingStudy" {
		return fmt.Errorf("Expected resourceType to be ImagingStudy, instead received %s", x.ResourceType)
	}
	return nil
}

// ImagingStudySeriesComponent ...
type ImagingStudySeriesComponent struct {
	BackboneElement   `bson:",inline"`
	UID               string                                 `bson:"uid,omitempty" json:"uid,omitempty"`
	Number            *uint32                                `bson:"number,omitempty" json:"number,omitempty"`
	Modality          *Coding                                `bson:"modality,omitempty" json:"modality,omitempty"`
	Description       string                                 `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *uint32                                `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Endpoint          []Reference                            `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	BodySite          *Coding                                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Laterality        *Coding                                `bson:"laterality,omitempty" json:"laterality,omitempty"`
	Specimen          []Reference                            `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Started           *FHIRDateTime                          `bson:"started,omitempty" json:"started,omitempty"`
	Performer         []ImagingStudySeriesPerformerComponent `bson:"performer,omitempty" json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstanceComponent  `bson:"instance,omitempty" json:"instance,omitempty"`
}

// ImagingStudySeriesPerformerComponent ...
type ImagingStudySeriesPerformerComponent struct {
	BackboneElement `bson:",inline"`
	Function        *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

// ImagingStudySeriesInstanceComponent ...
type ImagingStudySeriesInstanceComponent struct {
	BackboneElement `bson:",inline"`
	UID             string  `bson:"uid,omitempty" json:"uid,omitempty"`
	SopClass        *Coding `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Number          *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	Title           string  `bson:"title,omitempty" json:"title,omitempty"`
}

// ImagingStudyPlus ...
type ImagingStudyPlus struct {
	ImagingStudy                     `bson:",inline"`
	ImagingStudyPlusRelatedResources `bson:",inline"`
}

// ImagingStudyPlusRelatedResources ...
type ImagingStudyPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                         *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                             *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPractitionerResourcesReferencedByInterpreter                   *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByInterpreter,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByInterpreter               *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByInterpreter,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByReferrer                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByReferrer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByReferrer                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByReferrer,omitempty"`
	IncludedEndpointResourcesReferencedByEndpointPath1                     *[]Endpoint                   `bson:"_includedEndpointResourcesReferencedByEndpointPath1,omitempty"`
	IncludedEndpointResourcesReferencedByEndpointPath2                     *[]Endpoint                   `bson:"_includedEndpointResourcesReferencedByEndpointPath2,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedAppointmentResourcesReferencedByBasedon                        *[]Appointment                `bson:"_includedAppointmentResourcesReferencedByBasedon,omitempty"`
	IncludedAppointmentResponseResourcesReferencedByBasedon                *[]AppointmentResponse        `bson:"_includedAppointmentResponseResourcesReferencedByBasedon,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedTaskResourcesReferencedByBasedon                               *[]Task                       `bson:"_includedTaskResourcesReferencedByBasedon,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
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
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingDerivedfrom                  *[]Observation                `bson:"_revIncludedObservationResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedObservationResourcesReferencingPartof                       *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPartof,omitempty"`
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

// GetIncludedPractitionerResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*i.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedOrganizationResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPerformer() (organization *Organization, err error) {
	if i.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*i.IncludedOrganizationResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*i.IncludedOrganizationResourcesReferencedByPerformer))
	} else if len(*i.IncludedOrganizationResourcesReferencedByPerformer) == 1 {
		organization = &(*i.IncludedOrganizationResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedCareTeamResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedCareTeamResourceReferencedByPerformer() (careTeam *CareTeam, err error) {
	if i.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careteams not requested")
	} else if len(*i.IncludedCareTeamResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 careTeam, but found %d", len(*i.IncludedCareTeamResourcesReferencedByPerformer))
	} else if len(*i.IncludedCareTeamResourcesReferencedByPerformer) == 1 {
		careTeam = &(*i.IncludedCareTeamResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedDeviceResourceReferencedByPerformer() (device *Device, err error) {
	if i.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else if len(*i.IncludedDeviceResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*i.IncludedDeviceResourcesReferencedByPerformer))
	} else if len(*i.IncludedDeviceResourcesReferencedByPerformer) == 1 {
		device = &(*i.IncludedDeviceResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPatientResourceReferencedByPerformer() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPerformer))
	} else if len(*i.IncludedPatientResourcesReferencedByPerformer) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPerformer() (practitionerRole *PractitionerRole, err error) {
	if i.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*i.IncludedPractitionerRoleResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*i.IncludedPractitionerRoleResourcesReferencedByPerformer))
	} else if len(*i.IncludedPractitionerRoleResourcesReferencedByPerformer) == 1 {
		practitionerRole = &(*i.IncludedPractitionerRoleResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedRelatedPersonResourceReferencedByPerformer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPerformer() (relatedPerson *RelatedPerson, err error) {
	if i.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*i.IncludedRelatedPersonResourcesReferencedByPerformer))
	} else if len(*i.IncludedRelatedPersonResourcesReferencedByPerformer) == 1 {
		relatedPerson = &(*i.IncludedRelatedPersonResourcesReferencedByPerformer)[0]
	}
	return
}

// GetIncludedGroupResourceReferencedBySubject ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if i.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*i.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*i.IncludedGroupResourcesReferencedBySubject))
	} else if len(*i.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*i.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedDeviceResourceReferencedBySubject ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if i.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*i.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*i.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*i.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*i.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedBySubject ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedBySubject))
	} else if len(*i.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

// GetIncludedPractitionerResourcesReferencedByInterpreter ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByInterpreter() (practitioners []Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByInterpreter == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *i.IncludedPractitionerResourcesReferencedByInterpreter
	}
	return
}

// GetIncludedPractitionerRoleResourcesReferencedByInterpreter ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByInterpreter() (practitionerRoles []PractitionerRole, err error) {
	if i.IncludedPractitionerRoleResourcesReferencedByInterpreter == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *i.IncludedPractitionerRoleResourcesReferencedByInterpreter
	}
	return
}

// GetIncludedEncounterResourceReferencedByEncounter ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if i.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*i.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*i.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*i.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*i.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

// GetIncludedPractitionerResourceReferencedByReferrer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerResourceReferencedByReferrer() (practitioner *Practitioner, err error) {
	if i.IncludedPractitionerResourcesReferencedByReferrer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*i.IncludedPractitionerResourcesReferencedByReferrer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*i.IncludedPractitionerResourcesReferencedByReferrer))
	} else if len(*i.IncludedPractitionerResourcesReferencedByReferrer) == 1 {
		practitioner = &(*i.IncludedPractitionerResourcesReferencedByReferrer)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByReferrer ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByReferrer() (practitionerRole *PractitionerRole, err error) {
	if i.IncludedPractitionerRoleResourcesReferencedByReferrer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*i.IncludedPractitionerRoleResourcesReferencedByReferrer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*i.IncludedPractitionerRoleResourcesReferencedByReferrer))
	} else if len(*i.IncludedPractitionerRoleResourcesReferencedByReferrer) == 1 {
		practitionerRole = &(*i.IncludedPractitionerRoleResourcesReferencedByReferrer)[0]
	}
	return
}

// GetIncludedEndpointResourcesReferencedByEndpointPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpointPath1() (endpoints []Endpoint, err error) {
	if i.IncludedEndpointResourcesReferencedByEndpointPath1 == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *i.IncludedEndpointResourcesReferencedByEndpointPath1
	}
	return
}

// GetIncludedEndpointResourcesReferencedByEndpointPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedEndpointResourcesReferencedByEndpointPath2() (endpoints []Endpoint, err error) {
	if i.IncludedEndpointResourcesReferencedByEndpointPath2 == nil {
		err = errors.New("Included endpoints not requested")
	} else {
		endpoints = *i.IncludedEndpointResourcesReferencedByEndpointPath2
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPatient))
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedAppointmentResourcesReferencedByBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedAppointmentResourcesReferencedByBasedon() (appointments []Appointment, err error) {
	if i.IncludedAppointmentResourcesReferencedByBasedon == nil {
		err = errors.New("Included appointments not requested")
	} else {
		appointments = *i.IncludedAppointmentResourcesReferencedByBasedon
	}
	return
}

// GetIncludedAppointmentResponseResourcesReferencedByBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedAppointmentResponseResourcesReferencedByBasedon() (appointmentResponses []AppointmentResponse, err error) {
	if i.IncludedAppointmentResponseResourcesReferencedByBasedon == nil {
		err = errors.New("Included appointmentResponses not requested")
	} else {
		appointmentResponses = *i.IncludedAppointmentResponseResourcesReferencedByBasedon
	}
	return
}

// GetIncludedCarePlanResourcesReferencedByBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if i.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *i.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

// GetIncludedTaskResourcesReferencedByBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedTaskResourcesReferencedByBasedon() (tasks []Task, err error) {
	if i.IncludedTaskResourcesReferencedByBasedon == nil {
		err = errors.New("Included tasks not requested")
	} else {
		tasks = *i.IncludedTaskResourcesReferencedByBasedon
	}
	return
}

// GetIncludedServiceRequestResourcesReferencedByBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if i.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *i.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *i.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if i.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *i.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if i.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *i.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedChargeItemResourcesReferencingService ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if i.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *i.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedObservationResourcesReferencingDerivedfrom() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedObservationResourcesReferencingPartof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedObservationResourcesReferencingPartof() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *i.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if i.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *i.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingInvestigation ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingInvestigation() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingInvestigation == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingInvestigation
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*i.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*i.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*i.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*i.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*i.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*i.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedGroupResourcesReferencedBySubject {
			rsc := (*i.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*i.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedPatientResourcesReferencedBySubject {
			rsc := (*i.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerResourcesReferencedByInterpreter != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByInterpreter {
			rsc := (*i.IncludedPractitionerResourcesReferencedByInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByInterpreter != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByInterpreter {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *i.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*i.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerResourcesReferencedByReferrer != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByReferrer {
			rsc := (*i.IncludedPractitionerResourcesReferencedByReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByReferrer != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByReferrer {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEndpointResourcesReferencedByEndpointPath1 != nil {
		for idx := range *i.IncludedEndpointResourcesReferencedByEndpointPath1 {
			rsc := (*i.IncludedEndpointResourcesReferencedByEndpointPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEndpointResourcesReferencedByEndpointPath2 != nil {
		for idx := range *i.IncludedEndpointResourcesReferencedByEndpointPath2 {
			rsc := (*i.IncludedEndpointResourcesReferencedByEndpointPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedAppointmentResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedAppointmentResourcesReferencedByBasedon {
			rsc := (*i.IncludedAppointmentResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedAppointmentResponseResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedAppointmentResponseResourcesReferencedByBasedon {
			rsc := (*i.IncludedAppointmentResponseResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*i.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedTaskResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedTaskResourcesReferencedByBasedon {
			rsc := (*i.IncludedTaskResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*i.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (i *ImagingStudyPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *i.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*i.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*i.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (i *ImagingStudyPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*i.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*i.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*i.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*i.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*i.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *i.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*i.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedGroupResourcesReferencedBySubject {
			rsc := (*i.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*i.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *i.IncludedPatientResourcesReferencedBySubject {
			rsc := (*i.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerResourcesReferencedByInterpreter != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByInterpreter {
			rsc := (*i.IncludedPractitionerResourcesReferencedByInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByInterpreter != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByInterpreter {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByInterpreter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *i.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*i.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerResourcesReferencedByReferrer != nil {
		for idx := range *i.IncludedPractitionerResourcesReferencedByReferrer {
			rsc := (*i.IncludedPractitionerResourcesReferencedByReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPractitionerRoleResourcesReferencedByReferrer != nil {
		for idx := range *i.IncludedPractitionerRoleResourcesReferencedByReferrer {
			rsc := (*i.IncludedPractitionerRoleResourcesReferencedByReferrer)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEndpointResourcesReferencedByEndpointPath1 != nil {
		for idx := range *i.IncludedEndpointResourcesReferencedByEndpointPath1 {
			rsc := (*i.IncludedEndpointResourcesReferencedByEndpointPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedEndpointResourcesReferencedByEndpointPath2 != nil {
		for idx := range *i.IncludedEndpointResourcesReferencedByEndpointPath2 {
			rsc := (*i.IncludedEndpointResourcesReferencedByEndpointPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedAppointmentResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedAppointmentResourcesReferencedByBasedon {
			rsc := (*i.IncludedAppointmentResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedAppointmentResponseResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedAppointmentResponseResourcesReferencedByBasedon {
			rsc := (*i.IncludedAppointmentResponseResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*i.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedTaskResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedTaskResourcesReferencedByBasedon {
			rsc := (*i.IncludedTaskResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *i.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*i.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *i.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*i.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedObservationResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*i.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingInvestigation != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingInvestigation {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingInvestigation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
