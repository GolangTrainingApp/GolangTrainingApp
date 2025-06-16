package client

import (
	"GoTrainingProject/server"
)

type Client struct {
	Service server.WindyAPIParameters
}

func NewAPIClient(service server.WindyAPIParameters) *Client {
	return &Client{service}
}

func (client *Client) GetWindyAPIDtlsByLatLong(lat float64, long float64) (string, error) {
	client.Service.Latitude = lat
	client.Service.Longitude = long
	return client.Service.GetSingleResponseByLatAndLong()
}
