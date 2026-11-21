package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const (
	configFlagName  = "config"
	configFlagUsage = "path to the yaml config file"
)

const floatBitSize = 64

var errEmptyConfigPath = errors.New("config path is not specified")

var errUnsupportedCharset = errors.New("unsupported xml charset")

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type decimal float64

type valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    decimal `xml:"Value"`
}

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []valute `xml:"Valute"`
}

func main() {
	cfg, err := loadConfig(parseConfigPath())
	if err != nil {
		panic(err)
	}

	valutes, err := loadValutes(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(valutes))
}

func parseConfigPath() string {
	path := flag.String(configFlagName, "", configFlagUsage)

	flag.Parse()

	return *path
}

func loadConfig(path string) (config, error) {
	if path == "" {
		return config{}, errEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("reading config file: %w", err)
	}

	var cfg config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config{}, fmt.Errorf("decoding config file: %w", err)
	}

	return cfg, nil
}

func (value *decimal) UnmarshalText(text []byte) error {
	normalized := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)

	parsed, err := strconv.ParseFloat(normalized, floatBitSize)
	if err != nil {
		return fmt.Errorf("parsing value %q: %w", text, err)
	}

	*value = decimal(parsed)

	return nil
}

func loadValutes(path string) ([]valute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	var rates valCurs

	err = decoder.Decode(&rates)
	if err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}

	return rates.Valutes, nil
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, "windows-1251") {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, charset)
}
