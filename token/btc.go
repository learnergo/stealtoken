package token

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/learnergo/stealtoken/http"
	"github.com/learnergo/stealtoken/internal"
)

type btc struct {
	url       string
	assembler internal.Assembler
	parser    internal.Parser
}

func (b *btc) Generage() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	if err != nil {
		return "", "", err
	}

	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}

func (b *btc) Balance(address string) (amount float64, err error) {
	//baseUrl := "https://blockchain.info/q/addressbalance/"
	url := b.assembler.Combine(b.url, address)

	result, err := http.Get(url)
	if err != nil {
		return
	}

	return b.parser.Parse(result)

}

func (b *btc) Name() string {
	return "btc"
}

func Newbtc(url string, assembler internal.Assembler, parser internal.Parser) *btc {
	return &btc{
		url:       url,
		assembler: assembler,
		parser:    parser,
	}
}
