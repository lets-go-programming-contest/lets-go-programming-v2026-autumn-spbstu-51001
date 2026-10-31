package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("err in LoadConfig:\n ReadFile Config:\n  path: %s\n  err: %v", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("err in LoadConfig:\n Yaml Unmarshal cfg:\n  path: %s\n  err: %v", path, err))
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		panic(fmt.Errorf("InputFile file must not be empty"))
	}

	if cfg.OutputFile == "" {
		panic(fmt.Errorf("OutputFile file must not be empty"))
	}

	return cfg
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	if strings.TrimSpace(*configPath) == "" {
		panic(fmt.Errorf("config path is not set: use -config <path>"))
	}
}
