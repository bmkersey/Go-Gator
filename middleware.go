package main

import (
	"context"

	"github.com/bmkersey/Go-Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return err
		}

		// Call the original handler with user information
		return handler(s, cmd, user)
	}
}
