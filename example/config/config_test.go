package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigFromFile(t *testing.T) {
	configFile := "./conf/config_dev.toml"

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("FileName: %s, Error: %v", configFile, err))
	}

	var config *ConfigStruck = &ConfigStruck{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("FileName: %s, Error: %v", configFile, err))
	}

	fmt.Printf("config: %+v\n", config)
}
