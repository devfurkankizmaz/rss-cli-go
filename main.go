package main

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
)

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Feed struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.gamespot.com/feeds/news")
	if err != nil {
		log.Fatal(err)
	}

	jsonData := Feed{
		Title: feed.Title,
		Items: make([]Item, len(feed.Items)),
	}

	for i, item := range feed.Items {
		jsonData.Items[i] = Item{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
		}
	}

	file, err := os.Create("rss_data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RSS datas successfully saved")
}
