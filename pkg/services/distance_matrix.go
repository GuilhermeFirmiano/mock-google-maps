package services

import (
	"github.com/GuilhermeFirmiano/mock-google-maps/pkg/models"
)

// DistanceMatrixService ...
type DistanceMatrixService interface {
	Get() *models.DistanceMatrixResponse
}

//distanceMatrixService ...
type distanceMatrixService struct {
}

// NewDistanceMatrixService ...
func NewDistanceMatrixService() DistanceMatrixService {
	return &distanceMatrixService{}
}

//Get ...
func (service *distanceMatrixService) Get() *models.DistanceMatrixResponse {
	var distanceMatrix models.DistanceMatrixResult

	distanceMatrix.OriginAddresses = []string{"Vancouver, BC, Canada", "Seattle, État de Washington, États-Unis"}

	distanceMatrix.DestinationAddresses = []string{"San Francisco, Californie, États-Unis", "Victoria, BC, Canada"}

	distanceMatrix.Rows = []models.DistanceMatrixElementsRow{
		models.DistanceMatrixElementsRow{
			Elements: []*models.DistanceMatrixElement{
				&models.DistanceMatrixElement{
					Status: "OK",
					Duration: models.Duration{
						Time:          340110,
						HumanReadable: "3 jours 22 heures",
					},
					Distance: models.Distance{
						Meters:        1734542,
						HumanReadable: "1 735 km",
					},
				},
				&models.DistanceMatrixElement{
					Status: "OK",
					Duration: models.Duration{
						Time:          24487,
						HumanReadable: "6 heures 48 minutes",
					},
					Distance: models.Distance{
						Meters:        129324,
						HumanReadable: "129 km",
					},
				},
			},
		},
		models.DistanceMatrixElementsRow{
			Elements: []*models.DistanceMatrixElement{
				&models.DistanceMatrixElement{
					Status: "OK",
					Duration: models.Duration{
						Time:          288834,
						HumanReadable: "3 jours 8 heures",
					},
					Distance: models.Distance{
						Meters:        1489604,
						HumanReadable: "1 490 km",
					},
				},
				&models.DistanceMatrixElement{
					Status: "OK",
					Duration: models.Duration{
						Time:          14388,
						HumanReadable: "4 heures 0 minutes",
					},
					Distance: models.Distance{
						Meters:        135822,
						HumanReadable: "136 km",
					},
				},
			},
		},
	}

	response := models.DistanceMatrixResponse{
		DistanceMatrixResult: distanceMatrix,
		CommonResponse: models.CommonResponse{
			Status: "OK",
		},
	}

	return &response
}
