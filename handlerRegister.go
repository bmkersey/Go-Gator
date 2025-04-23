package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bmkersey/Go-Gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough args supplied: Register needs 1 arg")
	}

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		fmt.Printf("error creating user: %s", err)
		os.Exit(1)
	}

	s.cfg.SetUser(newUser.Name)
	fmt.Printf("User was successfully created: %v\n", newUser)
	return nil
}
