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

var errUnsupportedCharset = errors.New("unsupported XML charset")

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type xmlCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
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
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var settings config
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return config{}, fmt.Errorf("parse config %q: %w", path, err)
	}

	return settings, nil
}

func loadCurrencies(path string) ([]currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
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

	values := make([]currency, 0, len(rates.Valutes))

	for index, item := range rates.Valutes {
		numCode := 0

		if text := strings.TrimSpace(item.NumCode); text != "" {
			parsed, err := strconv.Atoi(text)
			if err != nil {
				return nil, fmt.Errorf("currency %d num code: %w", index+1, err)
			}

			numCode = parsed
		}

		normalized := strings.Replace(strings.TrimSpace(item.Value), ",", ".", 1)

		value, err := strconv.ParseFloat(normalized, 64)
		if err != nil {
			return nil, fmt.Errorf("currency %d value: %w", index+1, err)
		}

		values = append(values, currency{
			NumCode:  numCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return values, nil
}

func run(configPath string) error {
	settings, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	values, err := loadCurrencies(settings.InputFile)
	if err != nil {
		return err
	}

	slices.SortStableFunc(values, func(left, right currency) int {
		return cmp.Compare(right.Value, left.Value)
	})

	data, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		return fmt.Errorf("encode currencies: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(settings.OutputFile), directoryPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(settings.OutputFile, data, filePermissions); err != nil {
		return fmt.Errorf("write output %q: %w", settings.OutputFile, err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the YAML config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
