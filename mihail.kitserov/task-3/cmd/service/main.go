package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
)

var ErrUnsupportedCharset = errors.New("unsupported charset")

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
	jsonPermissions = 0o600
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("error reading config file %q: %v", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("error unmarshaling YAML config %q: %v", path, err))
	}

	if cfg.InputFile == "" {
		panic("in config not set field input-file")
	}

	if cfg.OutputFile == "" {
		panic("in config not set field output-file")
	}

	return cfg
}

func EnsureOutputPath(outputPath string) {
	dir := filepath.Dir(outputPath)

	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		panic(fmt.Sprintf("error creating directory %q: %v", dir, err))
	}

	outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, filePermissions)
	if err != nil {
		panic(fmt.Sprintf("error creating file %q: %v", outputPath, err))
	}

	_ = outputFile.Close()
}

func LoadValCurs(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error opening input-file %q: %v", path, err))
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "utf-8", "":
			return input, nil
		default:
			return nil, ErrUnsupportedCharset
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Sprintf("error decoding XML from %q: %v", path, err))
	}

	return valCurs
}

func parseValue(s string) (float64, error) {
	normalized := strings.ReplaceAll(s, ",", ".")

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float %q: %w", normalized, err)
	}

	return value, nil
}

func SortValutesByValueDesc(valCurs ValCurs) []Valute {
	sorted := make([]Valute, len(valCurs.Valutes))
	copy(sorted, valCurs.Valutes)

	sort.Slice(sorted, func(i, j int) bool {
		valueI, errI := parseValue(sorted[i].Value)
		valueJ, errJ := parseValue(sorted[j].Value)

		if errI != nil || errJ != nil {
			return false
		}

		return valueI > valueJ
	})

	return sorted
}

func SaveValutesJSON(valutes []Valute, outputPath string) {
	result := make([]ValuteJSON, 0, len(valutes))

	for _, valute := range valutes {
		numCode := 0

		numCodeStr := strings.TrimSpace(valute.NumCode)
		if numCodeStr != "" {
			parsedNumCode, err := strconv.Atoi(numCodeStr)
			if err != nil {
				panic(fmt.Sprintf("incorrect NumCode %q for %s: %v", valute.NumCode, valute.CharCode, err))
			}

			numCode = parsedNumCode
		}

		value, err := parseValue(valute.Value)
		if err != nil {
			value = 0
		}

		result = append(result, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling JSON: %v", err))
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		panic(fmt.Sprintf("error creating directory %q: %v", dir, err))
	}

	if err := os.WriteFile(outputPath, data, jsonPermissions); err != nil {
		panic(fmt.Sprintf("error writing file %q: %v", outputPath, err))
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

	cfg := LoadConfig(*configPath)
	EnsureOutputPath(cfg.OutputFile)

	valCurs := LoadValCurs(cfg.InputFile)

	sortedValutes := SortValutesByValueDesc(valCurs)

	SaveValutesJSON(sortedValutes, cfg.OutputFile)
}
