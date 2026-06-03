package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	fmt.Println("Hello, World!")
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
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
	concurr := make(chan struct{}, 3)
	defer close(concurr)
	cfg := config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 mu,
		concurrencyControl: concurr,
		wg:                 wg,
	}
	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("Finished crawl")
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
