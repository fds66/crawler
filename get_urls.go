package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	// Convert any relative URLs into absolute URLs using the base url passed in
	if htmlBody == "" {
		return []string{}, nil
	}
	if baseURL == nil {
		return []string{}, nil
	}
	fmt.Println("Inputs: ", htmlBody, baseURL.String())
	outputStrings := []string{}
	reader := strings.NewReader(htmlBody)
	// look for  all <a> tags if present
	doc, _ := goquery.NewDocumentFromReader(reader)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		link, exists := s.Attr("href")
		if exists {
			text := s.Text()
			fmt.Printf("Link #%d: text: %q, href: %q\n", i+1, text, link)
			linkUrl, _ := url.Parse(link)

			if linkUrl.Scheme == "" {
				linkUrl.Scheme = baseURL.Scheme
			}
			if linkUrl.Host == "" {
				linkUrl.Host = baseURL.Host
			}

			outputStrings = append(outputStrings, linkUrl.String())
		}
	})
	for _, output := range outputStrings {
		fmt.Println("output string: ", output)
	}

	return outputStrings, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	if htmlBody == "" {
		return []string{}, nil
	}
	if baseURL == nil {
		return []string{}, nil
	}
	fmt.Println("Inputs: ", htmlBody, baseURL.String())
	outputStrings := []string{}
	reader := strings.NewReader(htmlBody)
	// look for  all <a> tags if present
	doc, _ := goquery.NewDocumentFromReader(reader)

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		link, exists := s.Attr("src")
		if exists {
			text, _ := s.Attr("alt")
			fmt.Printf("Image #%d: text: %q, src: %q\n", i+1, text, link)
			linkUrl, _ := url.Parse(link)

			if linkUrl.Scheme == "" {
				linkUrl.Scheme = baseURL.Scheme
			}
			if linkUrl.Host == "" {
				linkUrl.Host = baseURL.Host
			}

			outputStrings = append(outputStrings, linkUrl.String())

		}
	})
	for _, output := range outputStrings {
		fmt.Println("output string: ", output)
	}
	return outputStrings, nil

}
