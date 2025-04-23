package main

import (
	"context"
	"fmt"

	"github.com/bmkersey/Go-Gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough args provided: follow requires url")
	}

	err := s.db.DeleteFeedFollowByURL(context.Background(), database.DeleteFeedFollowByURLParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	})
	if err != nil {
		return fmt.Errorf("error deleting follow: %s", err)
	}

	fmt.Println("Follow has been removed")
	return nil
}
