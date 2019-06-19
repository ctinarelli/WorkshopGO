//+build part2

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func requestChuckNorrisApi(out chan<- *http.Response) {
	for i := 0; i < goMax; i++ {
		resp, _ := http.Get(apiUrl)
		out <- resp
		time.Sleep(time.Second * 1)
	}
	close(out)
}

func readBodyOfResponse(in <-chan *http.Response, out chan<- []byte) {
	for r := range in {
		body, _ := ioutil.ReadAll(r.Body)
		out <- body
	}
	close(out)
}

func jsonToChuckNorrisFact(in <-chan []byte, out chan<- ChuckNorrisFact) {
	for bs := range in {
		var f ChuckNorrisFact
		json.Unmarshal(bs, &f)
		out <- f
	}
	close(out)
}

func getStringFact(in <-chan ChuckNorrisFact, out chan<- string) {
	for f := range in {
		out <- f.Value
	}
	close(out)
}

func main() {
	out1 := make(chan *http.Response, 2)
	go requestChuckNorrisApi(out1)

	out2 := make(chan []byte, 3)
	go readBodyOfResponse(out1, out2)

	out3 := make(chan ChuckNorrisFact, 4)
	go jsonToChuckNorrisFact(out2, out3)

	out4 := make(chan string, 5)
	go getStringFact(out3, out4)

	for s := range out4 {
		fmt.Println(s)
	}
}
