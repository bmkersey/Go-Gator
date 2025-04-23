package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bmkersey/Go-Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough args provided: follow requires url")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error finding feed: %s", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating follow for feed: %s", err)
	}

	fmt.Printf("Feed: %s\nUser: %s", feed.Name, user.Name)
	return nil
}
