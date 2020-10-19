package config

import "github.com/spf13/viper"

var config *viper.Viper
var Collection collection

func Init() {
	config = viper.New()

	config.SetConfigFile("./config/config.yaml")
	_ = config.ReadInConfig()

	config.SetConfigFile("./secret/secret.yaml")
	_ = config.MergeInConfig()
	config.Unmarshal(&Collection)
}

type collection struct {
	Server server
	Redis  redis
	Mysql  mysql
	JWT    jwt
	App    app
}
