package main

import (
	"fmt"
	"github.com/learnergo/stealtoken/internal"
	"github.com/learnergo/stealtoken/record"
	"github.com/learnergo/stealtoken/token"
)

func start(t token.Token) {

	for {

		private, address, err := t.Generage()
		if err != nil {
			fmt.Println("failed to generage", err)
		}
		balance, err := t.Balance(address)
		if err != nil {
			fmt.Println("failed to balance", err)
		}
		if balance > 0 {
			record.Record(fmt.Sprintf("private:%s\naddress:%s\nbalance:%d", private, address, balance))
			fmt.Println("found!!")
		}
	}

}

func main() {
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
