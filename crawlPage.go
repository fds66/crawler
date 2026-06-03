package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// print out the base and current urls so we can monitor how many times this is called, we can interupt if there are too many loops
	fmt.Println("Crawling CurrentURL", rawCurrentURL)
	// parse the urls so we can compare hosts, we only want to crawl this domain
	baseURL := cfg.baseURL
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

	isFirst := cfg.addPageVisit(normCurrent)
	fmt.Printf("crawl url %s, isFirst %t\n", normCurrent, isFirst)
	if isFirst == false {
		return
	}

	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error getting content from the currentURL", err.Error())
		return
	}
	fmt.Println("Content from the current page:", body[:200])
	// original version where I called extractPageData for every link
	/*
		pageData := extractPageData(body, rawCurrentURL)
		for _, link := range pageData.OutgoingLinks {
			fmt.Println("link ", link)
		}

		for _, eachLink := range pageData.OutgoingLinks {
			crawlPage(rawBaseURL, eachLink, pages)
		}
	*/
	// new version where I only get the links
	pageData := extractPageData(body, rawCurrentURL)
	cfg.pages[normCurrent] = pageData

	for _, nextURL := range pageData.OutgoingLinks {
		cfg.crawlPage(nextURL)
	}
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
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	for key, value := range cfg.pages {
		if key == normalizedURL {
			fmt.Println("duplicate page ", value.URL)
			return false
		}
	}
	return true
}
