package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

type eth struct {
	url       string
	assembler assembler
	parser    parser
}

func (b *eth) generage() (string, string, error) {
	// 创建账户
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	// 私钥:64个字符
	privateKey := hex.EncodeToString(key.D.Bytes())

	// 得到地址：42个字符
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	return privateKey, address, nil
}

func (b *eth) balance(address string) (amount float64, err error) {
	url := b.assembler.combine(b.url, address)

	result, err := Get(url)
	if err != nil {
		return
	}

	return b.parser.parse(result)

}

func neweth(url string, assembler assembler, parser parser) *eth {
	return &eth{
		url:       url,
		assembler: assembler,
		parser:    parser,
	}
}
