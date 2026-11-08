package serviceConfig

import (
	"os"

	"github.com/CWSHBR/task-3/internal/common"
	"gopkg.in/yaml.v3"
)

func ReadConfig(filePath string) (*common.Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config common.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
