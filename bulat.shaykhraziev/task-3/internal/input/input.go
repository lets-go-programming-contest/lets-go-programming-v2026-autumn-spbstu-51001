package input

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/CWSHBR/task-3/internal/common"
	"golang.org/x/net/html/charset"
)

func InputFile(filePath string) ([]common.Currency, error) {
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(rawData))
	decoder.CharsetReader = charset.NewReaderLabel

	var data common.CurrencyList

	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("decode input xml: %w", err)
	}

	return data.Currencies, nil
}
