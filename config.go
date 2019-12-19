package main

import (
	"github.com/astaxie/beego/config"
	"strings"
	"fmt"
)

type AppConfig struct {
	LogPath string
	LogLevel string
	kafkaAddr string
	KafkaThreadNum int
	LogFiles []string
}

var appConfig = &AppConfig{}

func initConfig(filename string)(err error){
	conf,err := config.NewConfig("ini",filename)
	if err != nil{
		return
	}
	logPath := conf.String("log_path")
	if len(logPath) == 0 {
		return
	}
	logLevel := conf.String("log_level")
	if len(logLevel) == 0 {
		return
	}
	kafkaAddr:= conf.String("kafka_addr")
	if len(kafkaAddr) == 0 {
		return
	}
	logFiles := conf.String("log_files")
	if len(logFiles) == 0 {
		return
	}

	appConfig.KafkaThreadNum,err = conf.Int("kafka_thread_num")
	if err!=nil || appConfig.KafkaThreadNum == 0{
		appConfig.KafkaThreadNum = 8
	}

	arr := strings.Split(logFiles,",")
	for _,v := range arr{
		str := strings.TrimSpace(v)
		if len(str) == 0{
			continue
		}
		appConfig.LogFiles = append(appConfig.LogFiles,str)
	}
	appConfig.kafkaAddr = kafkaAddr
	appConfig.LogLevel = logLevel
	appConfig.LogPath = logPath

	fmt.Printf("load conf success data:%v\n",appConfig)
	return
}