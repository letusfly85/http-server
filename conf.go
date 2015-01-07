package main

import (
	"encoding/json"
	"fmt"
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

func (cfg *Config) ToString() string {
	result, err := json.MarshalIndent(&cfg, "", "     ")
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf(string(result))
}

func GetConfig(confName string) (cfg Config, err error) {
	//conf.gcfgが存在しない場合は異常終了
	if _, err := os.Stat(confName); os.IsNotExist(err) {
		log.Println(err.Error())

		return cfg, err
	}

	err = gcfg.ReadFileInto(&cfg, confName)
	check(err)

	return cfg, err
}
