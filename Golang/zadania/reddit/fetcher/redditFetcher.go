package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type Reddit struct{}

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func (r *Reddit) Fetch() error {
	resp, err := http.Get("https://www.reddit.com/r/golang.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	for _, child := range res.Data.Children {
		fmt.Printf("Title: %s\nURL: %s\n\n", child.Data.Title, child.Data.URL)
	}

	return nil
}

func (r *Reddit) Save(w io.Writer) error {
	resp, err := http.Get("https://www.reddit.com/r/golang.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	for _, child := range res.Data.Children {
		fmt.Fprintf(w, "Title: %s\nURL: %s\n\n", child.Data.Title, child.Data.URL)
	}

	return nil
}
