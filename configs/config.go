package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type conf struct {
	VIACEPURL      string `mapstructure:"VIA_CEP_URL"`
	APICEPURL      string `mapstructure:"API_CEP_URL"`
	WEBSERVEPORT   string `mapstructure:"WEB_SERVER_PORT"`
	TIMEOUTCODE    int    `mapstructure:"TIMEOUTCODE"`
	TIMEOUTMESSAGE string `mapstructure:"TIMEOUTMESSAGE"`
}

func LoadEnvironmentVariables(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Panic(err)
	}
	return cfg, err
}
