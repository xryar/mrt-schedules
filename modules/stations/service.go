package stations

import (
	"encoding/json"
	"mrt-schedules/common/client"
	"net/http"
	"time"
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
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"
	// hit url
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	// keluarin response
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}
