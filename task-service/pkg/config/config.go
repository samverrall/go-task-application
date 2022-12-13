package config

type Config struct {
	Address     string
	DBDirectory string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Parse(configFilePath string) error {
	return nil
}
