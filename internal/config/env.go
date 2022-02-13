package config

import "github.com/spf13/viper"

func LoadEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		panic("Could not get env var " + key)
	}

	return value
}
