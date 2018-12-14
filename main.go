package main

import (
	"fmt"
)

func start(t token) {

	for {

		private, address, err := t.generage()
		if err != nil {
			fmt.Println("failed to generage", err)
		}
		balance, err := t.balance(address)
		if err != nil {
			fmt.Println("failed to balance", err)
		}
		if balance > 0 {
			recode(fmt.Sprintf("private:%s\naddress:%s\nbalance:%d", private, address, balance))
			fmt.Println("found!!")
		}
	}

}

func main() {
	btc := newbtc(
		"https://blockchain.info/q/addressbalance/",
		newAddAssemble(nil),
		newFloatParse(nil),
	)
	eth := neweth(
		"https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=eth",
		newOccupyAssemble(nil),
		newJsonParse("result"),
	)
	go start(btc)
	start(eth)
}
