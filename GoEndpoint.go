package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Slack          string `json:"Slack"`
	DayOfWeek      string `json:"Current day of the week"`
	CurrentUTCTime string `json:"Current utc time in Nigeria"`
	Track          string `json:"Track"`
	GithubURL      string `json:"Github_url"`
	GithubCodeURL  string `json:"Github_url_code"`
	StatusCode     int    `json:"Status_code"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	slack := r.URL.Query().Get("slack")
	track := r.URL.Query().Get("track")

	// Get the current day of the week and UTC time in Nigeria
	dayOfWeek := time.Now().Weekday().String()
	utcTime := time.Now().In(time.FixedZone("UTC+1", 3600)).Format("2006-01-02 15:04:05")

	// Create the response struct
	response := Response{
		Slack:          slack,
		DayOfWeek:      dayOfWeek,
		CurrentUTCTime: utcTime,
		Track:          track,
		GithubURL:      "https://github.com/ezrahel",
		GithubCodeURL:  "https://github.com/ezrahel/myproject",
		StatusCode:     200,
	}

	// Set the content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the response struct to JSON and send it as the response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the client
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/info", infoHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
