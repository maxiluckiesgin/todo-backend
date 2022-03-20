package config

import "github.com/spf13/viper"

var Config *viper.Viper

func LoadConfig(configPath string) (*viper.Viper, error) {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath(configPath)
	err := Config.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return Config, nil
}
