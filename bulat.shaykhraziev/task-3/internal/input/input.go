package input

import (
	"encoding/xml"
	"os"

	"github.com/CWSHBR/task-3/internal/common"
	"golang.org/x/net/html/charset"
)

func InputFile(filePath string) ([]common.Currency, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data common.CurrencyList
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Currencies, err
}
