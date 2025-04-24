package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bmkersey/Go-Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		timeBetweenReqs = 5 * time.Second
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed: %s", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %s", err)
	}

	fetchedFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching RSS feed: %s", err)
	}

	for _, f := range fetchedFeed.Channel.Item {
		pubDate, err := time.Parse(time.RFC3339, f.PubDate)
		if err != nil {
			pubDate, err = time.Parse(time.RFC1123, f.PubDate)
			if err != nil {
				pubDate, err = time.Parse(time.RFC1123Z, f.PubDate)
				if err != nil {
					fmt.Printf("Failed to parse date '%s' with error: %v\n", f.PubDate, err)
					// Use current time as fallback
					pubDate = time.Now()
				}
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       f.Title,
			Url:         f.Link,
			Description: f.Description,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") ||
				strings.Contains(err.Error(), "duplicate key value") {
				continue
			}

			fmt.Printf("error creating post: %v", err)
			continue
		}
		fmt.Println("post saved!")
	}

	return nil
}
