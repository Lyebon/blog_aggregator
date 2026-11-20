package main

import (
	"fmt"
	"log"

	"github.com/Lyebon/blog_aggregator/internal/config"	
)

type State struct {
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("lyebon")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}