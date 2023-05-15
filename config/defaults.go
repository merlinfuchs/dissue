package config

import "github.com/spf13/viper"

func setupDefaults() {
	v := viper.GetViper()

	v.SetDefault("db.path", "data")

	v.SetDefault("api.host", "localhost")
	v.SetDefault("api.port", 8080)
}
