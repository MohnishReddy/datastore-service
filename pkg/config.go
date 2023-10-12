package pkg

import (
	"datastore-service/constants"
	"datastore-service/models"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(env constants.ServiceMode) (config *models.Config, err error) {
	configPath, err := getConfigPath(env)
	if err != nil {
		return nil, err
	}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func getConfigPath(env constants.ServiceMode) (path string, err error) {

	switch env {
	case constants.TestMode:
		path = "./config.yml"
	case constants.DataStorePodMode:
		path = "./config-data-store.yml"
	}

	err = validateConfigPath(path)
	if err != nil {
		return "", nil
	}

	return path, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
