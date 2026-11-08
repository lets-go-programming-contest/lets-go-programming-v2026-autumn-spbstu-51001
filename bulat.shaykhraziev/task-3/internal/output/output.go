package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CWSHBR/task-3/internal/common"
)

const (
	dirPerm  = 0o750
	filePerm = 0o600
)

func WriteFile(path string, data []common.Currency) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("encode output json: %w", err)
	}

	err = os.WriteFile(path, fileData, filePerm)
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
