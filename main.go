package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	fmt.Println("Hello, World!")
	//# usage: ./crawler URL maxConcurrency maxPages
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error converting maxConcurrency to int", err.Error())
		return
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error converting maxPages to int", err.Error())
		return
	}
	fmt.Println("starting crawl of: ", rawBaseURL)
	/*HTMLString, err := getHTML(baseURL)
	if err != nil {
		fmt.Println("Error returned : ", err.Error())
	}
	fmt.Println(HTMLString)
	*/
	//exampleBaseURL := "https://learnwebscraping.dev/practice/ecommerce/"
	//exampleURL := "https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/"
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing the baseURL", err.Error())
		return
	}
	pages := make(map[string]PageData)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	concurr := make(chan struct{}, maxConcurrency)
	defer close(concurr)
	cfg := config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 mu,
		concurrencyControl: concurr,
		wg:                 wg,
		maxPages:           maxPages,
	}
	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("Finished crawl")
	writeJSONReport(cfg.pages, "report.json")
	fmt.Println("Final pages map:")
	for key, value := range pages {
		fmt.Printf("%s:  %s\n", key, value.Heading)
	}

}

/*
type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		wg.Go(func() {
			// Fetch the URL.
			http.Get(url)
		})
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
*/
