package services

import "github.com/GuilhermeFirmiano/mock-google-maps/pkg/models"

// GeocodeSerivce ...
type GeocodeSerivce interface {
	Get() *models.GeocodeResponse
}

//geocodeService ...
type geocodeService struct {
}

// NewGeocodeSerivce ...
func NewGeocodeSerivce() GeocodeSerivce {
	return &geocodeService{}
}

//Get ...
func (service *geocodeService) Get() *models.GeocodeResponse {
	var geocode models.GeocodeResult

	geocode.AddressComponents = []models.AddressComponent{
		models.AddressComponent{
			LongName:  "1600",
			ShortName: "1600",
			Types:     []string{"street_number"},
		},
		models.AddressComponent{
			LongName:  "Amphitheatre Pkwy",
			ShortName: "Amphitheatre Pkwy",
			Types:     []string{"route"},
		},
		models.AddressComponent{
			LongName:  "Mountain View",
			ShortName: "Mountain View",
			Types:     []string{"locality", "political"},
		},
		models.AddressComponent{
			LongName:  "Santa Clara County",
			ShortName: "Santa Clara County",
			Types:     []string{"administrative_area_level_2", "political"},
		},
		models.AddressComponent{
			LongName:  "California",
			ShortName: "CA",
			Types:     []string{"administrative_area_level_1", "political"},
		},
		models.AddressComponent{
			LongName:  "United States",
			ShortName: "US",
			Types:     []string{"country", "political"},
		},
		models.AddressComponent{
			LongName:  "94043",
			ShortName: "94043",
			Types:     []string{"postal_code"},
		},
	}

	geocode.FormattedAddress = "1600 Amphitheatre Parkway, Mountain View, CA 94043, USA"

	geocode.Geometry = models.AddressGeometry{
		Location: models.LatLng{
			Lat: 37.4224764,
			Lng: -122.0842499,
		},
		LocationType: "ROOFTOP",
		Viewport: models.LatLngBounds{
			NorthEast: models.LatLng{
				Lat: 37.4238253802915,
				Lng: -122.0829009197085,
			},
			SouthWest: models.LatLng{
				Lat: 37.4211274197085,
				Lng: -122.0855988802915,
			},
		},
	}

	geocode.PlaceID = "ChIJ2eUgeAK6j4ARbn5u_wAGqWA"

	geocode.Types = []string{"street_address"}

	response := models.GeocodeResponse{
		Results: []models.GeocodeResult{geocode},
		CommonResponse: models.CommonResponse{
			Status: "OK",
		},
	}

	return &response
}
