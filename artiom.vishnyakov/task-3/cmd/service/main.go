package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "path to YAML config file")
	flag.Parse()

	if *configPath == "" {
		panic("flag -config is required")
	}

	cfg := loadConfig(*configPath)

	fmt.Printf("input=%s output=%s\n", cfg.InputFile, cfg.OutputFile)
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("decode config %q: %w", path, err))
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("config must contain input-file and output-file")
	}

	return cfg
}
