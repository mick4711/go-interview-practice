// Package challenge11 contains the solution for Challenge 11.
package challenge11

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ContentFetcher defines an interface for fetching content from URLs
type ContentFetcher interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
}

// ContentProcessor defines an interface for processing raw content
type ContentProcessor interface {
	Process(ctx context.Context, content []byte) (ProcessedData, error)
}

// ProcessedData represents structured data extracted from raw content
type ProcessedData struct {
	Title       string
	Description string
	Keywords    []string
	Timestamp   time.Time
	Source      string
}

// ContentAggregator manages the concurrent fetching and processing of content
type ContentAggregator struct {
	// TODO: Add fields for fetcher, processor, worker count, rate limiter, etc.
	Fetcher           ContentFetcher
	Processor         ContentProcessor
	WorkerCount       int
	RequestsPerSecond int
}

// NewContentAggregator creates a new ContentAggregator with the specified configuration
func NewContentAggregator(
	fetcher ContentFetcher,
	processor ContentProcessor,
	workerCount int,
	requestsPerSecond int,
) *ContentAggregator {
	// validate inputs
	if fetcher == nil ||
		processor == nil ||
		workerCount < 1 ||
		requestsPerSecond < 1 {
		return nil
	}

	return &ContentAggregator{
		Fetcher:           fetcher,
		Processor:         processor,
		WorkerCount:       workerCount,
		RequestsPerSecond: requestsPerSecond,
	}
}

// FetchAndProcess concurrently fetches and processes content from multiple URLs
func (ca *ContentAggregator) FetchAndProcess(
	ctx context.Context,
	urls []string,
) ([]ProcessedData, error) {
	// TODO: Implement concurrent fetching and processing with proper error handling
	// ca.Fetcher.Fetch()
	result := []ProcessedData{}
	for _, url := range urls {
		fetchRes, err := ca.Fetcher.Fetch(ctx, url)
		if err != nil {
			return nil, err
		}

		processedData, err := ca.Processor.Process(ctx, fetchRes)
		if err != nil {
			return nil, err
		}
		result = append(result, processedData)

	}

	return result, nil
}

// Shutdown performs cleanup and ensures all resources are properly released
func (ca *ContentAggregator) Shutdown() error {
	// TODO: Implement proper shutdown logic
	return nil
}

// workerPool implements a worker pool pattern for processing content
func (ca *ContentAggregator) workerPool(
	ctx context.Context,
	jobs <-chan string,
	results chan<- ProcessedData,
	errors chan<- error,
) {
	// TODO: Implement worker pool logic
}

// fanOut implements a fan-out, fan-in pattern for processing multiple items concurrently
func (ca *ContentAggregator) fanOut(
	ctx context.Context,
	urls []string,
) ([]ProcessedData, []error) {
	// TODO: Implement fan-out, fan-in pattern
	return nil, nil
}

// HTTPFetcher is a simple implementation of ContentFetcher that uses HTTP
type HTTPFetcher struct {
	Client *http.Client
	// TODO: Add fields for rate limiting, etc.
}

// Fetch retrieves content from a URL via HTTP
func (hf *HTTPFetcher) Fetch(ctx context.Context, url string) ([]byte, error) {
	// TODO: Implement HTTP-based content fetching with context support
	resp, err := hf.Client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found")
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// HTMLProcessor is a basic implementation of ContentProcessor for HTML content
type HTMLProcessor struct {
	// TODO: Add any fields needed for HTML processing
}

// Process extracts structured data from HTML content
func (hp *HTMLProcessor) Process(ctx context.Context, content []byte) (ProcessedData, error) {
	// TODO: Implement HTML processing logic
	res := ProcessedData{}

	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return res, err
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.Title {
			res.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.DataAtom == atom.Meta {
			if n.Attr[0].Val == "description" {
				res.Description = n.Attr[1].Val
			}
			if n.Attr[0].Val == "keywords" {
				res.Keywords = strings.Split(n.Attr[1].Val, ",")
			}
		}
	}

	if res.Title == "" && res.Description == "" && res.Keywords == nil {
		return res, fmt.Errorf("no data")
	}

	return res, nil
}
