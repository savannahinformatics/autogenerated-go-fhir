package fhir

import (
	"encoding/json"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

// GetResourceID returns the ID of valid resources and a boolean that indicates if it succeeded
func GetResourceID(resource interface{}) (ID string, ok bool) {
	if value := reflect.ValueOf(resource).Elem().FieldByName("Id"); value.CanInterface() {
		ID, ok = value.Interface().(string)
	}
	return
}

// GetResourceMeta return the Meta of valid resources and a boolean that indicates if it succeeded
func GetResourceMeta(resource interface{}) (meta *Meta, ok bool) {
	if value := reflect.ValueOf(resource).Elem().FieldByName("Meta"); value.CanInterface() {
		meta, ok = value.Interface().(*Meta)
	}
	return
}

//MapToResource ensures that when FHIR JSON is unmarshalled, types that are
// interface{} just get unmarshaled to map[string]interface{}.
// This function converts that unmarshaled map to a specific resource type.
func MapToResource(resourceMap interface{}, asPointer bool) interface{} {
	b, _ := json.Marshal(&resourceMap)
	m := resourceMap.(map[string]interface{})
	t := m["resourceType"]
	switch t {
	case "Account":
		x := Account{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ActivityDefinition":
		x := ActivityDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "AdverseEvent":
		x := AdverseEvent{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "AllergyIntolerance":
		x := AllergyIntolerance{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Appointment":
		x := Appointment{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "AppointmentResponse":
		x := AppointmentResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "AuditEvent":
		x := AuditEvent{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Basic":
		x := Basic{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Binary":
		x := Binary{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "BiologicallyDerivedProduct":
		x := BiologicallyDerivedProduct{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "BodyStructure":
		x := BodyStructure{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Bundle":
		x := Bundle{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CapabilityStatement":
		x := CapabilityStatement{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CarePlan":
		x := CarePlan{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CareTeam":
		x := CareTeam{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CatalogEntry":
		x := CatalogEntry{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ChargeItem":
		x := ChargeItem{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ChargeItemDefinition":
		x := ChargeItemDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Claim":
		x := Claim{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ClaimResponse":
		x := ClaimResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ClinicalImpression":
		x := ClinicalImpression{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CodeSystem":
		x := CodeSystem{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Communication":
		x := Communication{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CommunicationRequest":
		x := CommunicationRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CompartmentDefinition":
		x := CompartmentDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Composition":
		x := Composition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ConceptMap":
		x := ConceptMap{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Condition":
		x := Condition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Consent":
		x := Consent{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Contract":
		x := Contract{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Coverage":
		x := Coverage{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CoverageEligibilityRequest":
		x := CoverageEligibilityRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "CoverageEligibilityResponse":
		x := CoverageEligibilityResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DetectedIssue":
		x := DetectedIssue{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Device":
		x := Device{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceDefinition":
		x := DeviceDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceMetric":
		x := DeviceMetric{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceRequest":
		x := DeviceRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceUseStatement":
		x := DeviceUseStatement{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DiagnosticReport":
		x := DiagnosticReport{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DocumentManifest":
		x := DocumentManifest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "DocumentReference":
		x := DocumentReference{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EffectEvidenceSynthesis":
		x := EffectEvidenceSynthesis{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Encounter":
		x := Encounter{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Endpoint":
		x := Endpoint{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EnrollmentRequest":
		x := EnrollmentRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EnrollmentResponse":
		x := EnrollmentResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EpisodeOfCare":
		x := EpisodeOfCare{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EventDefinition":
		x := EventDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Evidence":
		x := Evidence{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "EvidenceVariable":
		x := EvidenceVariable{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ExampleScenario":
		x := ExampleScenario{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ExplanationOfBenefit":
		x := ExplanationOfBenefit{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "FamilyMemberHistory":
		x := FamilyMemberHistory{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Flag":
		x := Flag{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Goal":
		x := Goal{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "GraphDefinition":
		x := GraphDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Group":
		x := Group{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "GuidanceResponse":
		x := GuidanceResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "HealthcareService":
		x := HealthcareService{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ImagingStudy":
		x := ImagingStudy{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Immunization":
		x := Immunization{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ImmunizationEvaluation":
		x := ImmunizationEvaluation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ImmunizationRecommendation":
		x := ImmunizationRecommendation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ImplementationGuide":
		x := ImplementationGuide{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "InsurancePlan":
		x := InsurancePlan{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Invoice":
		x := Invoice{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Library":
		x := Library{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Linkage":
		x := Linkage{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "List":
		x := List{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Location":
		x := Location{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Measure":
		x := Measure{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MeasureReport":
		x := MeasureReport{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Media":
		x := Media{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Medication":
		x := Medication{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationAdministration":
		x := MedicationAdministration{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationDispense":
		x := MedicationDispense{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationKnowledge":
		x := MedicationKnowledge{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationRequest":
		x := MedicationRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationStatement":
		x := MedicationStatement{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProduct":
		x := MedicinalProduct{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductAuthorization":
		x := MedicinalProductAuthorization{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductContraindication":
		x := MedicinalProductContraindication{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductIndication":
		x := MedicinalProductIndication{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductIngredient":
		x := MedicinalProductIngredient{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductInteraction":
		x := MedicinalProductInteraction{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductManufactured":
		x := MedicinalProductManufactured{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductPackaged":
		x := MedicinalProductPackaged{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductPharmaceutical":
		x := MedicinalProductPharmaceutical{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductUndesirableEffect":
		x := MedicinalProductUndesirableEffect{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MessageDefinition":
		x := MessageDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MessageHeader":
		x := MessageHeader{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "MolecularSequence":
		x := MolecularSequence{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "NamingSystem":
		x := NamingSystem{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "NutritionOrder":
		x := NutritionOrder{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Observation":
		x := Observation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ObservationDefinition":
		x := ObservationDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "OperationDefinition":
		x := OperationDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "OperationOutcome":
		x := OperationOutcome{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Organization":
		x := Organization{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "OrganizationAffiliation":
		x := OrganizationAffiliation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Parameters":
		x := Parameters{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Patient":
		x := Patient{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "PaymentNotice":
		x := PaymentNotice{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "PaymentReconciliation":
		x := PaymentReconciliation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Person":
		x := Person{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "PlanDefinition":
		x := PlanDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Practitioner":
		x := Practitioner{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "PractitionerRole":
		x := PractitionerRole{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Procedure":
		x := Procedure{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Provenance":
		x := Provenance{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Questionnaire":
		x := Questionnaire{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "QuestionnaireResponse":
		x := QuestionnaireResponse{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "RelatedPerson":
		x := RelatedPerson{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "RequestGroup":
		x := RequestGroup{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchDefinition":
		x := ResearchDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchElementDefinition":
		x := ResearchElementDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchStudy":
		x := ResearchStudy{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchSubject":
		x := ResearchSubject{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "RiskAssessment":
		x := RiskAssessment{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "RiskEvidenceSynthesis":
		x := RiskEvidenceSynthesis{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Schedule":
		x := Schedule{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SearchParameter":
		x := SearchParameter{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ServiceRequest":
		x := ServiceRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Slot":
		x := Slot{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Specimen":
		x := Specimen{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SpecimenDefinition":
		x := SpecimenDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "StructureDefinition":
		x := StructureDefinition{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "StructureMap":
		x := StructureMap{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Subscription":
		x := Subscription{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Substance":
		x := Substance{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceNucleicAcid":
		x := SubstanceNucleicAcid{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstancePolymer":
		x := SubstancePolymer{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceProtein":
		x := SubstanceProtein{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceReferenceInformation":
		x := SubstanceReferenceInformation{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceSourceMaterial":
		x := SubstanceSourceMaterial{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceSpecification":
		x := SubstanceSpecification{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SupplyDelivery":
		x := SupplyDelivery{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "SupplyRequest":
		x := SupplyRequest{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "Task":
		x := Task{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "TerminologyCapabilities":
		x := TerminologyCapabilities{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "TestReport":
		x := TestReport{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "TestScript":
		x := TestScript{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "ValueSet":
		x := ValueSet{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "VerificationResult":
		x := VerificationResult{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x
	case "VisionPrescription":
		x := VisionPrescription{}
		json.Unmarshal(b, &x)
		if asPointer {
			return &x
		}
		return x

	}
	return nil
}

// BSONMapToResource ensures that when bson objects from the database get
// unmarshaled, types that are interface{} just get unmarshaled to
// map[string]interface{}. This function converts that unmarshaled
// bson.M map to a specific resource type.
func BSONMapToResource(bsonMap bson.M, asPointer bool) interface{} {
	data, _ := bson.Marshal(bsonMap)
	raw := bson.Raw{Kind: 3, Data: data}
	t := bsonMap["resourceType"]
	switch t {
	case "Account":
		x := Account{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ActivityDefinition":
		x := ActivityDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "AdverseEvent":
		x := AdverseEvent{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "AllergyIntolerance":
		x := AllergyIntolerance{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Appointment":
		x := Appointment{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "AppointmentResponse":
		x := AppointmentResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "AuditEvent":
		x := AuditEvent{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Basic":
		x := Basic{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Binary":
		x := Binary{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "BiologicallyDerivedProduct":
		x := BiologicallyDerivedProduct{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "BodyStructure":
		x := BodyStructure{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Bundle":
		x := Bundle{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CapabilityStatement":
		x := CapabilityStatement{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CarePlan":
		x := CarePlan{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CareTeam":
		x := CareTeam{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CatalogEntry":
		x := CatalogEntry{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ChargeItem":
		x := ChargeItem{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ChargeItemDefinition":
		x := ChargeItemDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Claim":
		x := Claim{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ClaimResponse":
		x := ClaimResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ClinicalImpression":
		x := ClinicalImpression{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CodeSystem":
		x := CodeSystem{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Communication":
		x := Communication{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CommunicationRequest":
		x := CommunicationRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CompartmentDefinition":
		x := CompartmentDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Composition":
		x := Composition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ConceptMap":
		x := ConceptMap{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Condition":
		x := Condition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Consent":
		x := Consent{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Contract":
		x := Contract{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Coverage":
		x := Coverage{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CoverageEligibilityRequest":
		x := CoverageEligibilityRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "CoverageEligibilityResponse":
		x := CoverageEligibilityResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DetectedIssue":
		x := DetectedIssue{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Device":
		x := Device{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceDefinition":
		x := DeviceDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceMetric":
		x := DeviceMetric{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceRequest":
		x := DeviceRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DeviceUseStatement":
		x := DeviceUseStatement{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DiagnosticReport":
		x := DiagnosticReport{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DocumentManifest":
		x := DocumentManifest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "DocumentReference":
		x := DocumentReference{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EffectEvidenceSynthesis":
		x := EffectEvidenceSynthesis{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Encounter":
		x := Encounter{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Endpoint":
		x := Endpoint{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EnrollmentRequest":
		x := EnrollmentRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EnrollmentResponse":
		x := EnrollmentResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EpisodeOfCare":
		x := EpisodeOfCare{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EventDefinition":
		x := EventDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Evidence":
		x := Evidence{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "EvidenceVariable":
		x := EvidenceVariable{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ExampleScenario":
		x := ExampleScenario{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ExplanationOfBenefit":
		x := ExplanationOfBenefit{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "FamilyMemberHistory":
		x := FamilyMemberHistory{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Flag":
		x := Flag{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Goal":
		x := Goal{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "GraphDefinition":
		x := GraphDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Group":
		x := Group{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "GuidanceResponse":
		x := GuidanceResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "HealthcareService":
		x := HealthcareService{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ImagingStudy":
		x := ImagingStudy{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Immunization":
		x := Immunization{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ImmunizationEvaluation":
		x := ImmunizationEvaluation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ImmunizationRecommendation":
		x := ImmunizationRecommendation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ImplementationGuide":
		x := ImplementationGuide{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "InsurancePlan":
		x := InsurancePlan{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Invoice":
		x := Invoice{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Library":
		x := Library{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Linkage":
		x := Linkage{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "List":
		x := List{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Location":
		x := Location{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Measure":
		x := Measure{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MeasureReport":
		x := MeasureReport{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Media":
		x := Media{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Medication":
		x := Medication{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationAdministration":
		x := MedicationAdministration{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationDispense":
		x := MedicationDispense{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationKnowledge":
		x := MedicationKnowledge{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationRequest":
		x := MedicationRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicationStatement":
		x := MedicationStatement{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProduct":
		x := MedicinalProduct{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductAuthorization":
		x := MedicinalProductAuthorization{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductContraindication":
		x := MedicinalProductContraindication{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductIndication":
		x := MedicinalProductIndication{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductIngredient":
		x := MedicinalProductIngredient{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductInteraction":
		x := MedicinalProductInteraction{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductManufactured":
		x := MedicinalProductManufactured{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductPackaged":
		x := MedicinalProductPackaged{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductPharmaceutical":
		x := MedicinalProductPharmaceutical{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MedicinalProductUndesirableEffect":
		x := MedicinalProductUndesirableEffect{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MessageDefinition":
		x := MessageDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MessageHeader":
		x := MessageHeader{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "MolecularSequence":
		x := MolecularSequence{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "NamingSystem":
		x := NamingSystem{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "NutritionOrder":
		x := NutritionOrder{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Observation":
		x := Observation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ObservationDefinition":
		x := ObservationDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "OperationDefinition":
		x := OperationDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "OperationOutcome":
		x := OperationOutcome{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Organization":
		x := Organization{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "OrganizationAffiliation":
		x := OrganizationAffiliation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Parameters":
		x := Parameters{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Patient":
		x := Patient{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "PaymentNotice":
		x := PaymentNotice{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "PaymentReconciliation":
		x := PaymentReconciliation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Person":
		x := Person{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "PlanDefinition":
		x := PlanDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Practitioner":
		x := Practitioner{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "PractitionerRole":
		x := PractitionerRole{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Procedure":
		x := Procedure{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Provenance":
		x := Provenance{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Questionnaire":
		x := Questionnaire{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "QuestionnaireResponse":
		x := QuestionnaireResponse{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "RelatedPerson":
		x := RelatedPerson{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "RequestGroup":
		x := RequestGroup{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchDefinition":
		x := ResearchDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchElementDefinition":
		x := ResearchElementDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchStudy":
		x := ResearchStudy{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ResearchSubject":
		x := ResearchSubject{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "RiskAssessment":
		x := RiskAssessment{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "RiskEvidenceSynthesis":
		x := RiskEvidenceSynthesis{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Schedule":
		x := Schedule{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SearchParameter":
		x := SearchParameter{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ServiceRequest":
		x := ServiceRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Slot":
		x := Slot{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Specimen":
		x := Specimen{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SpecimenDefinition":
		x := SpecimenDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "StructureDefinition":
		x := StructureDefinition{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "StructureMap":
		x := StructureMap{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Subscription":
		x := Subscription{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Substance":
		x := Substance{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceNucleicAcid":
		x := SubstanceNucleicAcid{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstancePolymer":
		x := SubstancePolymer{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceProtein":
		x := SubstanceProtein{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceReferenceInformation":
		x := SubstanceReferenceInformation{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceSourceMaterial":
		x := SubstanceSourceMaterial{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SubstanceSpecification":
		x := SubstanceSpecification{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SupplyDelivery":
		x := SupplyDelivery{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "SupplyRequest":
		x := SupplyRequest{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "Task":
		x := Task{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "TerminologyCapabilities":
		x := TerminologyCapabilities{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "TestReport":
		x := TestReport{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "TestScript":
		x := TestScript{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "ValueSet":
		x := ValueSet{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "VerificationResult":
		x := VerificationResult{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	case "VisionPrescription":
		x := VisionPrescription{}
		raw.Unmarshal(&x)
		if asPointer {
			return &x
		}
		return x
	}
	return nil
}
