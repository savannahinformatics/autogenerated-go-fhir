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

type List struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string               `bson:"status,omitempty" json:"status,omitempty"`
	Mode           string               `bson:"mode,omitempty" json:"mode,omitempty"`
	Title          string               `bson:"title,omitempty" json:"title,omitempty"`
	Code           *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter      *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date           *FHIRDateTime        `bson:"date,omitempty" json:"date,omitempty"`
	Source         *Reference           `bson:"source,omitempty" json:"source,omitempty"`
	OrderedBy      *CodeableConcept     `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Note           []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Entry          []ListEntryComponent `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason    *CodeableConcept     `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *List) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "List"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

// "list" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type list List

// UnmarshalJSON is used to properly unmarshal embedded resources (represented as interface{})
func (x *List) UnmarshalJSON(data []byte) (err error) {
	x2 := list{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = List(x2)
		return x.checkResourceType()
	}
	return
}

func (x *List) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "List"
	} else if x.ResourceType != "List" {
		return errors.New(fmt.Sprintf("Expected resourceType to be List, instead received %s", x.ResourceType))
	}
	return nil
}

type ListEntryComponent struct {
	BackboneElement `bson:",inline"`
	Flag            *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted         *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date            *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Item            *Reference       `bson:"item,omitempty" json:"item,omitempty"`
}

type ListPlus struct {
	List                     `bson:",inline"`
	ListPlusRelatedResources `bson:",inline"`
}

type ListPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                             *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                           *[]Location                   `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedDeviceResourcesReferencedBySource                              *[]Device                     `bson:"_includedDeviceResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                             *[]Patient                    `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedPractitionerRoleResourcesReferencedBySource                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedBySource,omitempty"`
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

func (l *ListPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if l.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*l.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*l.IncludedGroupResourcesReferencedBySubject))
	} else if len(*l.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*l.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if l.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*l.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*l.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedBySubject))
	} else if len(*l.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if l.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*l.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*l.IncludedLocationResourcesReferencedBySubject))
	} else if len(*l.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*l.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedByPatient))
	} else if len(*l.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if l.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*l.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*l.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*l.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*l.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedDeviceResourceReferencedBySource() (device *Device, err error) {
	if l.IncludedDeviceResourcesReferencedBySource == nil {
		err = errors.New("Included devices not requested")
	} else if len(*l.IncludedDeviceResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*l.IncludedDeviceResourcesReferencedBySource))
	} else if len(*l.IncludedDeviceResourcesReferencedBySource) == 1 {
		device = &(*l.IncludedDeviceResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if l.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*l.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*l.IncludedPatientResourcesReferencedBySource))
	} else if len(*l.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*l.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedBySource() (practitionerRole *PractitionerRole, err error) {
	if l.IncludedPractitionerRoleResourcesReferencedBySource == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*l.IncludedPractitionerRoleResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*l.IncludedPractitionerRoleResourcesReferencedBySource))
	} else if len(*l.IncludedPractitionerRoleResourcesReferencedBySource) == 1 {
		practitionerRole = &(*l.IncludedPractitionerRoleResourcesReferencedBySource)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if l.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*l.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*l.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*l.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*l.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *l.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if l.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *l.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if l.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *l.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *l.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if l.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *l.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if l.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *l.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if l.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *l.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if l.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *l.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if l.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *l.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if l.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *l.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *l.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if l.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *l.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *l.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if l.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *l.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if l.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *l.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *l.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if l.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *l.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if l.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *l.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (l *ListPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (l *ListPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedGroupResourcesReferencedBySubject {
			rsc := (*l.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*l.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedPatientResourcesReferencedBySubject {
			rsc := (*l.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedLocationResourcesReferencedBySubject {
			rsc := (*l.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *l.IncludedPatientResourcesReferencedByPatient {
			rsc := (*l.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*l.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedDeviceResourcesReferencedBySource != nil {
		for idx := range *l.IncludedDeviceResourcesReferencedBySource {
			rsc := (*l.IncludedDeviceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPatientResourcesReferencedBySource {
			rsc := (*l.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*l.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *l.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*l.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *ListPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*l.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*l.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*l.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*l.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*l.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*l.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*l.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *l.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*l.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (l *ListPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedGroupResourcesReferencedBySubject {
			rsc := (*l.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedDeviceResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedDeviceResourcesReferencedBySubject {
			rsc := (*l.IncludedDeviceResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedPatientResourcesReferencedBySubject {
			rsc := (*l.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedLocationResourcesReferencedBySubject != nil {
		for idx := range *l.IncludedLocationResourcesReferencedBySubject {
			rsc := (*l.IncludedLocationResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *l.IncludedPatientResourcesReferencedByPatient {
			rsc := (*l.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*l.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedDeviceResourcesReferencedBySource != nil {
		for idx := range *l.IncludedDeviceResourcesReferencedBySource {
			rsc := (*l.IncludedDeviceResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPatientResourcesReferencedBySource {
			rsc := (*l.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *l.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*l.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *l.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*l.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*l.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*l.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *l.RevIncludedConsentResourcesReferencingData {
			rsc := (*l.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*l.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *l.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*l.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*l.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*l.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedContractResourcesReferencingSubject {
			rsc := (*l.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*l.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *l.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*l.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*l.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *l.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*l.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *l.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*l.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*l.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *l.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*l.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*l.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*l.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*l.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*l.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *l.RevIncludedListResourcesReferencingItem {
			rsc := (*l.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *l.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*l.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*l.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*l.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *l.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*l.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*l.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *l.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*l.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *l.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*l.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *l.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*l.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*l.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *l.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*l.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*l.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*l.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*l.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*l.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
