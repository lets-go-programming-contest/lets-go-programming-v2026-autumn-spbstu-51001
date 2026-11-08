package output

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/CWSHBR/task-3/internal/common"
)

func WriteFile(path string, data []common.Currency) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, fileData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
