package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ronaldcrb/amazon-scraper/scraper"
)

// HealthStatus is a struct with a status string and a timestamp used in server healthchecks
type HealthStatus struct {
	Status    string `json:"status"`
	TimeStamp string `json:"timestamp"`
}

// Start the server
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movie/amazon/{amazon_id}", scraper.AmazonScrapeMovieByID)
	r.HandleFunc("/healthz", healthCheck)
	log.Println("Proxy healthcheck online on http://localhost:8080/healthz")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// healthCheck basic function to assert that the http server is responsive
func healthCheck(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()

	healthStatus := HealthStatus{
		Status:    "amazon-scraper is Healthy!",
		TimeStamp: currentTime.Format("2006-01-02 15:04:05"),
	}

	hj, err := json.Marshal(healthStatus)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", hj)
}
