package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/testsuite", testSuite)
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
