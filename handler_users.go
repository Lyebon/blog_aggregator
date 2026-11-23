package main

import (
	"log"
	"fmt"
)

func handlerLogin(s *state, cmd command) error{
	if cmd.args == [] {
		return log.Fatalf("error reading config: %v", err)
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User name %v was set", cmd.args[0])
	return nil
}