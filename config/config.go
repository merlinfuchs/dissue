package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetEnvPrefix("dissue")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("WARN could not find config file", err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	setupDefaults()
}
