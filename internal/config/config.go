package config

import (
	"os"
	"encoding/json"
	"log"
	"strings"
)

type Config struct{
	DbURL 			string `json:"db_url"`
	CurrentUserName	string `json:"current_user_name"`
}

func basePath()(string, error){
	home, err := os.UserHomeDir()
	if err != nil{
		log.Fatal(err)
	}
	return home, nil
}

func ReadConfig()Config{
	home, err := basePath()
	if err != nil{
		log.Fatal(err)
	}

	path := home + "/.gatorconfig.json"
	file, err := os.ReadFile(path)
	if err != nil{
		log.Fatal(err)
	}
	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg = ReadUser(cfg)
	return cfg
}

func ReadUser(cfg Config) Config{
	home, err := basePath()
	if err != nil{
		log.Fatal(err)
	}

	options := strings.SplitAfter(home, "/")
	cfg.CurrentUserName = options[len(options)-1]
	return cfg
}

func WriteConfig(cfg Config) error{
	home, err := basePath()
	if err != nil{
		log.Fatal(err)
	}
	
	path := home + "/.gatorconfig.json"

	file, err := json.Marshal(cfg)
	if err != nil{
		log.Fatal(err)
	}

	err = os.WriteFile(path, file, 0666)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}