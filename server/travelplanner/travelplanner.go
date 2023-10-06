package main

import (
	"fmt"
	"net/http"
)

type TravelPlanner struct{}

func (a *TravelPlanner) Run() error {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}

	http.HandleFunc("/", handler)

	return http.ListenAndServe(":8080", nil)
}
