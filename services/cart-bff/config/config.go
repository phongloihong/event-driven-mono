package config

type ConfigData struct {
	ServiceName      string `mapstructure:"SERVICE_NAME" default:"cart-bff"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST" default:"localhost"`
	DatabasePort     string `mapstructure:"DATABASE_PORT" default:"27017"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME" default:"root"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD" default:"root"`
	DataBaseSSL      bool   `mapstructure:"DATABASE_SSL" default:"false"`
	DatabaseName     string `mapstructure:"DATABASE_NAME" default:"cart"`
}
