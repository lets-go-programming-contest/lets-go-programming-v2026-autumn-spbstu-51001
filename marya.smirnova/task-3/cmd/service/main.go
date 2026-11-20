package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	if config.InputFile == "" {
		return Config{}, errors.New("input-file is required")
	}

	if config.OutputFile == "" {
		return Config{}, errors.New("output-file is required")
	}

	return config, nil
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config loaded: input=%q, output=%q\n", config.InputFile, config.OutputFile)
}
