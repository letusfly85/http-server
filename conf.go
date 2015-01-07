package main

import (
	"log"
	"os"

	"code.google.com/p/gcfg"
)

type Config struct {
	Server struct {
		HostName       string
		PortNumber     string
		ConnectionType string
		DocumentRoot   string
	}
}

func GetConfig(confName string) (cfg Config, err error) {
	//conf.gcfgが存在しない場合は異常終了
	if _, err := os.Stat(confName); os.IsNotExist(err) {
		log.Fatal(err.Error())
	}

	err = gcfg.ReadFileInto(&cfg, confName)
	check(err)

	return cfg, err
}
