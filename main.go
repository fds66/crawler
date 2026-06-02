package main

import (
	"fmt"
	"os"
)

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
	baseURL := os.Args[1]
	fmt.Println("starting crawl of: ", baseURL)
	/*HTMLString, err := getHTML(baseURL)
	if err != nil {
		fmt.Println("Error returned : ", err.Error())
	}
	fmt.Println(HTMLString)
	*/
	//exampleBaseURL := "https://learnwebscraping.dev/practice/ecommerce/"
	//exampleURL := "https://learnwebscraping.dev/practice/ecommerce/products/ashenfang-longsword-fan-1001/"
	pages := make(map[string]int)
	fmt.Println("starting crawl at ", baseURL)
	crawlPage(baseURL, baseURL, pages)
	fmt.Println("Finished crawl")
	fmt.Println("Final pages map:")
	for key, value := range pages {
		fmt.Printf("%s:  %d\n", key, value)
	}

}
