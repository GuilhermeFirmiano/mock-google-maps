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
	PartialMatch      bool               `json:"partial_match,omitempty"`
}

// AddressPlusCode ...
type AddressPlusCode struct {
	GlobalCode   string `json:"global_code,omitempty"`
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
