package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile string `yaml:input-file`
	OutputFile string `yaml:output-file`
}

func GetFilenames(yaml_path string) (string, string) {
	data, err := os.ReadFile(yaml_path)
	if err!= nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, config)
	if (err != nil) {
		panic(err)
	}

	return config.InputFile, config.OutputFile
}

func Parse (input_filename string, output_filename string) {

}