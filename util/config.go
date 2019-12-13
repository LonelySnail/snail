package util

import (
	"encoding/json"
	"fmt"
	"github.com/snail/logger"
	"io/ioutil"
	"os"
)


type Config struct {
	Name 			string
	Version 		string
	Addr 			string
	MaxConn			int
	MAxPacketSize 	int
}

func LoadConfig(configPath string) *Config {
	if configPath == ""{
		dir,_ := os.Getwd()
		configPath =fmt.Sprintf("%s/config.json",dir)
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	conf := new(Config)
	err = json.Unmarshal(data,conf)
	if err != nil {
		panic(err)
	}
	logger.InitZapLog("snail.log")
	return conf
}