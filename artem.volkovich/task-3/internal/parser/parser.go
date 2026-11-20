package parser

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

func valueToFloat(value string) float64 {
	value = strings.ReplaceAll(value, ",", ".")

	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func getFilenames(yamlPath string) (string, string) {
	data, err := os.ReadFile(yamlPath)

	if err != nil {
		panic(fmt.Errorf("error while reading config %q\n %w", yamlPath, err))
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config.InputFile, config.OutputFile
}

func parseXML(inputFilename string) ValCurs {
	data, err := os.ReadFile(inputFilename)

	if err != nil {
		panic(err)
	}

	var fileData ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))

	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&fileData)
	if err != nil {
		panic(err)
	}

	return fileData
}

func SortValutes(yamlString string) {
	var input_path, outputPath = getFilenames(yamlString)
	currencies := parseXML(input_path)
	sort.Slice(currencies.Valutes, func(i, j int) bool {
		valueL := valueToFloat(currencies.Valutes[i].Value)
		valueR := valueToFloat(currencies.Valutes[j].Value)

		return valueL < valueR
	})
	writeJSON(outputPath, currencies)
}

func writeJSON(outputFilename string, currencies ValCurs) {
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

	err = os.MkdirAll(filepath.Dir(outputFilename), 0755)
    if err != nil {
        panic(err)
    }

	err = os.WriteFile(outputFilename, data, 0600)
	if err != nil {
		panic(err)
	}
}
