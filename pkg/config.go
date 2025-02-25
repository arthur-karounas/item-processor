package pkg

import "github.com/spf13/viper"

func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath) // полный путь до файла
	return viper.ReadInConfig()
}
