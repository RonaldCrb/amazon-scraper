package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

// AmazonMovieByIDResponse response for the amazon movie scraping functionality
type AmazonMovieByIDResponse struct {
	Title       string   `json:"title"`
	ReleaseYear uint64   `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

// AmazonScrapeMovieByID returns a request response with the structure of AmazonMovieByIDResponse
func AmazonScrapeMovieByID(res http.ResponseWriter, req *http.Request) {
	// capture the URL parameter in a variable
	vars := mux.Vars(req)
	amazonID := vars["amazon_id"]
	// format the targetURL using the amazonID variable
	targetURL := fmt.Sprintf("https://www.amazon.de/gp/product/%s", amazonID)

	movie := AmazonMovieByIDResponse{}

	// Instantiate default collector
	c := colly.NewCollector()

	// scrape movie title
	c.OnHTML(`h1[data-automation-id="title"]`, func(e *colly.HTMLElement) {
		movie.Title = e.Text
	})

	// scrape movie release year
	c.OnHTML(`span[data-automation-id="release-year-badge"]`, func(e *colly.HTMLElement) {
		pry, err := strconv.ParseUint(e.Text, 0, 16)
		if err != nil {
			log.Fatalf("Error converting release year to uint => %v", err)
		}
		movie.ReleaseYear = pry
	})

	// scrape the movie actors
	c.OnHTML(`#meta-info > div > dl:nth-child(2)`, func(e *colly.HTMLElement) {
		var a []string
		e.ForEach(`a[class="_1NNx6V"]`, func(i int, ac *colly.HTMLElement) {
			a = append(a, ac.Text)
		})
		movie.Actors = a
	})

	// scrape the movie poster
	c.OnHTML(`#a-page > div.av-page-desktop.avu-retail-page > div.DVWebNode-detail-atf-wrapper.DVWebNode > div > div > div._3KHiTg._2r7Wei.av-dp-container._13P0S3 > div._3I7QY7.dv-fallback-packshot-image > img`, func(e *colly.HTMLElement) {
		movie.Poster = e.Attr("src")
	})

	// scrape similar ids from recommended movies
	c.OnHTML(`#a-page > div.av-page-desktop.avu-retail-page > div.DVWebNode-detail-btf-wrapper.DVWebNode > div > div:nth-child(2) > div > div > ul`, func(e *colly.HTMLElement) {
		var a []string
		e.ForEach(`a`, func(i int, sm *colly.HTMLElement) {
			s := strings.Split(sm.Attr("href"), "/")
			a = append(a, s[4])
		})
		movie.SimilarIds = a
	})

	// run the scraping
	c.Visit(targetURL)

	// marshal the AmazonMovieByIDResponse struct into a JSON object
	movieJSON, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
	}

	// return the response with the appropriate headers
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", movieJSON)
}
