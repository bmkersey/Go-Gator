package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Printf("error deleting users: %s", err)
		os.Exit(1)
	}
	fmt.Println("Users successfully deleted.")
	return nil
}
