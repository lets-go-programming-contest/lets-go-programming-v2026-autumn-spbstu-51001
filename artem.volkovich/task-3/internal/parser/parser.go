package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)



func GetFilenames(yaml_path string) (string, string) {
	data, err := os.ReadFile(yaml_path)
	if err!= nil {
		panic(fmt.Errorf("Error while reading config %q\n %w", yaml_path, err))
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if (err != nil) {
		panic(err)
	}

	return config.InputFile, config.OutputFile
}

func Parse (input_filename string, output_filename string) {
	data, err := os.ReadFile(input_filename)
	if (err != nil) {
		panic(err)
	}

	var fileData ValCurs
	
	err = xml.Unmarshal(data, &fileData)
	if (err != nil) {
		panic(err)
	}
	
	
}