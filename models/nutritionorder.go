package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NutritionOrder ...
type NutritionOrder struct {
	DomainResource         `bson:",inline"`
	Identifier             []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical  []Canonical                            `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesURI        []string                               `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Instantiates           []string                               `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	Status                 string                                 `bson:"status,omitempty" json:"status,omitempty"`
	Intent                 string                                 `bson:"intent,omitempty" json:"intent,omitempty"`
	Patient                *Reference                             `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter              *Reference                             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               *FHIRDateTime                          `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Orderer                *Reference                             `bson:"orderer,omitempty" json:"orderer,omitempty"`
	AllergyIntolerance     []Reference                            `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept                      `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept                      `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDietComponent       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplementComponent    `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormulaComponent `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
	Note                   []Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *NutritionOrder) MarshalJSON() ([]byte, error) {
	x.ResourceType = "NutritionOrder"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "nutritionOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type nutritionOrder NutritionOrder

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *NutritionOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := nutritionOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NutritionOrder(x2)
		return x.checkResourceType()
	}
	return
}

func (x *NutritionOrder) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "NutritionOrder"
	} else if x.ResourceType != "NutritionOrder" {
		return fmt.Errorf("Expected resourceType to be NutritionOrder, instead received %s", x.ResourceType)
	}
	return nil
}

// NutritionOrderOralDietComponent ...
type NutritionOrderOralDietComponent struct {
	BackboneElement      `bson:",inline"`
	Type                 []CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []Timing                                  `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrientComponent `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTextureComponent  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                         `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          string                                    `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

// NutritionOrderOralDietNutrientComponent ...
type NutritionOrderOralDietNutrientComponent struct {
	BackboneElement `bson:",inline"`
	Modifier        *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

// NutritionOrderOralDietTextureComponent ...
type NutritionOrderOralDietTextureComponent struct {
	BackboneElement `bson:",inline"`
	Modifier        *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType        *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}

// NutritionOrderSupplementComponent ...
type NutritionOrderSupplementComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName     string           `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule        []Timing         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity        *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction     string           `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

// NutritionOrderEnteralFormulaComponent ...
type NutritionOrderEnteralFormulaComponent struct {
	BackboneElement           `bson:",inline"`
	BaseFormulaType           *CodeableConcept                                      `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    string                                                `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                                      `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       string                                                `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                             `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                                      `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministrationComponent `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                             `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction string                                                `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}

// NutritionOrderEnteralFormulaAdministrationComponent ...
type NutritionOrderEnteralFormulaAdministrationComponent struct {
	BackboneElement    `bson:",inline"`
	Schedule           *Timing   `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity           *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	RateSimpleQuantity *Quantity `bson:"rateSimpleQuantity,omitempty" json:"rateSimpleQuantity,omitempty"`
	RateRatio          *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
}

// NutritionOrderPlus ...
type NutritionOrderPlus struct {
	NutritionOrder                     `bson:",inline"`
	NutritionOrderPlusRelatedResources `bson:",inline"`
}

// NutritionOrderPlusRelatedResources ...
type NutritionOrderPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByProvider                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByProvider                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical       *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical   *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingActivityreference               *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
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

// GetIncludedPractitionerResourceReferencedByProvider ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if n.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*n.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*n.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*n.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*n.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

// GetIncludedPractitionerRoleResourceReferencedByProvider ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByProvider() (practitionerRole *PractitionerRole, err error) {
	if n.IncludedPractitionerRoleResourcesReferencedByProvider == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*n.IncludedPractitionerRoleResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*n.IncludedPractitionerRoleResourcesReferencedByProvider))
	} else if len(*n.IncludedPractitionerRoleResourcesReferencedByProvider) == 1 {
		practitionerRole = &(*n.IncludedPractitionerRoleResourcesReferencedByProvider)[0]
	}
	return
}

// GetIncludedPatientResourceReferencedByPatient ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if n.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*n.IncludedPatientResourcesReferencedByPatient))
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*n.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

// GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

// GetIncludedEncounterResourceReferencedByEncounter ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if n.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*n.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*n.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *n.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if n.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *n.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *n.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if n.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *n.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if n.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *n.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if n.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *n.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if n.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *n.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingActivityreference ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if n.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *n.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if n.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *n.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if n.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *n.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedObservationResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if n.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *n.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *n.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if n.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *n.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedDiagnosticReportResourcesReferencingBasedon ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if n.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *n.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if n.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *n.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if n.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *n.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *n.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetIncludedResources ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetRevIncludedResources ...
func (n *NutritionOrderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *n.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*n.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ...
func (n *NutritionOrderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *n.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*n.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*n.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*n.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *n.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*n.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.ID] = &rsc
		}
	}
	return resourceMap
}
