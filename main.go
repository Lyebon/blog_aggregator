package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lyebon/blog_aggregator/internal/config"	
)

type state struct {
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	stt := state{
		cfg: &cfg,
	}

	han := make(map[string]func(*state, command) error) 
	cmds := commands{
		handler:han,
	}
	cmds.register("login", handlerLogin)

	line := os.Args
	if len(line) < 2 {
		log.Fatalf("error: not command enter")
	}
	
	cmd := command{
		name: line[1],
		args: line[2:],
	}

	err = cmds.run(&stt, cmd)
	if err != nil{
		log.Fatalf("error login: %v", err)
	}
}