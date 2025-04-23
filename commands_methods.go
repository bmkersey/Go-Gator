package main

import "fmt"

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdNames[name] = f

}

func (c *commands) run(s *state, cmd command) error {
	handlerFunc, exists := c.cmdNames[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handlerFunc(s, cmd)
}
