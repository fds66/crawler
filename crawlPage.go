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
	// lock the numtex for check if first visit and write the initial entry
	cfg.mu.Lock()
	isFirst := cfg.addPageVisit(normCurrent)
	fmt.Printf("crawl url %s, isFirst %t\n", normCurrent, isFirst)
	if isFirst == false {
		cfg.mu.Unlock()
		return
	}

	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error getting content from the currentURL", err.Error())
		return
	}
	fmt.Println("Content from the current page:", body[:200])

	pageData := extractPageData(body, rawCurrentURL)
	// uses mutex.Lock and mutex.Unlock to make sure only one go routine writes at once

	cfg.pages[normCurrent] = pageData
	cfg.mu.Unlock()
	// add in concurrency here
	// uses mutex.Lock and mutex.Unlock to make sure only one go routine writes at once
	// use maxConcurrency to limit how many go routines can run at the same time

	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)

		go func() {
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			defer func() { <-cfg.concurrencyControl }()
			cfg.crawlPage(nextURL)
		}()

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

//func (cfg *config) receive() {
//	fmt.Println(<-cfg.concurrencyControl)

//}
