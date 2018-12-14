package record

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var recodeFile = "/opt/gopath/src/github.com/learnergo/stealtoken/data"
var mu sync.Mutex

func Record(item string) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.OpenFile(recodeFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666) //打开文件
	if err != nil {
		fmt.Println(item)
		os.Exit(1)
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, item) //写入文件(字符串)
	if err != nil {
		fmt.Println(item)
		return err
	}
	return nil
}
