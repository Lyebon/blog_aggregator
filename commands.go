package main

type command struct {
	name string
	args []string
}

type commands struct {
	command map[string]func(*state, command) error
}

func (c *commands)