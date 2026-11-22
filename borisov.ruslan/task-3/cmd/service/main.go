package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CurrencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCursXML struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []CurrencyXML `xml:"Valute"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, fmt.Errorf("не заполнены обязательные поля в %s", path)
	}

	return config, nil
}

func loadRates(path string) ([]CurrencyXML, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var rates ValCursXML

	err = decoder.Decode(&rates)
	if err != nil {
		return nil, err
	}

	return rates.Currencies, nil
}

func main() {
	configPath := flag.String("config", "", "Путь к конфигурации")
	flag.Parse()

	if *configPath == "" {
		panic("Не указан путь к конфигурации")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := loadRates(config.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Количество валют:", len(currencies))

	for i, currency := range currencies {
		if i >= 3 {
			break
		}

		fmt.Println(currency.NumCode, currency.CharCode, currency.Value)
	}
}
