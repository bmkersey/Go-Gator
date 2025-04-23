package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bmkersey/Go-Gator/internal/config"
	"github.com/bmkersey/Go-Gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Could not reac config: %s\n", err)
		os.Exit(1)
		return
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("error opening postgres: %s\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	appState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	appCommands := commands{
		cmdNames: make(map[string]func(*state, command) error),
	}

	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)
	appCommands.register("reset", handlerReset)
	appCommands.register("users", handlerUsers)
	appCommands.register("agg", handlerAgg)
	appCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	appCommands.register("feeds", handlerFeeds)
	appCommands.register("follow", middlewareLoggedIn(handlerFollow))
	appCommands.register("following", middlewareLoggedIn(handlerFollowing))

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
