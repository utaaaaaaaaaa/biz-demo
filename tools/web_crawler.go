package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"golang.org/x/net/html"
)

// WordCount stores word frequency for a URL
type WordCount struct {
	URL   string
	Words map[string]int
	Error error
}

// crawlPage fetches and counts words from a URL
func crawlPage(url string, ch chan<- WordCount, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate network delay
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get(url)
	if err != nil {
		ch <- WordCount{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- WordCount{URL: url, Error: err}
		return
	}

	// Extract text from HTML
	words := extractWords(string(body))
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	ch <- WordCount{URL: url, Words: wordMap}
}

// extractWords gets words from HTML content
func extractWords(content string) []string {
	var words []string
	tokenizer := html.NewTokenizer(strings.NewReader(content))
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.TextToken {
			text := strings.TrimSpace(string(tokenizer.Text()))
			if text != "" {
				// Simple word splitting
				for _, word := range strings.Fields(strings.ToLower(text)) {
					// Remove punctuation
					word = strings.Trim(word, ".,!?\"';:()")
					if len(word) > 2 {
						words = append(words, word)
					}
				}
			}
		}
	}
	return words
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://go.dev/blog",
		"https://go.dev/doc",
	}

	// Channel for collecting results
	resultCh := make(chan WordCount, len(urls))
	var wg sync.WaitGroup

	// Start goroutines for each URL
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go crawlPage(url, resultCh, &wg)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect and display results
	totalWords := make(map[string]int)
	for result := range resultCh {
		if result.Error != nil {
			fmt.Printf("Error crawling %s: %v\n", result.URL, result.Error)
			continue
		}
		fmt.Printf("Results for %s:\n", result.URL)
		for word, count := range result.Words {
			fmt.Printf("\t%s: %d\n", word, count)
			totalWords[word] += count
		}
	}

	// Display total stats
	fmt.Printf("\nTotal unique words across all sites: %d\n", len(totalWords))
	fmt.Printf("Time taken: %v\n", time.Since(start))
}