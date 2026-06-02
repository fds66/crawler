package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// print out the base and current urls so we can monitor how many times this is called, we can interupt if there are too many loops
	fmt.Println("Crawling CurrentURL", rawCurrentURL)
	// parse the urls so we can compare hosts, we only want to crawl this domain
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing the baseURL", err.Error())
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing the currentURL", err.Error())
		return
	}
	if baseURL.Host != currentURL.Host {
		fmt.Println("Current URL is not in the same domain as Base URL")
		return
	}
	// normalise the currentURL
	normCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalising the currentURL", err.Error())
		return
	}
	for key, value := range pages {
		if key == normCurrent {
			pages[key] += 1
			fmt.Println("duplicate page ", value)
			return
		}
	}
	pages[normCurrent] = 1
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error getting content from the currentURL", err.Error())
		return
	}
	fmt.Println("Content from the current page:", body[:200])
	pageData := extractPageData(body, rawCurrentURL)
	for _, link := range pageData.OutgoingLinks {
		fmt.Println("link ", link)
	}

	for _, eachLink := range pageData.OutgoingLinks {
		crawlPage(rawBaseURL, eachLink, pages)
	}
	/*
		fmt.Println("pages map:")
		for key, value := range pages {
			fmt.Printf("%s:  %d\n", key, value)
		}
	*/
}

/*
type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}
*/
