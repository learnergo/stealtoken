package main

import (
	"fmt"
	"github.com/learnergo/stealtoken/internal"
	"github.com/learnergo/stealtoken/token"
	"io/ioutil"
	"strconv"
	"syscall"
)

func start(t token.Token) {

	for {

		private, address, err := t.Generage()
		if err != nil {
			ErrorLog.Println("failed to generage", err)
			continue
		}
		DebugLog.Println("----------- begin ----------")
		DebugLog.Println(fmt.Sprintf("%s:\nprivate:%s\naddress:%s", t.Name(), private, address))
		DebugLog.Println("----------- end -----------")
		balance, err := t.Balance(address)
		if err != nil {
			ErrorLog.Println("failed to balance", err)
			continue
		}
		if balance > 0 {
			SuccessLog.Println("----------- begin ----------")
			SuccessLog.Println(fmt.Sprintf("%s:\nprivate:%s\naddress:%s\nbalance:%d", t.Name(), private, address, balance))
			SuccessLog.Println("----------- end -----------")
			SuccessLog.Println("found!!")
		}
	}

}

// 启动命令
// nohup ./stealtoken 1> stealtoken.out 2> stealtoken.err

func main() {
	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("/opt/gopath/src/github.com/learnergo/stealtoken/stealtoken.pid", []byte(strconv.Itoa(pid)), 0777)
		//defer os.Remove("game_server.pid")
	}

	btc := token.Newbtc(
		"https://blockchain.info/q/addressbalance/",
		internal.NewAddAssemble(nil),
		internal.NewFloatParse(nil),
	)
	eth := token.Neweth(
		"https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=eth",
		internal.NewOccupyAssemble(nil),
		internal.NewJsonParse("result"),
	)
	go start(btc)
	start(eth)
}
