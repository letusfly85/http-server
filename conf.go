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

func GetConfig() Config {
	var cfg Config
	//TODO flgを利用して、設定ファイルを引数指定で変更できるように改修
	confName := "conf.gcfg"

	//conf.gcfgが存在しない場合は異常終了
	if _, err := os.Stat(confName); os.IsNotExist(err) {
		log.Panicln(err.Error())
	}

	err := gcfg.ReadFileInto(&cfg, "conf.gcfg")
	check(err)

	return cfg
}
