package main

import (
	"fmt"
	"github.com/GolangTrainingApp/GolangTrainingApp/client"
	"github.com/GolangTrainingApp/GolangTrainingApp/server"
)

func main() {

	service := server.WindyAPIParameters{}

	apiClient := client.NewAPIClient(service) //Dependency injection
	resp, err := apiClient.GetWindyAPIDtlsByLatLong(53.1900, -112.2500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

/*
	const WINDYAPI_ENDPOINT = "https://api.windy.com/api/point-forecast/v2"

	jsonReq := `{
		    "lat": 53.1900,
		    "lon": -112.2500,
		    "model": "gfs",
		    "parameters": ["temp","dewpoint","precip","convPrecip","snowPrecip","wind","windGust","cape","ptype","lclouds","mclouds","hclouds","rh","gh","pressure"],
		    "levels": ["surface", "1000h", "800h","400h","200h"],
		    "key": "mxJW8fEadecqILVj7RWBdhUfJ38Ou0Bv"
		}`

	req, err := http.NewRequest("POST", WINDYAPI_ENDPOINT, strings.NewReader(jsonReq))
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	io.Copy(os.Stdout, resp.Body)

*/
