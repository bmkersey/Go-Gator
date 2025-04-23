package main

import (
	"github.com/bmkersey/Go-Gator/internal/config"
	"github.com/bmkersey/Go-Gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdNames map[string]func(*state, command) error
}
