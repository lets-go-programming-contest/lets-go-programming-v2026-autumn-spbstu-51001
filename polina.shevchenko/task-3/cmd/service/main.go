package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"Value"`
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
