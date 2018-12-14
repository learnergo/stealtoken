package http

import (
	"io/ioutil"
	"net/http"
)

// 可能以后有代理，所以提取出来
func Get(url string) (result []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
