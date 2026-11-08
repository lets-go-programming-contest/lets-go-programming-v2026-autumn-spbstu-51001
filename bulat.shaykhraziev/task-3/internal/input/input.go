package input

import (
	"encoding/xml"
	"os"

	"github.com/CWSHBR/task-3/internal/common"
)

func InputFile(filePath string) ([]common.Currency, error) {
	rawData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data []common.Currency
	err = xml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return nil, err
}
