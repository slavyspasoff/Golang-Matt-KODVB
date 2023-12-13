package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d</h1>", <-ch)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func Race() {
	var nextID nextCh = make(chan int)
	go counter(nextID)

	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
