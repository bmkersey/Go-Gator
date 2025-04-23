package main

import "github.com/bmkersey/Go-Gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
