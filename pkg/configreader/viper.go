package configreader

import "github.com/spf13/viper"

func Init[T any](configFile string) (t T, err error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&t)
	if err != nil {
		return
	}
	return
}
