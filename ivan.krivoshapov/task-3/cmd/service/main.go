package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var config Config
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("decode YAML %q: %w", path, err))
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		if err != nil {
			panic(fmt.Errorf("decode YAML %q: %w", path, err))
		}
		panic(fmt.Errorf("config %q must contain one YAML document", path))
	}

	config.InputFile = strings.TrimSpace(config.InputFile)
	config.OutputFile = strings.TrimSpace(config.OutputFile)
	if config.InputFile == "" || config.OutputFile == "" {
		panic(fmt.Errorf("config %q must specify input-file and output-file", path))
	}

	return config
}
