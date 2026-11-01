package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const (
	directoryPermissions = 0o755
	filePermissions      = 0o644
)

var (
	errEmptyConfigPath    = errors.New("config path is empty")
	errEmptyInputPath     = errors.New("input-file is empty")
	errEmptyOutputPath    = errors.New("output-file is empty")
	errUnsupportedCharset = errors.New("unsupported XML charset")
)

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type xmlCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type xmlRates struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Valutes []xmlCurrency `xml:"Valute"`
}

type currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(path string) (config, error) {
	if strings.TrimSpace(path) == "" {
		return config{}, errEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var settings config
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return config{}, fmt.Errorf("parse config %q: %w", path, err)
	}

	settings.InputFile = strings.TrimSpace(settings.InputFile)
	settings.OutputFile = strings.TrimSpace(settings.OutputFile)

	if settings.InputFile == "" {
		return config{}, errEmptyInputPath
	}

	if settings.OutputFile == "" {
		return config{}, errEmptyOutputPath
	}

	return settings, nil
}

func parseValue(text string) (float64, error) {
	normalized := strings.Replace(strings.TrimSpace(text), ",", ".", 1)

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("parse currency value %q: %w", text, err)
	}

	return value, nil
}

func decodeCurrencies(input io.Reader) ([]currency, error) {
	decoder := xml.NewDecoder(input)
	decoder.CharsetReader = func(label string, source io.Reader) (io.Reader, error) {
		switch strings.ToLower(label) {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(source), nil
		case "utf-8":
			return source, nil
		default:
			return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, label)
		}
	}

	var rates xmlRates
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode currencies: %w", err)
	}

	result := make([]currency, 0, len(rates.Valutes))

	for index, item := range rates.Valutes {
		numCode := 0

		if text := strings.TrimSpace(item.NumCode); text != "" {
			parsed, err := strconv.Atoi(text)
			if err != nil {
				return nil, fmt.Errorf("currency %d num code: %w", index+1, err)
			}

			numCode = parsed
		}

		value, err := parseValue(item.Value)
		if err != nil {
			return nil, fmt.Errorf("currency %d: %w", index+1, err)
		}

		result = append(result, currency{
			NumCode:  numCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return result, nil
}

func sortCurrencies(values []currency) {
	slices.SortStableFunc(values, func(left, right currency) int {
		return cmp.Compare(right.Value, left.Value)
	})
}

func writeJSON(path string, values []currency) error {
	data, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		return fmt.Errorf("encode currencies: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), directoryPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write output %q: %w", path, err)
	}

	return nil
}

func run(configPath string) error {
	settings, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(settings.InputFile)
	if err != nil {
		return fmt.Errorf("read input %q: %w", settings.InputFile, err)
	}

	values, err := decodeCurrencies(bytes.NewReader(data))
	if err != nil {
		return err
	}

	sortCurrencies(values)

	return writeJSON(settings.OutputFile, values)
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the YAML config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
