package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configFlagName  = "config"
	configFlagUsage = "path to the yaml config file"
)

var errEmptyConfigPath = errors.New("config path is not specified")

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	cfg, err := loadConfig(parseConfigPath())
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.InputFile, cfg.OutputFile)
}

func parseConfigPath() string {
	path := flag.String(configFlagName, "", configFlagUsage)

	flag.Parse()

	return *path
}

func loadConfig(path string) (config, error) {
	if path == "" {
		return config{}, errEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("reading config file: %w", err)
	}

	var cfg config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config{}, fmt.Errorf("decoding config file: %w", err)
	}

	return cfg, nil
}
