package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

var errMissingConfigFields = errors.New("required configuration fields are missing")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CurrencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCursXML struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []CurrencyXML `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func sortCurrencies(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func loadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", path, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("decode YAML %q: %w", path, err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, fmt.Errorf("config %q: %w", path, errMissingConfigFields)
	}

	return config, nil
}

func loadRates(path string) ([]CurrencyXML, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML %q: %w", path, err)
	}

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var rates ValCursXML

	decodeErr := decoder.Decode(&rates)
	closeErr := xmlFile.Close()

	if decodeErr != nil {
		return nil, fmt.Errorf("decode XML %q: %w", path, decodeErr)
	}

	if closeErr != nil {
		return nil, fmt.Errorf("close XML %q: %w", path, closeErr)
	}

	return rates.Currencies, nil
}

func convertCurrencies(raw []CurrencyXML) ([]Currency, error) {
	currencies := make([]Currency, 0, len(raw))

	for _, currency := range raw {
		numCode, err := strconv.Atoi(currency.NumCode)
		if err != nil {
			return nil, fmt.Errorf("parse NumCode %q: %w", currency.NumCode, err)
		}

		valueText := strings.ReplaceAll(currency.Value, ",", ".")

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			return nil, fmt.Errorf("parse Value %q: %w", currency.Value, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: currency.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func saveJSON(path string, currencies []Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, 0o750)
	if err != nil {
		return fmt.Errorf("create directory %q: %w", dir, err)
	}

	err = os.WriteFile(path, data, 0o600)
	if err != nil {
		return fmt.Errorf("write JSON %q: %w", path, err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "", "Path to config")
	flag.Parse()

	if *configPath == "" {
		panic("Expected path to config")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	rawCurrencies, err := loadRates(config.InputFile)
	if err != nil {
		panic(err)
	}

	currencies, err := convertCurrencies(rawCurrencies)
	if err != nil {
		panic(err)
	}

	sortCurrencies(currencies)

	err = saveJSON(config.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}
