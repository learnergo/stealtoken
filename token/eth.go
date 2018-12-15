package token

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/learnergo/stealtoken/http"
	"github.com/learnergo/stealtoken/internal"
)

type eth struct {
	url       string
	assembler internal.Assembler
	parser    internal.Parser
}

func (b *eth) Generage() (string, string, error) {
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

func (b *eth) Balance(address string) (amount float64, err error) {
	url := b.assembler.Combine(b.url, address)

	result, err := http.Get(url)
	if err != nil {
		return
	}

	return b.parser.Parse(result)

}

func (b *eth) Name() string {
	return "eth"
}

func Neweth(url string, assembler internal.Assembler, parser internal.Parser) *eth {
	return &eth{
		url:       url,
		assembler: assembler,
		parser:    parser,
	}
}
