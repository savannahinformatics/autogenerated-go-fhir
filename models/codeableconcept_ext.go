package fhir

// CodeableConcepts ... // TODO Write proper comment
type CodeableConcepts []CodeableConcept

// AnyMatchesCode ... // TODO Write proper comment
func (slice CodeableConcepts) AnyMatchesCode(system string, code string) bool {
	for _, concept := range slice {
		if concept.MatchesCode(system, code) {
			return true
		}
	}
	return false
}

// AnyMatchesAnyCode ... // TODO Write proper comment
func (slice CodeableConcepts) AnyMatchesAnyCode(codings []Coding) bool {
	for _, coding := range codings {
		if slice.AnyMatchesCode(coding.System, coding.Code) {
			return true
		}
	}
	return false
}

// MatchesCode ... // TODO Write proper comment
func (c *CodeableConcept) MatchesCode(system string, code string) bool {
	for _, coding := range c.Coding {
		if coding.System == system && coding.Code == code {
			return true
		}
	}
	return false
}
