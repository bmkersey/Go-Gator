package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bmkersey/Go-Gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.args) < 1 {
		limit = 2
	} else {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("error converting args: %s", err)
		}
		limit = int32(parsedLimit)
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %s", err)
	}

	for i, post := range posts {
		fmt.Println("------- Post", i+1, "-------")
		fmt.Println("Title:", post.Title)
		fmt.Println("Published:", post.PublishedAt.Format("Jan 02, 2006 15:04:05"))
		fmt.Println("URL:", post.Url)
		fmt.Println("Description:", post.Description)
		fmt.Println()
	}
	return nil
}
