package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug  bool
	Server struct {
		Address string
	}
	Context struct {
		Timeout int
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	JWT struct {
		Secret  string
		Expired int
	} `mapstructure:"jwt"`
	Minio struct {
		Host          string
		AccessKey     string
		SecretKey     string
		UseSSL        bool
		BaseURL       string
		Duration      int
		DefaultBucket string
	}
	Google struct {
		Key         interface{}
		RedirectURL string
	}
}

func GetConfig(path string) Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}
