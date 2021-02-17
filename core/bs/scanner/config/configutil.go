/*
2021-02-10

Written by wowlsh93
*/

package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func InitConfig(configPath string) Configuration {

	viper.SetConfigType("yaml")
	viper.SetConfigName("bs_config") // name of config file (without extension)

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		home, _ := homedir.Dir()
		configPath = home + "/.filecoinscanner/configurations"
		viper.AddConfigPath(configPath)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s", err)
	}

	var configuration Configuration
	err2 := viper.Unmarshal(&configuration)

	if err2 != nil {
		fmt.Println("unable to decode into configuration struct, %v", err2)
	}

	return configuration
}
