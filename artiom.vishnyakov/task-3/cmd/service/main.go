package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to YAML config file")
	flag.Parse()

	if *configPath == "" {
		panic("flag -config is required")
	}

	cfg := loadConfig(*configPath)
	currencies := loadCurrencies(cfg.InputFile)

	fmt.Printf("parsed %d currencies, output=%s\n", len(currencies), cfg.OutputFile)
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("decode config %q: %w", path, err))
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("config must contain input-file and output-file")
	}

	return cfg
}

func loadCurrencies(path string) []Currency {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open XML %q: %w", path, err))
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charsetReader

	var doc ValCurs
	if err := decoder.Decode(&doc); err != nil {
		panic(fmt.Errorf("decode XML %q: %w", path, err))
	}

	result := make([]Currency, 0, len(doc.Valutes))

	for _, v := range doc.Valutes {
		numCode, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			panic(fmt.Errorf("parse NumCode %q: %w", v.NumCode, err))
		}

		valueText := strings.Replace(strings.TrimSpace(v.Value), ",", ".", 1)
		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			panic(fmt.Errorf("parse Value %q: %w", v.Value, err))
		}

		result = append(result, Currency{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(v.CharCode),
			Value:    value,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(strings.TrimSpace(charset)) {
	case "windows-1251", "cp1251":
		return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
	default:
		return nil, fmt.Errorf("unsupported charset %q", charset)
	}
}
