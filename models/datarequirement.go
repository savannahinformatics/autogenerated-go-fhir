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

type DataRequirement struct {
	Type                   string                               `bson:"type,omitempty" json:"type,omitempty"`
	Profile                []canonical                          `bson:"profile,omitempty" json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept                     `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                           `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	MustSupport            []string                             `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilterComponent `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilterComponent `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
	Limit                  *uint32                              `bson:"limit,omitempty" json:"limit,omitempty"`
	Sort                   []DataRequirementSortComponent       `bson:"sort,omitempty" json:"sort,omitempty"`
}

type DataRequirementCodeFilterComponent struct {
	BackboneElement `bson:",inline"`
	Path            string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam     string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueSet        *canonical `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Code            []Coding   `bson:"code,omitempty" json:"code,omitempty"`
}

type DataRequirementDateFilterComponent struct {
	BackboneElement `bson:",inline"`
	Path            string        `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam     string        `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueDateTime   *FHIRDateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod     *Period       `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration   *Quantity     `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}

type DataRequirementSortComponent struct {
	BackboneElement `bson:",inline"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Direction       string `bson:"direction,omitempty" json:"direction,omitempty"`
}
