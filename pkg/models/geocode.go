package models

//GeocodeResponse ...
type GeocodeResponse struct {
	Results []GeocodeResult `json:"results"`
	CommonResponse
}

// GeocodeResult is a single geocoded address
type GeocodeResult struct {
	AddressComponents []AddressComponent `json:"address_components,omitempty"`
	FormattedAddress  string             `json:"formatted_address,omitempty"`
	Geometry          AddressGeometry    `json:"geometry,omitempty"`
	Types             []string           `json:"types,omitempty"`
	PlaceID           string             `json:"place_id,omitempty"`

	// PartialMatch indicates that the geocoder did not return an exact match for
	// the original request, though it was able to match part of the requested address.
	// You may wish to examine the original request for misspellings and/or an incomplete address.
	// Partial matches most often occur for street addresses that do not exist within the
	// locality you pass in the request.
	// Partial matches may also be returned when a request matches two or more locations in
	// the same locality. For example, "21 Henr St, Bristol, UK" will return a partial match
	// for both Henry Street and Henrietta Street.
	// Note that if a request includes a misspelled address component, the geocoding service may
	// suggest an alternative address.
	// Suggestions triggered in this way will also be marked as a partial match.
	PartialMatch bool `json:"partial_match,omitempty"`
}

// AddressPlusCode (see https://en.wikipedia.org/wiki/Open_Location_Code and https://plus.codes/)
// is an encoded location reference, derived from latitude and longitude coordinates,
// that represents an area: 1/8000th of a degree by 1/8000th of a degree (about 14m x 14m at the equator)
// or smaller.
//
// Plus codes can be used as a replacement for street addresses in places where they do not exist
// (where buildings are not numbered or streets are not named).
// The plus code is formatted as a global code and a compound code:
// Typically, both the global code and compound code are returned.
// However, if the result is in a remote location (for example, an ocean or desert)
// only the global code may be returned.
type AddressPlusCode struct {
	// GlobalCode is a 4 character area code and 6 character or longer local code (849VCWC8+R9).
	GlobalCode string `json:"global_code,omitempty"`
	// CompoundCode is a 6 character or longer local code with an explicit location (CWC8+R9, Mountain View, CA, USA).
	CompoundCode string `json:"compound_code,omitempty"`
}

// AddressComponent is a part of an address
type AddressComponent struct {
	LongName  string   `json:"long_name,omitempty"`
	ShortName string   `json:"short_name,omitempty"`
	Types     []string `json:"types,omitempty"`
}

// AddressGeometry is the location of a an address
type AddressGeometry struct {
	Location     LatLng       `json:"location,omitempty"`
	LocationType string       `json:"location_type,omitempty"`
	Bounds       LatLngBounds `json:"bounds,omitempty"`
	Viewport     LatLngBounds `json:"viewport,omitempty"`
	Types        []string     `json:"types,omitempty"`
}
