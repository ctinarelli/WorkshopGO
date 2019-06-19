//+build part1

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	resp, _ := http.Get(apiUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &f)

	return f.Value
}

func pushChuckNorrisFact(c chan<- string, i int) {
	c <- fmt.Sprintf("%d `%s`", i, getChuckNorrisFact())
}

func main() {
	c := make(chan string)
	for i := 1; i < goMax+1; i++ {
		go pushChuckNorrisFact(c, i)
	}

	i := 0
	for f := range c {
		fmt.Println(f)
		i++
		if i >= goMax {
			break
		}
	}

	close(c)
}
