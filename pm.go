package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	urlListFile string
	outputFile  string
)

func init() {
	flag.StringVar(&urlListFile, "l", "", "File containing a list of URLs")
	flag.StringVar(&outputFile, "o", "output.txt", "Output file to store the results")
	flag.Parse()
}

func out(s string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(s + "\n"); err != nil {
		log.Fatal(err)
	}
}

func remover(url string) {
	countOfAmpersands := strings.Count(url, "=")

	fmt.Printf("Number of '=' characters in the URL: %d\n", countOfAmpersands)
	fmt.Printf("Iteration %d: %s\n", '1'-48, url)
	out(url)

	for i := 0; i < countOfAmpersands; i++ {
		reversedURL := reverseString(url)
		indexOfAmpersand := strings.Index(reversedURL, "=")
		substring := reversedURL[indexOfAmpersand+1:]
		finalResult := reverseString(substring)
		fmt.Printf("Iteration %d: %s\n", i+2, finalResult)
		url = finalResult
		out(url)
	}
}

func reverseString(s string) string {
	runeString := []rune(s)
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}
	return string(runeString)
}

func main() {
	if urlListFile == "" {
		fmt.Println("Please provide a file containing a list of URLs using the -l flag.")
		return
	}

	file, err := os.Open(urlListFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	urlRegex := regexp.MustCompile(`(?i)\b(?:https?://|www\.)\S+\b`)

	for scanner.Scan() {
		line := scanner.Text()
		urls := urlRegex.FindAllString(line, -1)

		for _, url := range urls {
			remover(url)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
