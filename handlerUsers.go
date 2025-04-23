package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error getting users from DB: %s\n", err)
		os.Exit(1)
	}

	currentUser := s.cfg.CurrentUser

	for i := range users {
		if users[i].Name != currentUser {
			fmt.Printf("* %s\n", users[i].Name)
			continue
		}

		fmt.Printf("* %s (current)\n", users[i].Name)

	}
	return nil
}
