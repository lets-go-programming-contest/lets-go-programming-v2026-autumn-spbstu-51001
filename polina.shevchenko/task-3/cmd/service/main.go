package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
}

func readConfig(path string) map[string]string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := map[string]string{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
