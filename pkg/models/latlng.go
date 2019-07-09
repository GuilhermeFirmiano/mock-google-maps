package models

// LatLng represents a location on the Earth.
type LatLng struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

// LatLngBounds represents a bounded square area on the Earth.
type LatLngBounds struct {
	NorthEast LatLng `json:"northeast,omitempty"`
	SouthWest LatLng `json:"southwest,omitempty"`
}
