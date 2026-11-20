package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)


func valueToFloat (value string) float64 {
	value = strings.ReplaceAll(value, ",", ".")
	res, err := strconv.ParseFloat(value, 64)
	if (err != nil) {
		panic(err)
	} 
	return res
}

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

func parseXML (input_filename string) ValCurs {
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

func SortValutes(yaml_string string) {
	var input_path, output_path = getFilenames(yaml_string)
	currencies := parseXML(input_path)
	type sortableValute struct {
		valute Valute
		value  float64
	}
	sort.Slice(currencies.Valutes, func(i, j int) bool {
		valueL := valueToFloat(currencies.Valutes[i].Value)
		valueR := valueToFloat(currencies.Valutes[j].Value)
		return valueL < valueR
	})
	writeJSON(output_path, currencies)
}

func writeJSON (output_filename string, currencies ValCurs) {
	result := make([]ValJSON, 0, len(currencies.Valutes))
	for _, valute := range currencies.Valutes {
        result = append(result, ValJSON{
            NumCode:  valute.NumCode,
            CharCode: valute.CharCode,
            Value:    valute.Value,
        })
    }

    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        panic(err)
    }

    os.WriteFile(output_filename, data, 0644)
}