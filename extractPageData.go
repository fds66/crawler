package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	//fmt.Println("Input: ", html)
	// declare the link slices for normalising ?REMOVE?
	//outgoingLinks := []string{}
	//imageLinks := []string{}
	// convert the page url into a url.Url
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("Error in parsing url argument to extractPageData", err.Error())
		return PageData{}
	}
	// get the links from the page
	links, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Println("Error in getting outgoing links", err.Error())
		return PageData{}
	}
	//REMOVE?
	// normalise the links
	/*
		for _, link := range links {
			newlink, err := normalizeURL(link)
			if err != nil {
				fmt.Println("Error normalising links", err.Error())
				return PageData{}
			}
			// add the normalised link to the list (?? check if it's alreadyy in the list)
			outgoingLinks = append(outgoingLinks, newlink)
		}
	*/
	// get the image urls from the page
	iLinks, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		fmt.Println("Error in getting image links", err.Error())
		return PageData{}
	}
	// REMOVE?
	// normalise the links
	/*
		for _, link := range iLinks {
			newlink, err := normalizeURL(link)
			if err != nil {
				fmt.Println("Error normalising links", err.Error())
				return PageData{}
			}
			// add the normalised link to the list (?? check if it's alreadyy in the list)
			imageLinks = append(imageLinks, newlink)
		}
	*/
	output := PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  links,
		ImageURLs:      iLinks,
	}
	//fmt.Printf("output from page struct %+v", output)
	return output
}
