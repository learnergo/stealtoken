package main

import (
	"fmt"
	"log"
	"os"
)

// 相对$GOPATH/src 目录
var recodeFile = "github.com/learnergo/stealtoken/stealtoken.suc"
var (
	ErrorLog   = log.New(os.Stderr, "[ERROR]", log.LstdFlags)
	DebugLog   = log.New(os.Stdout, "[DEBUG]", log.LstdFlags)
	SuccessLog = NewLog(recodeFile)
)

func NewLog(path string) *log.Logger {

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666) //打开文件
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return log.New(file, "[SUCCESS]", log.LstdFlags)
}
