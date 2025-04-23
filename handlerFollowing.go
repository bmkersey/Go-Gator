package main

import (
	"context"
	"fmt"

	"github.com/bmkersey/Go-Gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	followed_feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not find follows: %s", err)
	}

	for _, feed := range followed_feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
