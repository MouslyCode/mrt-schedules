package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MouslyCode/mrt-schedules/common/client"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationResponse, err error) {
	// layer Service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"
	// url := "https://www.jakartamrt.co.id/"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	// Response
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}
