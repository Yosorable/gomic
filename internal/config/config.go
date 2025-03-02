package config

type Config struct {
	Host      string
	Port      int
	LogLevel  string `mapstructure:"log-level"`
	DataPath  string `mapstructure:"data"`
	MediaPath string `mapstructure:"media"`
	Secret    string
}
