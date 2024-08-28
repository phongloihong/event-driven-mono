package configLoader

import (
	"os"

	"github.com/spf13/viper"
)

type ConfigData struct {
	ServiceName      string `mapstructure:"SERVICE_NAME" default:"cart-bff"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST" default:"localhost"`
	DatabasePort     string `mapstructure:"DATABASE_PORT" default:"27017"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME" default:"root"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD" default:"root"`
	DataBaseSSL      bool   `mapstructure:"DATABASE_SSL" default:"false"`
	DatabaseName     string `mapstructure:"DATABASE_NAME" default:"cart"`
}

func LoadConfig[T any](path string) (config T, err error) {
	// Check if we're in development mode
	isDev := os.Getenv("APP_ENV") == "development"
	if isDev {
		// Development mode: read from app.env
		viper.AddConfigPath(path)
		viper.SetConfigFile("app.env")
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
	} else {
		// Production mode: read from environment variables
		viper.AutomaticEnv()
	}

	// Unmarshal the configuration
	err = viper.Unmarshal(&config)
	return
}
