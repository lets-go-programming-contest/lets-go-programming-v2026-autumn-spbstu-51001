package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func main() {
}

func readConfig(path string) map[string]string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := map[string]string{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func readValutes(path string) []Valute {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	data = bytes.ReplaceAll(data, []byte(","), []byte("."))

	decoder := xml.NewDecoder(bytes.NewReader(data))

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	return valCurs.Valutes
}

func sortValutes(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

func saveJSON(path string, valutes []Valute) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
}
