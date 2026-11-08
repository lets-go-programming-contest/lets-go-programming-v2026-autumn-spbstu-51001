package main

import (
	"flag"
	"sort"

	"github.com/CWSHBR/task-3/internal/config"
	"github.com/CWSHBR/task-3/internal/input"
	"github.com/CWSHBR/task-3/internal/output"
)

func main() {
	configPath := flag.String("config", "", "path to the YAML config file")
	flag.Parse()

	if *configPath == "" {
		panic("provide a config file in -config <filepath> format")
	}

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		panic("an error occurred while reading config: " + err.Error())
	}

	currencies, err := input.InputFile(cfg.InputFile)
	if err != nil {
		panic("an error occurred while reading currencies data: " + err.Error())
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = output.WriteFile(cfg.OutputFile, currencies)
	if err != nil {
		panic("an error occurred while writing sorted currencies data: " + err.Error())
	}

}
