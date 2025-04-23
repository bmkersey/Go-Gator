package main

import (
	"fmt"

	"github.com/bmkersey/Go-Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Could not reac config: %s\n", err)
		return
	}

	err = cfg.SetUser("Brendan")
	if err != nil {
		fmt.Printf("Error setting user: %s\n", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Could not reac config: %s\n", err)
		return
	}

	fmt.Printf("%+v\n", cfg)
}
