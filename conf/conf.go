package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("dev-env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) string {
	return viper.GetString(key)
}
