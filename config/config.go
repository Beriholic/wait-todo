package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Charset    string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Configuration file loading error,check path: ", err)
	}
	LoadMysql(file)

}

func LoadMysql(file *ini.File) {
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}
