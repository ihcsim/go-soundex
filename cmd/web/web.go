package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ihcsim/soundex"
)

func home(response http.ResponseWriter, request *http.Request) {
	if queries := request.URL.Query(); len(queries) > 0 {
		if names, ok := queries["name"]; ok {
			var body []string
			for _, name := range names {
				code := soundex.Encode(name)
				body = append(body, fmt.Sprintf("%s => [%s]", name, code))
			}

			response.WriteHeader(http.StatusOK)
			response.Write([]byte(strings.Join(body, "\n")))
		}
	} else {
		defaultContent := `
From Wikipedia, https://en.wikipedia.org/wiki/Soundex

Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English. The goal is for homophones to be encoded to the same representation so that they can be matched despite minor differences in spelling.[1] The algorithm mainly encodes consonants; a vowel will not be encoded unless it is the first letter. Soundex is the most widely known of all phonetic algorithms (in part because it is a standard feature of popular database software such as DB2, PostgreSQL,[2] MySQL,[3] Ingres, MS SQL Server[4] and Oracle[5]) and is often used (incorrectly) as a synonym for "phonetic algorithm".[citation needed] Improvements to Soundex are the basis for many modern phonetic algorithms.[6].
`
		if _, err := response.Write([]byte(defaultContent)); err != nil {
			log.Printf("%v\n", err)
		}
	}
}

func testSuite(response http.ResponseWriter, request *http.Request) {
	testData, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Printf("%s\n", err)
		response.Write([]byte(fmt.Sprintf("%v", err)))
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	var testResult bytes.Buffer
	for _, test := range strings.Split(string(testData), "\n") {
		if len(test) == 0 {
			continue
		}
		inputs := strings.Split(test, " ")
		actual, expected := soundex.Encode(inputs[1]), inputs[0]
		output := fmt.Sprintf("%s %s %s\n", inputs[1], expected, actual)
		testResult.Write([]byte(output))
	}

	if _, err := response.Write(testResult.Bytes()); err != nil {
		log.Printf("%v\n", err)
	}
}
