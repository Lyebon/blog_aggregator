package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	handler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error{
	command := cmd.name
	if command == "" {
		return fmt.Errorf("the command is was not pass it")
	}

	function := c.handler[command]
	if function == nil {
		return fmt.Errorf("theres no command %v set", command)
	}

	err := function(s, cmd)
	return err
}