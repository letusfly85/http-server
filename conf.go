package main

import "code.google.com/p/gcfg"

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

	//TODO conf.gcfgが存在しない場合のチェック処理
	//TODO flgを利用して、設定ファイルを引数指定で変更できるように改修
	err := gcfg.ReadFileInto(&cfg, "conf.gcfg")
	check(err)

	return cfg
}
