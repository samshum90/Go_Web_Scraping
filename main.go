package main

import (
	"fmt"
	"net/http"

	"github.com/golang-web-scraping/pkg/actions"
	logr "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/scrape", actions.Scrape)
	http.HandleFunc("/crawl", actions.Crawl)
	logr.Info("Starting up on 8080")

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, welcome to your app. Use the following suffix's on the URL to show the different results.\n1)'/scrape' to show results of web scraping.\n2)'/crawl' to show results of a web crawler"
	logr.Info("Received request for the home page")
	w.Write([]byte(fmt.Sprintf(msg)))
}
