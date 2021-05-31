package config

import (
	"bytes"
	_ "embed"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var configBytes []byte

type Config struct {
	Name string
}

func (c *Config) WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config file changed: %s", e.Name)
	})
}

func (c *Config) Init() error {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBufferString(string(configBytes)))
	if err != nil {
		return err
	}
	viper.AutomaticEnv()

	return nil
}

// main.go init config
func InitConfig(cfg string) error {
	c := Config{
		Name: cfg,
	}

	//init config file
	if err := c.Init(); err != nil {
		return err
	}

	// watch hot update
	c.WatchConfig()

	return nil
}
