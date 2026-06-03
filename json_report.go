package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	// sort the map keys for deterministic output
	// tip : use sort.Strings(keys)
	fmt.Printf("checks, len(pages) %d\n", len(pages))
	var keyList []string
	for key := range pages {
		fmt.Println("key ", key)
		keyList = append(keyList, key)
	}

	sort.Strings(keyList)
	for key := range pages {
		fmt.Println("sorted key ", key)

	}

	fmt.Println("keyList", keyList)
	//build a []pageData slice in sorted order

	sorted := make([]PageData, len(pages))
	for i, key := range keyList {
		sorted[i] = pages[key]
	}
	for _, entry := range sorted {
		fmt.Printf("page data %+v\n\n", entry)
	}
	//marshal it with json.MarshalIndent using indent = 2
	// tip: data, err := json.MarshalIndent(sorted, "", " ")
	data, err := json.MarshalIndent(sorted, "", " ")
	if err != nil {
		fmt.Println("Error marshalling data,", err.Error())
		return err
	}
	fmt.Println(string(data))
	// write the results to disk with os.WriteFile
	// tip os.WriteFile(filename, data, 0644)
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file", err.Error())
		return err
	}
	fmt.Println("output file written")
	return nil
}
