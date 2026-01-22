package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MouslyCode/mrt-jakarta-api/common/client"
	"github.com/MouslyCode/mrt-jakarta-api/common/helper"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
	CheckScheduleByStations(id string) (response []StationScheduleResponse, err error)
	CheckEstimateByStations(id string) (response []StationEstimateResponse, err error)
	CheckFacilityByStations(id string) (response []StationFacilityResponse, err error)
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

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	// Response
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}

func (s *service) CheckScheduleByStations(id string) (response []StationScheduleResponse, err error) {
	// Layer Service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	// Response
	for _, station := range stations {
		if station.Id != id {
			continue
		}

		schedules := []ScheduleResponse{
			{
				BundaranHIRegular: helper.SplitSchedule(station.StationScheduleHI),
				LebakBulusRegular: helper.SplitSchedule(station.StationScheduleLB),
			},
		}

		response = append(response, StationScheduleResponse{
			StationName: station.Name,
			Schedules:   schedules,
		})

		return

	}

	err = errors.New("Station Not Found")

	return

}

func (s *service) CheckEstimateByStations(id string) (response []StationEstimateResponse, err error) {
	// Layer Service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	// Hit URL & Response
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	stationNameById := make(map[string]string, len(stations))
	for _, s := range stations {
		stationNameById[s.Id] = s.Name
	}

	for _, station := range stations {

		if station.Id != id {
			continue
		}

		var estimates []EstimateResponse
		for _, est := range station.StationEstimate {
			stationName := stationNameById[est.StationId]
			estimates = append(estimates, EstimateResponse{
				StationName: stationName,
				Fare:        est.Fare,
				Time:        est.Time,
			})
		}

		response = append(response, StationEstimateResponse{
			StationName: station.Name,
			Estimates:   estimates,
		})

		return
	}

	err = errors.New("Station Not Found")

	return
}

func (s *service) CheckFacilityByStations(id string) (response []StationFacilityResponse, err error) {
	// Layer Service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	for _, station := range stations {

		if station.Id != id {
			continue
		}

		var facilities []FacilityResponse
		for _, facility := range station.StationFacility {
			facilities = append(facilities, FacilityResponse{
				Title: facility.Title,
				Type:  facility.Type,
				Img:   facility.Img,
			})
		}

		response = append(response, StationFacilityResponse{
			StationName: station.Name,
			Facilities:  facilities,
		})

		return
	}

	err = errors.New("Station Not found")

	return
}
