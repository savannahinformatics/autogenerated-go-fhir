package fhir

// Canonical is a URI that refers to a resource by its Canonical URL
// (resources with a url property). The Canonical type differs from a uri in
// that it has special meaning in FHIR, and in that it may have a version
// appended, separated by a vertical bar (|). Note that the type Canonical
// is not used for the actual Canonical URLs that are the target of these
// references, but for the URIs that refer to them, and may have the version
// suffix in them. Like other URIs, elements of type Canonical may also have
// #fragment references	xs:anyURI	A JSON string - a Canonical URL
type Canonical string
