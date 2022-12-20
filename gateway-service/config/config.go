package config

import "github.com/spf13/viper"

type Config struct {
	path string
	*viper.Viper
}

func New(configPath string) *Config {
	v := viper.New()
	return &Config{
		path:  configPath,
		Viper: v,
	}
}

func (c *Config) ParseConfig() error {
	c.SetConfigFile(c.path)
	c.SetConfigType("yaml")
	if err := c.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
