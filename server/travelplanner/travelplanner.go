package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"main/config"
	"main/opencage"
	"main/tripadvisor"
)

type TravelPlanner struct {
	config config.Config
}

func (tp *TravelPlanner) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.RedirectSlashes)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	handler := Handler(tp, WithRouter(router), WithServerBaseURL("/api"))
	return http.ListenAndServe(":8080", handler)
}

func (tp *TravelPlanner) GetLatLng(location string, geocode *opencage.Geocode) {
	key := tp.config.OpenCageKey
	url := "https://api.opencagedata.com/geocode/v1/json?q=" + url.QueryEscape(location) + "&key=" + key + "&language=en&pretty=1&no_annotations=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &geocode)
}

func (tp *TravelPlanner) GetDetails(details *tripadvisor.Details, locationID string) {
	key := tp.config.TripAdvisorKey

	url := "https://api.content.tripadvisor.com/api/v1/location/" + locationID + "/details?key=" + key + "&language=en&currency=USD"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &details)
}

func (tp *TravelPlanner) GetPhotos(photos *tripadvisor.Photos, locationID string) {
	key := tp.config.TripAdvisorKey
	url := "https://api.content.tripadvisor.com/api/v1/location/" + locationID + "/photos?key=" + key + "&language=en"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &photos)
}

func (tp *TravelPlanner) GetPlaces(lat string, long string, radius int32, filter string, places *tripadvisor.Places) {
	key := tp.config.TripAdvisorKey
	latLong := lat + "%252C" + long
	// url := "https://api.content.tripadvisor.com/api/v1/location/nearby_search?latLong=" + latLong + "&key=" + key + "&category=" + filter + "&radius=" + fmt.Sprint(radius) + "&radiusUnit=mi&language=en"

	baseURL := "https://api.content.tripadvisor.com/api/v1/location/nearby_search"
	params := url.Values{"key": {key},
		"latLong":    {latLong},
		"category":   {filter},
		"radius":     {fmt.Sprint(radius)},
		"radiusUnit": {"mi"},
		"language":   {"en"},
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
	}

	u.RawQuery = params.Encode()
	url := u.String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &places)
}

func (tp *TravelPlanner) Search(w http.ResponseWriter, r *http.Request, params SearchParams) *Response {
	// fmt.Printf("Ayo we got shit")
	// fmt.Println(fmt.Sprint(params.Radius) + " " + params.Filter + " " + params.Location)

	geocode := opencage.Geocode{}
	tp.GetLatLng(params.Location, &geocode)
	lat := fmt.Sprint(geocode.Results[0].Geometry.Lat)
	long := fmt.Sprint(geocode.Results[0].Geometry.Lng)

	places := tripadvisor.Places{}
	tp.GetPlaces(lat, long, params.Radius, params.Filter, &places)

	fmt.Println(fmt.Sprint(places))

	response := []Place{}
	for i := range places.Data {
		photos := tripadvisor.Photos{}
		tp.GetPhotos(&photos, places.Data[i].LocationID)
		photosList := []string{}
		for j := range photos.Data {
			photosList = append(photosList, photos.Data[j].Images.Original.URL)
		}

		details := tripadvisor.Details{}
		tp.GetDetails(&details, places.Data[i].LocationID)

		response = append(response, Place{
			Address: places.Data[i].AddressObj.AddressString,
			Name:    places.Data[i].Name,
			Photos:  photosList,     // FOR TESTING
			Rating:  details.Rating, // FOR TESTING
			URL:     details.WebURL,
		})
	}

	return SearchJSON200Response(response)
}
