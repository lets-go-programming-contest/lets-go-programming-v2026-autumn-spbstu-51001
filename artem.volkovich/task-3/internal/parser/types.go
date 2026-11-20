package parser

type Config struct {
	InputFile string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode int `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal string `xml:"Nominal"`
	Name string `xml:"Name"`
	Value string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`	
}

type ValCurs struct {
	Date string `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type ValJSON struct {
    NumCode  int    `json:"NumCode"`
    CharCode string `json:"CharCode"`
    Value    string `json:"Value"`
}