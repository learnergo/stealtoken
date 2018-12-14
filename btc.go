package main

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type btc struct {
	url       string
	assembler assembler
	parser    parser
}

func (b *btc) generage() (string, string, error) {
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

func (b *btc) balance(address string) (amount float64, err error) {
	//baseUrl := "https://blockchain.info/q/addressbalance/"
	url := b.assembler.combine(b.url, address)

	result, err := Get(url)
	if err != nil {
		return
	}

	return b.parser.parse(result)

}

func newbtc(url string, assembler assembler, parser parser) *btc {
	return &btc{
		url:       url,
		assembler: assembler,
		parser:    parser,
	}
}
