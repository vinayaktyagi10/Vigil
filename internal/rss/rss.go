package rss

import (
	"github.com/mmcdole/gofeed"
	"vigil/internal/item"
	"time"
)


func Fetch(url string) ([]item.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	var results []item.Item
	for _, feedItem := range feed.Items {

		var pubTime time.Time
		var authorName string
		if feedItem.PublishedParsed != nil {
			pubTime = *feedItem.PublishedParsed
		}
		if feedItem.Author == nil || feedItem.Author.Name == "" {
			authorName = "Unknown"
		} else {
			authorName = feedItem.Author.Name
		}
		myFeedItem := item.Item{
			Title: feedItem.Title,
			URL:   feedItem.Link,
			Source: authorName,
			PublishedAt: pubTime,
			RawContent: feedItem.Description,
		}
		results = append(results, myFeedItem)
	}
	return results, nil
}

