package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return &RSSFeed{}, err
	}

	client := &http.Client{}
	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching feed: %s\n", err)
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return &RSSFeed{}, err
	}

	var feed RSSFeed

	err = xml.Unmarshal(bodyData, &feed)
	if err != nil {
		fmt.Printf("Error unmarshalling xml: %s\n", err)
		return &RSSFeed{}, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}
