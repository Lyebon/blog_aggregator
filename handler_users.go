package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.args) == 0 {
		return fmt.Errorf("error: not username pass")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User name %v was set", cmd.args[0])
	return nil
}