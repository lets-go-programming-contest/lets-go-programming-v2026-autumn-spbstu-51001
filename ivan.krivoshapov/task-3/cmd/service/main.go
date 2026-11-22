package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
)

const dirPermissions = 0o755

var (
	errMultipleConfigDocuments = errors.New("must contain one YAML document")
	errMissingConfigPaths      = errors.New("must specify input-file and output-file")
	errUnsupportedXMLEncoding  = errors.New("unsupported XML encoding")
	errTrailingXML             = errors.New("unexpected content after ValCurs")
	errNonFiniteRate           = errors.New("exchange rate must be finite")
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
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	var config Config

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("decode YAML %q: %w", path, err))
	}

	var extra yaml.Node

	err = decoder.Decode(&extra)
	if err == nil {
		panic(fmt.Errorf("config %q: %w", path, errMultipleConfigDocuments))
	}

	if !errors.Is(err, io.EOF) {
		panic(fmt.Errorf("decode YAML %q: %w", path, err))
	}

	config.InputFile = strings.TrimSpace(config.InputFile)
	config.OutputFile = strings.TrimSpace(config.OutputFile)

	if config.InputFile == "" || config.OutputFile == "" {
		panic(fmt.Errorf("config %q: %w", path, errMissingConfigPaths))
	}

	return config
}

func LoadValCurs(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open XML %q: %w", path, err))
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("close XML %q: %v", path, err)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("%w %q", errUnsupportedXMLEncoding, charset)
		}
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Errorf("decode XML %q: %w", path, err))
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return valCurs
			}

			panic(fmt.Errorf("decode XML %q: %w", path, err))
		}

		switch token := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if strings.TrimSpace(string(token)) == "" {
				continue
			}
		}

		panic(fmt.Errorf("%w in %q", errTrailingXML, path))
	}
}

func SortValutesByValueDesc(valutes []Valute) []ValuteJSON {
	result := make([]ValuteJSON, 0, len(valutes))

	for _, valute := range valutes {
		charCode := strings.TrimSpace(valute.CharCode)

		// As in the example, an empty numerical code is represented as 0.
		numCode := 0

		if code := strings.TrimSpace(valute.NumCode); code != "" {
			parsedCode, err := strconv.Atoi(code)
			if err != nil {
				panic(fmt.Errorf("invalid NumCode for currency %q: %w", valute.ID, err))
			}

			numCode = parsedCode
		}

		valueText := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")
		value, err := strconv.ParseFloat(valueText, 64)

		if err == nil && (math.IsNaN(value) || math.IsInf(value, 0)) {
			err = errNonFiniteRate
		}

		if err != nil {
			panic(fmt.Errorf("invalid Value %q for currency %q: %w", valute.Value, valute.ID, err))
		}

		result = append(result, ValuteJSON{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func SaveValutesJSON(valutes []ValuteJSON, path string) {
	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		panic(fmt.Errorf("encode JSON: %w", err))
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		panic(fmt.Errorf("create output directory %q: %w", dir, err))
	}

	// Write in the destination directory so renaming replaces the file atomically.
	file, err := os.CreateTemp(dir, ".valutes-*.json")
	if err != nil {
		panic(fmt.Errorf("create output file for %q: %w", path, err))
	}

	tempPath := file.Name()

	defer func() {
		if err := os.Remove(tempPath); err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Printf("remove temporary JSON %q: %v", tempPath, err)
		}
	}()

	_, writeErr := file.Write(data)
	closeErr := file.Close()

	if writeErr != nil {
		panic(fmt.Errorf("write JSON %q: %w", path, writeErr))
	}

	if closeErr != nil {
		panic(fmt.Errorf("close JSON %q: %w", path, closeErr))
	}

	if err := os.Rename(tempPath, path); err != nil {
		panic(fmt.Errorf("save JSON %q: %w", path, err))
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

	config := LoadConfig(*configPath)
	valCurs := LoadValCurs(config.InputFile)
	sortedValutes := SortValutesByValueDesc(valCurs.Valutes)
	SaveValutesJSON(sortedValutes, config.OutputFile)
}
