package output

import (
	"encoding/json"
	"os"

	"github.com/CWSHBR/task-3/internal/common"
)

func checkFile(filepath string) error {
	f, openErr := os.Open(filepath)
	if openErr != nil {
		f, crErr := os.Create(filepath)
		if crErr != nil {
			return crErr
		}
		closeErr := f.Close()
		if closeErr != nil {
			return closeErr
		}
	}

	closeErr := f.Close()
	if closeErr != nil {
		return closeErr
	}

	return nil
}

func WriteFile(filepath string, data []common.Currency) error {
	err := checkFile(filepath)
	if err != nil {
		return err
	}

	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, fileData, 0666)
	if err != nil {
		return err
	}

	return nil
}
