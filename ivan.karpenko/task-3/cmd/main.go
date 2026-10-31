package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type commaFloat float64

type Valute struct {
	NumCode  int        `xml:"NumCode"`
	CharCode string     `xml:"CharCode"`
	Value    commaFloat `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

func (f *commaFloat) UnmarshalText(b []byte) error {
	s := strings.Replace(strings.TrimSpace(string(b)), ",", ".", 1)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", b, err)
	}
	*f = commaFloat(v)
	return nil
}

func ReadXML(path string) []Valute {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open input %q: %w", path, err))
	}
	defer file.Close()

	dec := xml.NewDecoder(file)
	dec.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(label, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset %q", label)
	}

	var vc ValCurs
	if err := dec.Decode(&vc); err != nil {
		panic(fmt.Errorf("decode xml %q: %w", path, err))
	}

	if len(vc.Valutes) == 0 {
		panic("xml: no <Valute> elements")
	}
	for i, v := range vc.Valutes {
		if v.NumCode == 0 || v.CharCode == "" {
			panic(fmt.Sprintf("xml: valute #%d has empty NumCode or CharCode", i))
		}
	}

	return vc.Valutes
}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("parse config %q: %w", path, err))
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		panic("config: input-file is empty")
	}

	if cfg.OutputFile == "" {
		panic("config: output-file is empty")
	}

	return cfg
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is not set: use -config <path>")
	}

	cfg := LoadConfig(*configPath)
	valutes := ReadXML(cfg.InputFile)

}
