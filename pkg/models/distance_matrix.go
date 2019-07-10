package models

//DistanceMatrixResponse ...
type DistanceMatrixResponse struct {
	DistanceMatrixResult
	CommonResponse
}

// DistanceMatrixResult ...
type DistanceMatrixResult struct {
	OriginAddresses      []string                    `json:"origin_addresses,omitempty"`
	DestinationAddresses []string                    `json:"destination_addresses,omitempty"`
	Rows                 []DistanceMatrixElementsRow `json:"rows,omitempty"`
}

// DistanceMatrixElementsRow ...
type DistanceMatrixElementsRow struct {
	Elements []*DistanceMatrixElement `json:"elements,omitempty"`
}

// DistanceMatrixElement ...
type DistanceMatrixElement struct {
	Status   string   `json:"status,omitempty"`
	Duration Duration `json:"duration,omitempty"`
	Distance Distance `json:"distance,omitempty"`
}

// Distance ...
type Distance struct {
	HumanReadable string `json:"text,omitempty"`
	Meters        int    `json:"value,omitempty"`
}

// Duration ...
type Duration struct {
	HumanReadable string `json:"text,omitempty"`
	Time          int    `json:"value,omitempty"`
}
