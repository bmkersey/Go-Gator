package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %s", err)
	}

	for _, feed := range feeds {
		fmt.Println("Feed:", feed.Name)
		fmt.Println("URL:", feed.Url)
		fmt.Println("User:", feed.Name_2)
		fmt.Println() // empty line between entries
	}

	return nil
}
