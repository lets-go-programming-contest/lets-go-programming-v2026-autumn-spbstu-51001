package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/text/encoding/charmap"
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
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unsupported XML encoding %q", charset)
		}
	}

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

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func SortValutesByValueDesc(valutes []Valute) []ValuteJSON {
	result := make([]ValuteJSON, 0, len(valutes))

	for _, valute := range valutes {
		charCode := strings.TrimSpace(valute.CharCode)
		if charCode == "" {
			panic(fmt.Errorf("missing CharCode for currency %q", valute.ID))
		}

		// As in the example, an empty numerical code is represented as 0.
		numCode := 0
		if code := strings.TrimSpace(valute.NumCode); code != "" {
			parsedCode, err := strconv.Atoi(code)
			if err != nil {
				panic(fmt.Errorf("invalid NumCode for %s: %w", charCode, err))
			}
			numCode = parsedCode
		}

		valueText := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")
		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
			panic(fmt.Errorf("invalid Value %q for %s", valute.Value, charCode))
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

const dirPermissions = 0o755

func SaveValutesJSON(valutes []ValuteJSON, path string) {
	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		panic(fmt.Errorf("encode JSON: %w", err))
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		panic(fmt.Errorf("create output directory %q: %w", dir, err))
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		panic(fmt.Errorf("save JSON %q: %w", path, err))
	}
}
