package common

type Currency struct {
	NumCode  uint16        `xml:"NumCode" json:"num_code"`
	CharCode string        `xml:"CharCode" json:"char_code"`
	Value    CurrencyValue `xml:"Value" json:"value"`
}

type CurrencyList struct {
	Currencies []Currency `xml:"Valute"`
}

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}
