package main

import (
	"bytes"
	"fmt"
	"os"
	"reddit/fetcher"
)

func main() {
	var r fetcher.RedditFetcher = &fetcher.Reddit{}
	var w bytes.Buffer

	err := r.Save(&w)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write(w.Bytes())
}
