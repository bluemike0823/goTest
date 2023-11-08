package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Env *viper.Viper

func init() {
	Env = initViper()
}

func initViper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("env")
	v.SetConfigType("yaml")
	v.AddConfigPath("_assets/")

	err := v.ReadInConfig()

	if err != nil {
		fmt.Println("[ERROR] loading config failed : ", err)
		panic(err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("testGo")
	v.AutomaticEnv()

	fmt.Println("%+v\n", v.AllSettings())
	return v
}
