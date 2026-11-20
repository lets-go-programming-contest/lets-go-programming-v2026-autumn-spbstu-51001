package parser

type Config struct {
	InputFile string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode int `xml:"numCode"`
	CharCode string `xml:"charCode"`
	Nominal string `xml:"Nominal"`
	Name string `xml:"Name"`
	Value string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`	
}

type ValCurs struct {
	Date string `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}
