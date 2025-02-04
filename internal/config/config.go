package config

type Config struct {
	Host     string
	Port     int
	LogLevel string `mapstructure:"log-level"`
	Data     string
	Media    string
}
