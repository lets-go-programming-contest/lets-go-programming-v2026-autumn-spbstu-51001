package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is empty")
	}

}

func LoadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("read config %q: %w", path, err))
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var config Config
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Errorf("decode YAML %q: %w", path, err))
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		if err != nil {
			panic(fmt.Errorf("decode YAML %q: %w", path, err))
		}
		panic(fmt.Errorf("config %q must contain one YAML document", path))
	}

	config.InputFile = strings.TrimSpace(config.InputFile)
	config.OutputFile = strings.TrimSpace(config.OutputFile)
	if config.InputFile == "" || config.OutputFile == "" {
		panic(fmt.Errorf("config %q must specify input-file and output-file", path))
	}

	return config
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

func LoadValCurs(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open XML %q: %w", path, err))
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Errorf("decode XML %q: %w", path, err))
	}

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return valCurs
		}
		if err != nil {
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
		panic(fmt.Errorf("unexpected content after ValCurs in %q", path))
	}
}
