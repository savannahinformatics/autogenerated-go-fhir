package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Claim struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string                                `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	SubType              *CodeableConcept                      `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  string                                `bson:"use,omitempty" json:"use,omitempty"`
	Patient              *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period                               `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *FHIRDateTime                         `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference                            `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                            `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference                            `bson:"provider,omitempty" json:"provider,omitempty"`
	Priority             *CodeableConcept                      `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve         *CodeableConcept                      `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related              []ClaimRelatedClaimComponent          `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference                            `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference                            `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ClaimPayeeComponent                  `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral             *Reference                            `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                            `bson:"facility,omitempty" json:"facility,omitempty"`
	CareTeam             []ClaimCareTeamComponent              `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	SupportingInfo       []ClaimSupportingInformationComponent `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Diagnosis            []ClaimDiagnosisComponent             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedureComponent             `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Insurance            []ClaimInsuranceComponent             `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ClaimAccidentComponent               `bson:"accident,omitempty" json:"accident,omitempty"`
	Item                 []ClaimItemComponent                  `bson:"item,omitempty" json:"item,omitempty"`
	Total                *Quantity                             `bson:"total,omitempty" json:"total,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Claim) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Claim"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "claim" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type claim Claim

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Claim) UnmarshalJSON(data []byte) (err error) {
	x2 := claim{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Claim(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Claim) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Claim"
	} else if x.ResourceType != "Claim" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Claim, instead received %s", x.ResourceType))
	}
	return nil
}

type ClaimRelatedClaimComponent struct {
	BackboneElement `bson:",inline"`
	Claim           *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship    *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ClaimPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Party           *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}

type ClaimCareTeamComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider        *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible     *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification   *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type ClaimSupportingInformationComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate      *FHIRDateTime    `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod    *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueBoolean    *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueString     string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity   *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference  *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClaimDiagnosisComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}

type ClaimProcedureComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Date                     *FHIRDateTime     `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept  `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference        `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	Udi                      []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimInsuranceComponent struct {
	BackboneElement     `bson:",inline"`
	Sequence            *uint32     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               *bool       `bson:"focal,omitempty" json:"focal,omitempty"`
	Identifier          *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Coverage            *Reference  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string      `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimAccidentComponent struct {
	BackboneElement   `bson:",inline"`
	Date              *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}

type ClaimItemComponent struct {
	BackboneElement         `bson:",inline"`
	Sequence                *uint32                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamSequence        []uint32               `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []uint32               `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []uint32               `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	InformationSequence     []uint32               `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService        *CodeableConcept       `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier                []CodeableConcept      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *FHIRDateTime          `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept       `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address               `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference             `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Quantity              `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64               `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Quantity              `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []Reference            `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []Reference            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Detail                  []ClaimDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimDetailComponent struct {
	BackboneElement  `bson:",inline"`
	Sequence         *uint32                   `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue          *CodeableConcept          `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category         *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService *CodeableConcept          `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept         `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity         *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64                  `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi              []Reference               `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail        []ClaimSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimSubDetailComponent struct {
	BackboneElement  `bson:",inline"`
	Sequence         *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue          *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category         *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService *CodeableConcept  `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity         *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity         `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64          `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity         `bson:"net,omitempty" json:"net,omitempty"`
	Udi              []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimPlus struct {
	Claim                     `bson:",inline"`
	ClaimPlusRelatedResources `bson:",inline"`
}

type ClaimPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByCareteam                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByCareteam,omitempty"`
	IncludedOrganizationResourcesReferencedByCareteam                      *[]Organization               `bson:"_includedOrganizationResourcesReferencedByCareteam,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByCareteam                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByCareteam,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByPayee                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPayee,omitempty"`
	IncludedOrganizationResourcesReferencedByPayee                         *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPayee,omitempty"`
	IncludedPatientResourcesReferencedByPayee                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPayee                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPayee,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPayee                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerResourcesReferencedByProvider                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedOrganizationResourcesReferencedByProvider                      *[]Organization               `bson:"_includedOrganizationResourcesReferencedByProvider,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByProvider                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByInsurer                       *[]Organization               `bson:"_includedOrganizationResourcesReferencedByInsurer,omitempty"`
	IncludedDeviceResourcesReferencedByDetailudi                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByDetailudi,omitempty"`
	IncludedPractitionerResourcesReferencedByEnterer                       *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByEnterer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByEnterer                   *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByEnterer,omitempty"`
	IncludedDeviceResourcesReferencedByProcedureudi                        *[]Device                     `bson:"_includedDeviceResourcesReferencedByProcedureudi,omitempty"`
	IncludedDeviceResourcesReferencedBySubdetailudi                        *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubdetailudi,omitempty"`
	IncludedLocationResourcesReferencedByFacility                          *[]Location                   `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
	IncludedDeviceResourcesReferencedByItemudi                             *[]Device                     `bson:"_includedDeviceResourcesReferencedByItemudi,omitempty"`
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
	RevIncludedExplanationOfBenefitResourcesReferencingClaim               *[]ExplanationOfBenefit       `bson:"_revIncludedExplanationOfBenefitResourcesReferencingClaim,omitempty"`
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
	RevIncludedClaimResponseResourcesReferencingRequest                    *[]ClaimResponse              `bson:"_revIncludedClaimResponseResourcesReferencingRequest,omitempty"`
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

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCareteam() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByCareteam))
	} else if len(*c.IncludedPractitionerResourcesReferencedByCareteam) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByCareteam)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByCareteam() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByCareteam == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByCareteam))
	} else if len(*c.IncludedOrganizationResourcesReferencedByCareteam) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByCareteam)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByCareteam() (practitionerRole *PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*c.IncludedPractitionerRoleResourcesReferencedByCareteam))
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByCareteam) == 1 {
		practitionerRole = &(*c.IncludedPractitionerRoleResourcesReferencedByCareteam)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedEncounterResourcesReferencedByEncounter() (encounters []Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else {
		encounters = *c.IncludedEncounterResourcesReferencedByEncounter
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPayee() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByPayee == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByPayee))
	} else if len(*c.IncludedPractitionerResourcesReferencedByPayee) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPayee() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPayee == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByPayee))
	} else if len(*c.IncludedOrganizationResourcesReferencedByPayee) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPayee() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPayee == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPayee))
	} else if len(*c.IncludedPatientResourcesReferencedByPayee) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPayee() (practitionerRole *PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByPayee == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*c.IncludedPractitionerRoleResourcesReferencedByPayee))
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByPayee) == 1 {
		practitionerRole = &(*c.IncludedPractitionerRoleResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPayee() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPayee == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByPayee))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByPayee) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByPayee)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*c.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByProvider() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByProvider == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByProvider))
	} else if len(*c.IncludedOrganizationResourcesReferencedByProvider) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByProvider)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByProvider() (practitionerRole *PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByProvider == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*c.IncludedPractitionerRoleResourcesReferencedByProvider))
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByProvider) == 1 {
		practitionerRole = &(*c.IncludedPractitionerRoleResourcesReferencedByProvider)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedOrganizationResourceReferencedByInsurer() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByInsurer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByInsurer))
	} else if len(*c.IncludedOrganizationResourcesReferencedByInsurer) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByInsurer)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedDeviceResourcesReferencedByDetailudi() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByDetailudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByDetailudi
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerResourceReferencedByEnterer() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByEnterer))
	} else if len(*c.IncludedPractitionerResourcesReferencedByEnterer) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByEnterer)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByEnterer() (practitionerRole *PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*c.IncludedPractitionerRoleResourcesReferencedByEnterer))
	} else if len(*c.IncludedPractitionerRoleResourcesReferencedByEnterer) == 1 {
		practitionerRole = &(*c.IncludedPractitionerRoleResourcesReferencedByEnterer)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedDeviceResourcesReferencedByProcedureudi() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByProcedureudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByProcedureudi
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedDeviceResourcesReferencedBySubdetailudi() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedBySubdetailudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedBySubdetailudi
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByFacility))
	} else if len(*c.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedDeviceResourcesReferencedByItemudi() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByItemudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByItemudi
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *c.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if c.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *c.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedExplanationOfBenefitResourcesReferencingClaim() (explanationOfBenefits []ExplanationOfBenefit, err error) {
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim == nil {
		err = errors.New("RevIncluded explanationOfBenefits not requested")
	} else {
		explanationOfBenefits = *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedClaimResponseResourcesReferencingRequest() (claimResponses []ClaimResponse, err error) {
	if c.RevIncludedClaimResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded claimResponses not requested")
	} else {
		claimResponses = *c.RevIncludedClaimResponseResourcesReferencingRequest
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *ClaimPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *ClaimPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*c.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayee {
			rsc := (*c.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByProvider {
			rsc := (*c.IncludedOrganizationResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByDetailudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByDetailudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByDetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByProcedureudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByProcedureudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByProcedureudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedBySubdetailudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedBySubdetailudi {
			rsc := (*c.IncludedDeviceResourcesReferencedBySubdetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacility {
			rsc := (*c.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByItemudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByItemudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByItemudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *ClaimPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*c.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByCareteam {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPayee {
			rsc := (*c.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPayee {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByProvider {
			rsc := (*c.IncludedOrganizationResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByInsurer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByInsurer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByInsurer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByDetailudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByDetailudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByDetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByEnterer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByEnterer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByProcedureudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByProcedureudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByProcedureudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedBySubdetailudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedBySubdetailudi {
			rsc := (*c.IncludedDeviceResourcesReferencedBySubdetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *c.IncludedLocationResourcesReferencedByFacility {
			rsc := (*c.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByItemudi != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByItemudi {
			rsc := (*c.IncludedDeviceResourcesReferencedByItemudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedExplanationOfBenefitResourcesReferencingClaim != nil {
		for idx := range *c.RevIncludedExplanationOfBenefitResourcesReferencingClaim {
			rsc := (*c.RevIncludedExplanationOfBenefitResourcesReferencingClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClaimResponseResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedClaimResponseResourcesReferencingRequest {
			rsc := (*c.RevIncludedClaimResponseResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
