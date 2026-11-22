package main

import (
	"flag"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "Path to config file")

	flag.Parse()

	if *configPath == "" {
		panic("Can`t fing a path to config")
	}

	data, err := os.ReadFile(*configPath)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		panic(err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("Config file hasn`t expected fields")
	}

	fmt.Println("Begin XML:", config.InputFile)
	fmt.Println("Result JSON:", config.OutputFile)
}
