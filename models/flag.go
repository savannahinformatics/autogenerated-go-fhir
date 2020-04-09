// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Flag struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string            `bson:"status,omitempty" json:"status,omitempty"`
	Category       []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code           *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Period         *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Encounter      *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author         *Reference        `bson:"author,omitempty" json:"author,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *Flag) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Flag"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "flag" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type flag Flag

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *Flag) UnmarshalJSON(data []byte) (err error) {
	x2 := flag{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Flag(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Flag) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Flag"
	} else if x.ResourceType != "Flag" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Flag, instead received %s", x.ResourceType))
	}
	return nil
}

type FlagPlus struct {
	Flag                     `bson:",inline"`
	FlagPlusRelatedResources `bson:",inline"`
}

type FlagPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedBySubject                       *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedBySubject                       *[]Organization               `bson:"_includedOrganizationResourcesReferencedBySubject,omitempty"`
	IncludedMedicationResourcesReferencedBySubject                         *[]Medication                 `bson:"_includedMedicationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPlanDefinitionResourcesReferencedBySubject                     *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedProcedureResourcesReferencedBySubject                          *[]Procedure                  `bson:"_includedProcedureResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                           *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                              *[]Device                     `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                             *[]Patient                    `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByAuthor                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByAuthor,omitempty"`
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

func (f *FlagPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySubject() (practitioner *Practitioner, err error) {
	if f.IncludedPractitionerResourcesReferencedBySubject == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedPractitionerResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedPractitionerResourcesReferencedBySubject))
	} else if len(*f.IncludedPractitionerResourcesReferencedBySubject) == 1 {
		practitioner = &(*f.IncludedPractitionerResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if f.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*f.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*f.IncludedGroupResourcesReferencedBySubject))
	} else if len(*f.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*f.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySubject() (organization *Organization, err error) {
	if f.IncludedOrganizationResourcesReferencedBySubject == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedOrganizationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedOrganizationResourcesReferencedBySubject))
	} else if len(*f.IncludedOrganizationResourcesReferencedBySubject) == 1 {
		organization = &(*f.IncludedOrganizationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedMedicationResourceReferencedBySubject() (medication *Medication, err error) {
	if f.IncludedMedicationResourcesReferencedBySubject == nil {
		err = errors.New("Included medications not requested")
	} else if len(*f.IncludedMedicationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*f.IncludedMedicationResourcesReferencedBySubject))
	} else if len(*f.IncludedMedicationResourcesReferencedBySubject) == 1 {
		medication = &(*f.IncludedMedicationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedBySubject))
	} else if len(*f.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPlanDefinitionResourceReferencedBySubject() (planDefinition *PlanDefinition, err error) {
	if f.IncludedPlanDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included plandefinitions not requested")
	} else if len(*f.IncludedPlanDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 planDefinition, but found %d", len(*f.IncludedPlanDefinitionResourcesReferencedBySubject))
	} else if len(*f.IncludedPlanDefinitionResourcesReferencedBySubject) == 1 {
		planDefinition = &(*f.IncludedPlanDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedProcedureResourceReferencedBySubject() (procedure *Procedure, err error) {
	if f.IncludedProcedureResourcesReferencedBySubject == nil {
		err = errors.New("Included procedures not requested")
	} else if len(*f.IncludedProcedureResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 procedure, but found %d", len(*f.IncludedProcedureResourcesReferencedBySubject))
	} else if len(*f.IncludedProcedureResourcesReferencedBySubject) == 1 {
		procedure = &(*f.IncludedProcedureResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if f.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*f.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*f.IncludedLocationResourcesReferencedBySubject))
	} else if len(*f.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*f.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByPatient))
	} else if len(*f.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if f.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*f.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*f.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*f.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*f.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if f.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*f.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*f.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*f.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*f.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedDeviceResourceReferencedByAuthor() (device *Device, err error) {
	if f.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else if len(*f.IncludedDeviceResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*f.IncludedDeviceResourcesReferencedByAuthor))
	} else if len(*f.IncludedDeviceResourcesReferencedByAuthor) == 1 {
		device = &(*f.IncludedDeviceResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if f.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*f.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*f.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*f.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*f.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByAuthor() (practitionerRole *PractitionerRole, err error) {
	if f.IncludedPractitionerRoleResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*f.IncludedPractitionerRoleResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*f.IncludedPractitionerRoleResourcesReferencedByAuthor))
	} else if len(*f.IncludedPractitionerRoleResourcesReferencedByAuthor) == 1 {
		practitionerRole = &(*f.IncludedPractitionerRoleResourcesReferencedByAuthor)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if f.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*f.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*f.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*f.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*f.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *f.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if f.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *f.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *f.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if f.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *f.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *f.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *f.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if f.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *f.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if f.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *f.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *f.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if f.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *f.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if f.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *f.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if f.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *f.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if f.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *f.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if f.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *f.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if f.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *f.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if f.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *f.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if f.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *f.RevIncludedListResourcesReferencingItem
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *f.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if f.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *f.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if f.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *f.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *f.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if f.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *f.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if f.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *f.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if f.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *f.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if f.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *f.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if f.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *f.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *f.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (f *FlagPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (f *FlagPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*f.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedGroupResourcesReferencedBySubject {
			rsc := (*f.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*f.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*f.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPatientResourcesReferencedBySubject {
			rsc := (*f.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedProcedureResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedProcedureResourcesReferencedBySubject {
			rsc := (*f.IncludedProcedureResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedLocationResourcesReferencedBySubject {
			rsc := (*f.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*f.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*f.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*f.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *f.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*f.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingData {
			rsc := (*f.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*f.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*f.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*f.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*f.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*f.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (f *FlagPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if f.IncludedPractitionerResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedBySubject {
			rsc := (*f.IncludedPractitionerResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedGroupResourcesReferencedBySubject {
			rsc := (*f.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedBySubject {
			rsc := (*f.IncludedOrganizationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedMedicationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedMedicationResourcesReferencedBySubject {
			rsc := (*f.IncludedMedicationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPatientResourcesReferencedBySubject {
			rsc := (*f.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*f.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedProcedureResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedProcedureResourcesReferencedBySubject {
			rsc := (*f.IncludedProcedureResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *f.IncludedLocationResourcesReferencedBySubject {
			rsc := (*f.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByPatient {
			rsc := (*f.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*f.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedDeviceResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedDeviceResourcesReferencedByAuthor {
			rsc := (*f.IncludedDeviceResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*f.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *f.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*f.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *f.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*f.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*f.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *f.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*f.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *f.RevIncludedConsentResourcesReferencingData {
			rsc := (*f.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*f.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *f.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*f.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*f.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*f.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedContractResourcesReferencingSubject {
			rsc := (*f.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *f.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*f.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *f.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*f.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*f.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *f.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*f.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *f.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*f.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*f.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *f.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*f.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *f.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*f.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*f.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*f.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*f.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *f.RevIncludedListResourcesReferencingItem {
			rsc := (*f.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *f.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*f.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*f.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*f.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *f.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*f.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*f.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *f.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*f.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *f.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*f.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *f.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*f.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*f.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *f.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*f.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *f.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*f.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *f.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*f.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*f.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*f.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
