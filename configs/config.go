package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConfigFileName string
}

func (c *Config) Init() error {
	if c.ConfigFileName != "" {
		viper.SetConfigFile(c.ConfigFileName)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func initConfig(cfg string) error {
	c := Config{
		ConfigFileName: cfg,
	}
	if err := c.Init(); err != nil {
		return err
	}
	return nil
}
