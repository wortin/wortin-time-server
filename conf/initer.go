package conf

import "github.com/spf13/viper"

var Vip *viper.Viper

func init() {
	Vip = viper.New()
	Vip.SetConfigFile("conf.yaml")
	err := Vip.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
