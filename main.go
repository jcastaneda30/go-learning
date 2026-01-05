package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
