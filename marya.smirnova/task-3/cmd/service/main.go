package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func loadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	if config.InputFile == "" {
		return Config{}, errors.New("input-file is required")
	}

	if config.OutputFile == "" {
		return Config{}, errors.New("output-file is required")
	}

	return config, nil
}

func parseValue(s string) (float64, error) {
	normalized := strings.ReplaceAll(strings.TrimSpace(s), ",", ".")

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float %q: %w", normalized, err)
	}

	return value, nil
}

func readXML(filename string) ([]Currency, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file: %w", err)
	}
	defer file.Close()

	var data ValCurs

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	currencies := make([]Currency, 0, len(data.Valutes))

	for _, item := range data.Valutes {
		value, err := parseValue(item.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %s: %w", item.CharCode, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func writeJSON(filename string, currencies []Currency) (err error) {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close output file: %w", closeErr))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config loaded: input=%q, output=%q\n", config.InputFile, config.OutputFile)

	currencies, err := readXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully parsed %d currencies\n", len(currencies))

	if err := writeJSON(config.OutputFile, currencies); err != nil {
		panic(err)
	}

	fmt.Printf("Successfully saved currencies to %s\n", config.OutputFile)
}
