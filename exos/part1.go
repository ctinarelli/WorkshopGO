//+build part1

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ChuckNorrisFact struct {
	Category []string `json:"category"`
	IconURL  string   `json:"icon_url"`
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	Value    string   `json:"value"`
}

const (
	apiUrl = "https://api.chucknorris.io/jokes/random"
	goMax  = 10
)

func getChuckNorrisFact() string {
	var f ChuckNorrisFact
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &f)

	return f.Value
}

func printChuckNorrisFact(c <-chan string, i int, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("%d `%s`\n", i, <-c)
	wg.Done()
}

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < goMax; i++ {
		go printChuckNorrisFact(c, i, &wg)
	}

	for i := 0; i < goMax; i++ {
		c <- getChuckNorrisFact()
	}

	wg.Wait()
}
