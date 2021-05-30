package config

import (
	"bytes"
	"embed"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//go:embed conf
var configBytes embed.FS

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
	by, err := configBytes.ReadFile("conf/config.yaml")
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBufferString(string(by)))
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
