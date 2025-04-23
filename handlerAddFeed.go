package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bmkersey/Go-Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough args supplied: addfeed takes <name> <url> args")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("could not find user: %s", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating new feed: %s", err)
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

	fmt.Printf("Feed was successfully created and followed: %v\n", feed)
	return nil
}
