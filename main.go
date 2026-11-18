package main

import (
	"github.com/Lyebon/blog_aggregator/internal/config"	
	"fmt"
)

func main(){
	cfg := config.ReadConfig()
	fmt.Printf("%v",cfg)
	err := config.WriteConfig(cfg)
	if err != nil{
		fmt.Errorf("f")
	}
}