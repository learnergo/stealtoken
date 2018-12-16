package main

import (
	"fmt"
	"github.com/learnergo/stealtoken/config"
	"github.com/learnergo/stealtoken/internal"
	"github.com/learnergo/stealtoken/token"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Rich struct {
	name    string
	private string
	address string
	balance float64
}

var isDebug = false

func start(ctx context.Context, rich chan<- Rich, t token.Token) {

	go func() {
		for {

			private, address, err := t.Generage()
			if err != nil {
				ErrorLog.Println("failed to generage", err)
				continue
			}
			if isDebug {
				DebugLog.Println("----------- begin ----------")
				DebugLog.Println(fmt.Sprintf("%s:\nprivate:%s\naddress:%s", t.Name(), private, address))
				DebugLog.Println("----------- end -----------")
			}

			balance, err := t.Balance(address)
			if err != nil {
				fmt.Println("failed to balance", err)
				continue
			}
			if balance > 0 {
				rich <- Rich{
					name:    t.Name(),
					private: private,
					address: address,
					balance: balance,
				}
			}

			select {
			case <-ctx.Done():
				return
			default:

			}
		}
	}()

}

// 启动命令
// ./setup.sh up
// 如果开启debug 日志会打印很多日志
func main() {
	// 监控被杀死信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("stealtoken.pid", []byte(strconv.Itoa(pid)), 0777)
	}

	// 读取配置
	c, err := config.NewConfig("config.json")
	if err != nil {
		ErrorLog.Println(err)
		return
	}
	isDebug = c.Debug

	if len(c.TokenConfig) == 0 {
		ErrorLog.Println("empty token")
		return
	}

	// 取消context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rich := make(chan Rich)

	for _, value := range c.TokenConfig {
		t, err := parseToken(value)
		if err != nil {
			ErrorLog.Println(err)
			os.Exit(1)
		}
		go start(ctx, rich, t)
	}

	// 终止监听
	for {
		select {
		case <-interrupt:
			signal.Stop(interrupt)
			cancel()
			time.Sleep(1000 * time.Millisecond) // 为了取消操作的继续进行
			return
		case r := <-rich:
			SuccessLog.Println("----------- begin ----------")
			SuccessLog.Println(fmt.Sprintf("%s:\nprivate:%s\naddress:%s\nbalance:%d", r.name, r.private, r.address, r.balance))
			SuccessLog.Println("----------- end -----------")
			SuccessLog.Println("found!!")
		default:

		}
	}
}

func parseToken(c config.TokenConfig) (token.Token, error) {
	var assembler internal.Assembler
	var parser internal.Parser

	switch c.Assemble {
	case "add":
		assembler = internal.NewAddAssemble(nil)
	case "occupy":
		assembler = internal.NewOccupyAssemble(nil)
	default:
		return nil, errors.New(fmt.Sprintf("Wrong assembler config,c=%+v", c))
	}

	switch c.Parser.Method {
	case "float":
		parser = internal.NewFloatParse(nil)
	case "json":
		parser = internal.NewJsonParse(c.Parser.Route)
	default:
		return nil, errors.New(fmt.Sprintf("Wrong parser config,c=%+v", c))
	}

	switch c.Name {
	case "btc":
		return token.Newbtc(c.Url, assembler, parser), nil
	case "eth":
		return token.Neweth(c.Url, assembler, parser), nil
	default:
		return nil, errors.New(fmt.Sprintf("Wrong token name,c=%+v", c))
	}

}
