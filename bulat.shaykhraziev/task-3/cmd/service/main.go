package main

import (
	"os"
	"sort"

	serviceConfig "github.com/CWSHBR/task-3/internal/config"
	"github.com/CWSHBR/task-3/internal/input"
	"github.com/CWSHBR/task-3/internal/output"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] == "-config" {
		panic("provide a config file in -config <filepath> format")
	}

	config, err := serviceConfig.ReadConfig(os.Args[2])
	if err != nil {
		panic("an error occurred while reading config: " + err.Error())
	}

	currencies, err := input.InputFile(config.InputFile)
	if err != nil {
		panic("an error occurred while reading currencies data: " + err.Error())
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = output.WriteFile(config.OutputFile, currencies)
	if err != nil {
		panic("an error occurred while writing sorted currencies data: " + err.Error())
	}

}
