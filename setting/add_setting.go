package setting

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type AppSetting struct {
	Server struct{
		Ip string
		Port int
	}
	Storage struct {
		Data struct {
			PQ struct {
				ConnectionString string
			}
		}
	}
	Service struct{
		Identity struct {
			Address string
			Timeout int
		}
	}
}

const FilePath string = "config.json"
var instance *AppSetting

func GetAppSetting() *AppSetting {
	if instance == nil {
		instance = getAppConfig()
	}
	return instance
}
func getAppConfig() *AppSetting {
	var config = AppSetting{}
	var v = configureViper()
	setSettingValue(v, &config)
	return &config
}
func setSettingValue(v *viper.Viper, setting *AppSetting) {
	if err := v.Unmarshal(&setting); err != nil {
		log.Fatalln(err.Error())
	}
}
func configureViper() *viper.Viper {
	var v = viper.New()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	setConfigFile(v)
	v.AutomaticEnv()
	return v
}
func setConfigFile(v *viper.Viper)  {
	v.SetConfigFile(FilePath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err.Error())
	}
}