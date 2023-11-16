package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type TravelPlanner struct{}

func (a *TravelPlanner) Run() error {
	apiKey := os.Getenv("tripadvisorkey")
	latLong := "42.360083%2C-71.05888"
	radius := "25"

	url := "https://api.content.tripadvisor.com/api/v1/location/nearby_search?latLong=" + latLong + "&key=" + apiKey + "&radius=" + radius + "&radiusUnit=mi&language=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(string(body))

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(body))
	}

	http.HandleFunc("/", handler)

	return http.ListenAndServe(":8080", nil)
}
