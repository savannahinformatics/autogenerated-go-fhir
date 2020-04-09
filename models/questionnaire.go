package fhir

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Questionnaire ... // TODO Write proper comment
type Questionnaire struct {
	DomainResource  `bson:",inline"`
	Url             string                       `bson:"url,omitempty" json:"url,omitempty"`
	Identifier      []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version         string                       `bson:"version,omitempty" json:"version,omitempty"`
	Name            string                       `bson:"name,omitempty" json:"name,omitempty"`
	Title           string                       `bson:"title,omitempty" json:"title,omitempty"`
	DerivedFrom     []Canonical                  `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Status          string                       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental    *bool                        `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectType     []string                     `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Date            *FHIRDateTime                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher       string                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact         []ContactDetail              `bson:"contact,omitempty" json:"contact,omitempty"`
	Description     string                       `bson:"description,omitempty" json:"description,omitempty"`
	UseContext      []UsageContext               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction    []CodeableConcept            `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose         string                       `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright       string                       `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate    *FHIRDateTime                `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate  *FHIRDateTime                `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod *Period                      `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Code            []Coding                     `bson:"code,omitempty" json:"code,omitempty"`
	Item            []QuestionnaireItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (x *Questionnaire) MarshalJSON() ([]byte, error) {
	x.ResourceType = "Questionnaire"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*x)
}

// "questionnaire" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type questionnaire Questionnaire

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Questionnaire) UnmarshalJSON(data []byte) (err error) {
	x2 := questionnaire{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Questionnaire(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Questionnaire) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Questionnaire"
	} else if x.ResourceType != "Questionnaire" {
		return fmt.Errorf("Expected resourceType to be Questionnaire, instead received %s", x.ResourceType)
	}
	return nil
}

// QuestionnaireItemComponent ... // TODO Write proper comment
type QuestionnaireItemComponent struct {
	BackboneElement `bson:",inline"`
	LinkId          string                                   `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition      string                                   `bson:"definition,omitempty" json:"definition,omitempty"`
	Code            []Coding                                 `bson:"code,omitempty" json:"code,omitempty"`
	Prefix          string                                   `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text            string                                   `bson:"text,omitempty" json:"text,omitempty"`
	Type            string                                   `bson:"type,omitempty" json:"type,omitempty"`
	EnableWhen      []QuestionnaireItemEnableWhenComponent   `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	EnableBehavior  string                                   `bson:"enableBehavior,omitempty" json:"enableBehavior,omitempty"`
	Required        *bool                                    `bson:"required,omitempty" json:"required,omitempty"`
	Repeats         *bool                                    `bson:"repeats,omitempty" json:"repeats,omitempty"`
	ReadOnly        *bool                                    `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	MaxLength       *int32                                   `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	AnswerValueSet  *Canonical                               `bson:"answerValueSet,omitempty" json:"answerValueSet,omitempty"`
	AnswerOption    []QuestionnaireItemAnswerOptionComponent `bson:"answerOption,omitempty" json:"answerOption,omitempty"`
	Initial         []QuestionnaireItemInitialComponent      `bson:"initial,omitempty" json:"initial,omitempty"`
	Item            []QuestionnaireItemComponent             `bson:"item,omitempty" json:"item,omitempty"`
}

// QuestionnaireItemEnableWhenComponent ... // TODO Write proper comment
type QuestionnaireItemEnableWhenComponent struct {
	BackboneElement `bson:",inline"`
	Question        string        `bson:"question,omitempty" json:"question,omitempty"`
	Operator        string        `bson:"operator,omitempty" json:"operator,omitempty"`
	AnswerBoolean   *bool         `bson:"answerBoolean,omitempty" json:"answerBoolean,omitempty"`
	AnswerDecimal   *float64      `bson:"answerDecimal,omitempty" json:"answerDecimal,omitempty"`
	AnswerInteger   *int32        `bson:"answerInteger,omitempty" json:"answerInteger,omitempty"`
	AnswerDate      *FHIRDateTime `bson:"answerDate,omitempty" json:"answerDate,omitempty"`
	AnswerDateTime  *FHIRDateTime `bson:"answerDateTime,omitempty" json:"answerDateTime,omitempty"`
	AnswerTime      *FHIRDateTime `bson:"answerTime,omitempty" json:"answerTime,omitempty"`
	AnswerString    string        `bson:"answerString,omitempty" json:"answerString,omitempty"`
	AnswerCoding    *Coding       `bson:"answerCoding,omitempty" json:"answerCoding,omitempty"`
	AnswerQuantity  *Quantity     `bson:"answerQuantity,omitempty" json:"answerQuantity,omitempty"`
	AnswerReference *Reference    `bson:"answerReference,omitempty" json:"answerReference,omitempty"`
}

// QuestionnaireItemAnswerOptionComponent ... // TODO Write proper comment
type QuestionnaireItemAnswerOptionComponent struct {
	BackboneElement `bson:",inline"`
	ValueInteger    *int32        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueTime       *FHIRDateTime `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueCoding     *Coding       `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueReference  *Reference    `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	InitialSelected *bool         `bson:"initialSelected,omitempty" json:"initialSelected,omitempty"`
}

// QuestionnaireItemInitialComponent ... // TODO Write proper comment
type QuestionnaireItemInitialComponent struct {
	BackboneElement `bson:",inline"`
	ValueBoolean    *bool         `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal    *float64      `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger    *int32        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate       *FHIRDateTime `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueTime       *FHIRDateTime `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString     string        `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri        string        `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment *Attachment   `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding     *Coding       `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity   *Quantity     `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference  *Reference    `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
}

// QuestionnairePlus ... // TODO Write proper comment
type QuestionnairePlus struct {
	Questionnaire                     `bson:",inline"`
	QuestionnairePlusRelatedResources `bson:",inline"`
}

// QuestionnairePlusRelatedResources ... // TODO Write proper comment
type QuestionnairePlusRelatedResources struct {
	RevIncludedAppointmentResourcesReferencingSupportinginfo                *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                 *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                 *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                     *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref               *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                              *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                         *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                    *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                    *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                 *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource           *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                 *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                          *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                     *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor              *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1         *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2         *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource              *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical *[]FamilyMemberHistory        `bson:"_revIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor       *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1  *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2  *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCommunicationResourcesReferencingInstantiatescanonical       *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingInstantiatescanonical,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                      *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor              *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1         *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2         *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                              *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                            *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                     *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest                *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                       *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation    *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                         *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                         *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                              *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                                *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                              *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingInstantiatescanonical            *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingInstantiatescanonical,omitempty"`
	RevIncludedProcedureResourcesReferencingInstantiatescanonical           *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingInstantiatescanonical,omitempty"`
	RevIncludedListResourcesReferencingItem                                 *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor                *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson                *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                         *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                         *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                         *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon              *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                             *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                        *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                        *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                         *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                  *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                       *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                         *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                  *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire       *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingQuestionnaire,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo         *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                  *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1             *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2             *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDefinition                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDefinition,omitempty"`
}

// GetRevIncludedAppointmentResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *q.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEventDefinitionResourcesReferencingDependson ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if q.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *q.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingItem ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

// GetRevIncludedDocumentManifestResourcesReferencingRelatedref ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *q.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

// GetRevIncludedConsentResourcesReferencingData ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if q.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *q.RevIncludedConsentResourcesReferencingData
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedMeasureResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *q.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedDocumentReferenceResourcesReferencingRelated ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *q.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

// GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

// GetRevIncludedVerificationResultResourcesReferencingTarget ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if q.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *q.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

// GetRevIncludedContractResourcesReferencingSubject ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if q.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *q.RevIncludedContractResourcesReferencingSubject
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingRequest ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

// GetRevIncludedPaymentNoticeResourcesReferencingResponse ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *q.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedImplementationGuideResourcesReferencingResource ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if q.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *q.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

// GetRevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical() (familyMemberHistories []FamilyMemberHistory, err error) {
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical == nil {
		err = errors.New("RevIncluded familyMemberHistories not requested")
	} else {
		familyMemberHistories = *q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingInstantiatescanonical ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingInstantiatescanonical() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingPartof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

// GetRevIncludedCommunicationResourcesReferencingBasedon ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if q.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *q.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingItem ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if q.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *q.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

// GetRevIncludedLinkageResourcesReferencingSource ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if q.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *q.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedDeviceRequestResourcesReferencingPriorrequest ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

// GetRevIncludedMessageHeaderResourcesReferencingFocus ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if q.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *q.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

// GetRevIncludedImmunizationRecommendationResourcesReferencingInformation ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingEntity ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

// GetRevIncludedProvenanceResourcesReferencingTarget ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if q.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *q.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

// GetRevIncludedTaskResourcesReferencingSubject ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

// GetRevIncludedTaskResourcesReferencingFocus ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

// GetRevIncludedTaskResourcesReferencingBasedon ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if q.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *q.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedCarePlanResourcesReferencingInstantiatescanonical ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingInstantiatescanonical() (carePlans []CarePlan, err error) {
	if q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical
	}
	return
}

// GetRevIncludedProcedureResourcesReferencingInstantiatescanonical ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingInstantiatescanonical() (procedures []Procedure, err error) {
	if q.RevIncludedProcedureResourcesReferencingInstantiatescanonical == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *q.RevIncludedProcedureResourcesReferencingInstantiatescanonical
	}
	return
}

// GetRevIncludedListResourcesReferencingItem ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if q.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *q.RevIncludedListResourcesReferencingItem
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceVariableResourcesReferencingDependson ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *q.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

// GetRevIncludedObservationResourcesReferencingFocus ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if q.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *q.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedLibraryResourcesReferencingDependson ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if q.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *q.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

// GetRevIncludedCommunicationRequestResourcesReferencingBasedon ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *q.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

// GetRevIncludedBasicResourcesReferencingSubject ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if q.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *q.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedEvidenceResourcesReferencingDependson ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if q.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *q.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

// GetRevIncludedAuditEventResourcesReferencingEntity ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if q.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *q.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

// GetRevIncludedConditionResourcesReferencingEvidencedetail ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if q.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *q.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingSubject ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

// GetRevIncludedCompositionResourcesReferencingEntry ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if q.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *q.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

// GetRevIncludedDetectedIssueResourcesReferencingImplicated ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *q.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire
	}
	return
}

// GetRevIncludedQuestionnaireResponseResourcesReferencingSubject ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

// GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingSuccessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingPredecessor ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingComposedof ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2 ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

// GetRevIncludedPlanDefinitionResourcesReferencingDefinition ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDefinition() (planDefinitions []PlanDefinition, err error) {
	if q.RevIncludedPlanDefinitionResourcesReferencingDefinition == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *q.RevIncludedPlanDefinitionResourcesReferencingDefinition
	}
	return
}

// GetIncludedResources ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

// GetRevIncludedResources ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*q.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*q.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*q.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*q.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedProcedureResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedProcedureResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*q.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDefinition {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

// GetIncludedAndRevIncludedResources ... // TODO Write proper comment
func (q *QuestionnairePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if q.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*q.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *q.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*q.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *q.RevIncludedConsentResourcesReferencingData {
			rsc := (*q.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*q.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *q.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*q.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*q.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*q.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedContractResourcesReferencingSubject {
			rsc := (*q.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *q.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*q.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *q.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*q.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedFamilyMemberHistoryResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*q.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *q.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*q.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *q.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*q.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*q.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *q.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*q.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *q.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*q.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*q.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*q.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*q.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedCarePlanResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedProcedureResourcesReferencingInstantiatescanonical != nil {
		for idx := range *q.RevIncludedProcedureResourcesReferencingInstantiatescanonical {
			rsc := (*q.RevIncludedProcedureResourcesReferencingInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *q.RevIncludedListResourcesReferencingItem {
			rsc := (*q.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *q.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*q.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*q.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*q.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *q.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*q.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*q.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *q.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*q.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *q.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*q.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *q.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*q.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*q.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *q.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*q.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *q.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*q.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingQuestionnaire)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *q.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*q.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*q.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if q.RevIncludedPlanDefinitionResourcesReferencingDefinition != nil {
		for idx := range *q.RevIncludedPlanDefinitionResourcesReferencingDefinition {
			rsc := (*q.RevIncludedPlanDefinitionResourcesReferencingDefinition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
