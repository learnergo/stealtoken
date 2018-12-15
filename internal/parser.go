/*
请求结果解析方式
*/
package internal

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Parser interface {
	Parse([]byte) (float64, error)
}

type floatParser struct {
}

func (f *floatParser) Parse(result []byte) (amount float64, err error) {
	amount, err = strconv.ParseFloat(string(result), 10)
	return
}

func NewFloatParse(ctx Context) Parser {
	return &floatParser{}
}

type jsonParser struct {
	path string
}

func (f *jsonParser) Parse(result []byte) (amount float64, err error) {
	josnValue := string(result)
	var dat map[string]interface{}
	err = json.Unmarshal([]byte(josnValue), &dat)
	if err != nil {
		return
	}

	if v, ok := dat[f.path]; ok {
		strAmount := v.(string)
		amount, err = strconv.ParseFloat(string(strAmount), 10)
		return
	}
	return 0, errors.New("Not found")
}

func NewJsonParse(ctx Context) Parser {
	return &jsonParser{
		path: ctx.(string),
	}
}
