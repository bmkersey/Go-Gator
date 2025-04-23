package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough args supplied: Login needs 1 arg")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error occured while setting user: %s", err)
	}

	fmt.Printf("User: %s has been set\n", cmd.args[0])

	return nil
}
