package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type WindyAPIParameters struct {
	Latitude  float64
	Longitude float64
}

const WINDYAPI_ENDPOINT = "https://api.windy.com/api/point-forecast/v2"

func (api *WindyAPIParameters) GetSingleResponseByLatAndLong() (string, error) {
	buildReq := api.buildAPIRequest()
	req, err := http.NewRequest("POST", WINDYAPI_ENDPOINT, strings.NewReader(buildReq))
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	windyJsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(windyJsonResp), nil
}

func (api *WindyAPIParameters) buildAPIRequest() string {
	mapRequest := make(map[string]any)
	mapRequest["lat"] = api.Latitude
	mapRequest["lon"] = api.Longitude
	mapRequest["model"] = "gfs"
	mapRequest["parameters"] = []string{"temp", "dewpoint", "precip", "convPrecip", "snowPrecip", "wind", "windGust", "cape", "ptype", "lclouds", "mclouds", "hclouds", "rh", "gh", "pressure"}
	mapRequest["levels"] = []string{"surface", "1000h", "800h", "400h", "200h"}
	mapRequest["key"] = "mxJW8fEadecqILVj7RWBdhUfJ38Ou0Bv"
	jsonRequest, _ := json.Marshal(mapRequest)
	return string(jsonRequest)
}
