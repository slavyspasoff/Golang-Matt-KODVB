package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	err     error
	url     string
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{err, url, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{nil, url, t}
		resp.Body.Close()
	}
}

func CallURLs() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
