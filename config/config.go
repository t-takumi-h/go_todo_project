package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)



type ConfigList struct{
	DbName string
	SQLDriver string
	SQLUsername string
	SQLPassword string
	SQLAddress string
	WebPort int
}

var Config ConfigList

func init(){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	Config = ConfigList{
		DbName: cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		SQLUsername: cfg.Section("db").Key("username").String(),
		SQLPassword: cfg.Section("db").Key("password").String(),
		SQLAddress: cfg.Section("db").Key("address").String(),
		WebPort: cfg.Section("web").Key("port").MustInt(),
	}
}