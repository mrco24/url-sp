package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Define command-line flags
	fileFlag := flag.String("f", "", "File containing URLs")
	flag.Parse()

	// Check if the -f flag is provided
	if *fileFlag == "" {
		fmt.Println("Usage: go run main.go -f <file>")
		return
	}

	// Read the content of the file
	content := readFile(*fileFlag)

	// Extract URLs
	urls := extractURLs(content)

	// Print the result
	printResult(urls)
}

// readFile reads the content of the specified file
func readFile(filename string) string {
	// Implement file reading logic here
	// You can use ioutil.ReadFile or other file reading methods
	// For simplicity, let's assume the content is provided directly
	return `main url https://metalexpert.com/service/banners.nsf/fixbannerclick?OpenAgent&Redirect=http://www.electrotherm.com/&BannerName=Electrothem&bannerid=8DB3E1347CE2E277C2258450003F3FC1
fast separate url https://metalexpert.com/service/banners.nsf/fixbannerclick?OpenAgent&Redirect=http://www.electrotherm.com/&BannerName=Electrothem
second separate url https://metalexpert.com/service/banners.nsf/fixbannerclick?OpenAgent&Redirect=http://www.electrotherm.com/`
}

// extractURLs extracts URLs from the content
func extractURLs(content string) []string {
	var urls []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		for _, field := range fields {
			// Check if the field looks like a URL
			if strings.HasPrefix(field, "http://") || strings.HasPrefix(field, "https://") {
				urls = append(urls, field)
			}
		}
	}
	return urls
}

// printResult prints the URLs
func printResult(urls []string) {
	for _, url := range urls {
		fmt.Printf("%s\n", url)
	}
}
