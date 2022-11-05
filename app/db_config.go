package app

import (
	"fmt"
	"github.com/spf13/viper"
)

var cv *Configs

type Configs struct {
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Server             string `yaml:"server"`
	Port               string `yaml:"port"`
	Database           string `yaml:"database"`
	ConfigInfo         string
	ConfigPath         string
	ErrorPath          string
	ConnectionMasterDB string
}

func InitConfigs() *Configs {
	cv = GetConfig()
	return cv
}

func GetConfig() *Configs {
	cv := &Configs{}
	cv.ConfigPath = "./config/db_connect.yaml"
	cv.ErrorPath = "./config/error.yaml"

	viper.SetConfigFile(cv.ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(cv)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %+v", err)
	}

	cv.ConfigInfo = fmt.Sprintf("server : %s username : %s password : %s port : %s database : %s",
		cv.Server,
		cv.Username,
		cv.Password,
		cv.Port,
		cv.Database)

	cv.ConnectionMasterDB = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cv.Username,
		cv.Password,
		cv.Server,
		cv.Port,
		cv.Database)

	return cv
}
