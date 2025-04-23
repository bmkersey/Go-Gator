package main

import "github.com/bmkersey/Go-Gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdNames map[string]func(*state, command) error
}
