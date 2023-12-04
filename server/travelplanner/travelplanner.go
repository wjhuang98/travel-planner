package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"main/config"
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
} // TODO

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
} // TODO

func (tp *TravelPlanner) GetPlaces(places *tripadvisor.Places) {
	key := tp.config.TripAdvisorKey
	url := "https://api.content.tripadvisor.com/api/v1/location/nearby_search?latLong=42.3455%252C-71.10767&key=" + key + "&category=restaurants&radius=100&radiusUnit=mi&language=en"

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
	fmt.Println(fmt.Sprint(params.Distance) + " " + params.Filter + " " + params.Location)

	places := tripadvisor.Places{}
	tp.GetPlaces(&places)

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
