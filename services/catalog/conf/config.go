package conf

import (
	"os"

	globalConfig "github.com/eavillacis/velociraptor/pkg/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// DBConfiguration ...
type DBConfiguration struct {
	URL string
}

// ErrorTrackerConfiguration ...
type ErrorTrackerConfiguration struct {
	URL string
}

// Configuration ...
type Configuration struct {
	Port         int `default:"4000"`
	DB           *DBConfiguration
	JWT          *globalConfig.JWTConfiguration
	ProjectID    string                     `split_words:"true"`
	ErrorTracker *ErrorTrackerConfiguration `split_words:"true"`	
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()

		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

// LoadConfig loads the configuration from a file
func LoadConfig(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)

	if err := envconfig.Process("CATALOG", config); err != nil {
		return nil, err
	}

	return config, nil
}
