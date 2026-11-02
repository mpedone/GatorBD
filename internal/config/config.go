package config

import (
	"fmt"
	"os"
)

type Config struct {
	db_url            string
	current_user_name string
}

const configFileName = ".gatorconfig.json"

func Read() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	cfgLoc := dir + "/" + configFileName
	return cfgLoc
}
