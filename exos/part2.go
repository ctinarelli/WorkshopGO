//+build part2

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func readBodyOfResponse(in <-chan *http.Response, out chan<- []byte) {
	for r := range in {
		body, _ := ioutil.ReadAll(r.Body)
		out <- body
	}
}

func jsonToChuckNorrisFact(in <-chan []byte, out chan<- ChuckNorrisFact) {
	for bs := range in {
		var f ChuckNorrisFact
		json.Unmarshal(bs, &f)
		out <- f
	}
}

func getStringFact(in <-chan ChuckNorrisFact, out chan<- string) {
	for f := range in {
		out <- f.Value
	}
}

func dispFact(in <-chan string) {
	i := 0
	for f := range in {
		fmt.Println(i, f)
		i++
	}
}

func main() {

	entry := make(chan *http.Response)

	out1 := make(chan []byte)
	go readBodyOfResponse(entry, out1)

	out2 := make(chan ChuckNorrisFact)
	go jsonToChuckNorrisFact(out1, out2)

	out3 := make(chan string)
	go getStringFact(out2, out3)

	go dispFact(out3)

	input := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	for in, err := input.ReadString('\n'); err == nil || err == io.EOF; in, err = input.ReadString('\n') {
		in = strings.TrimRight(in, "\n")
		if err != nil {
			fmt.Println("\nGoodbye :)")
			break
		}
		if in == "end" || in == "quit" || in == "exit" {
			fmt.Println("Goodbye :)")
			break
		}

		if n, parseErr := strconv.ParseInt(in, 10, 64); parseErr == nil {
			for i := int64(0); i < n; i++ {
				if resp, err := http.Get(apiUrl); err == nil {
					entry <- resp
				} else {
					fmt.Printf("Error on request %d: ``.", i, err.Error())
				}
			}
		} else {
			fmt.Printf("Can't parse `%s` as an integer\n", in)
		}

		fmt.Print("> ")
	}
}
