package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
)

var (
	url = flag.String("url", "http://naver.com", "request url")
)

func main() {
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(strings.NewReader(string(data)))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Type: %s\n", og.Type)
	fmt.Printf("Title: %s\n", og.Title)
	fmt.Printf("URL: %s\n", og.URL)
	fmt.Printf("String/JSON Representation: %s\n", og)
}
