package main

import (
	"flag"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

}
