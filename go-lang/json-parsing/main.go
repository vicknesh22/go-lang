package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// {"page":"words","input":"word1","words":["word1"]}

//creating a struct and adding metadata for the vars
type Words struct {
	Page string `json:"page"`
	Input string `json:"input"`
	Words []string `json:"words"`
}

func main() {

	args := os.Args

	if len(args) < 1 {
		fmt.Print("Usage has to pass valid url\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("pass a valid url - current one is invalid")
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)

	}
	
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}


	if response.StatusCode != 200 {
		fmt.Printf("Invalid status code output: %d  %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	json.Unmarshal(body, &words)

	fmt.Printf("Json parsed \nPage: %s\nWords: %v\nInput: %s ", words.Page, strings.Join(words.Words, ", "), words.Input)

}
