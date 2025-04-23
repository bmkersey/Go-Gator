package main

import (
	"fmt"
	"os"

	"github.com/bmkersey/Go-Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Could not reac config: %s\n", err)
		os.Exit(1)
		return
	}

	appState := state{
		cfg: &cfg,
	}

	appCommands := commands{
		cmdNames: make(map[string]func(*state, command) error),
	}

	appCommands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments supplied")
		os.Exit(1)
		return
	}
	commandName := args[1]
	commandArgs := args[2:]
	command := command{
		name: commandName,
		args: commandArgs,
	}

	err = appCommands.run(&appState, command)
	if err != nil {
		fmt.Printf("error occured while running command: %s\n", err)
		os.Exit(1)
		return
	}

	os.Exit(0)
}
