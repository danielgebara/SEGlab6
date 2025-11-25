package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// This struct stores what each worker returns
type FetchResult struct {
	URL        string
	StatusCode int
	Size       int
	Error      error
}

// Worker function that keeps reading from the jobs channel
// and pushes results into the results channel.
func worker(id int, jobs <-chan string, results chan<- FetchResult) {
	for url := range jobs {
		// fetchURL is where we actually download & measure size
		result := fetchURL(url)
		fmt.Printf("Worker %d finished: %s\n", id, url)
		results <- result
	}
}

// Helper function to fetch a URL and return information
func fetchURL(url string) FetchResult {
	resp, err := http.Get(url)
	if err != nil {
		return FetchResult{URL: url, Error: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FetchResult{URL: url, Error: err}
	}

	return FetchResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		Size:       len(body),
		Error:      nil,
	}
}

func main() {
	start := time.Now()

	// The list of URLs we want to scrape concurrently
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://uottawa.ca",
		"https://httpbin.org/get",
		"https://github.com",
	}

	numWorkers := 5 // required fixed worker count
	jobs := make(chan string, len(urls))
	results := make(chan FetchResult, len(urls))

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send all URLs to the jobs channel
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// Collect results
	fmt.Println("Fetching URLs concurrently using worker pool...\n")

	for i := 0; i < len(urls); i++ {
		res := <-results
		if res.Error != nil {
			fmt.Printf("%s -> ERROR: %v\n", res.URL, res.Error)
		} else {
			fmt.Printf("%s -> Status: %d | Size: %d bytes\n",
				res.URL, res.StatusCode, res.Size)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nScraping complete in %v.\n", elapsed)
}
