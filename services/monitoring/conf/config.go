package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// MongoConfiguration for mongo config
type MongoConfiguration struct {
	DB       string `split_words:"true"`
	User     string `split_words:"true"`
	Password string `split_words:"true"`
	Port     string `split_words:"true"`
	Host     string `split_words:"true"`
}

// Configuration is the main configuration struct
type Configuration struct {
	ProjectID string `split_words:"true"`
	Mongo     *MongoConfiguration
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

	if err := envconfig.Process("MONITORING", config); err != nil {
		return nil, err
	}

	return config, nil
}
