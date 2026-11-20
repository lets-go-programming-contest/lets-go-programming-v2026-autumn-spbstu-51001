package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)



func getFilenames(yaml_path string) (string, string) {
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

func parse (input_filename string) ValCurs {
	data, err := os.ReadFile(input_filename)
	if (err != nil) {
		panic(err)
	}

	var fileData ValCurs
	
	err = xml.Unmarshal(data, &fileData)
	if (err != nil) {
		panic(err)
	}
	return fileData
}

func SortValutes(yaml_string string) error {
	var input_path, output_path = getFilenames(yaml_string)
	currencies := parse(input_path)

	writeJSON(output_path, currencies)
	return nil
}

func writeJSON (output_filename string, currencies ValCurs) {
	
}