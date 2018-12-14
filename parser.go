package main

import (
	"encoding/json"
	"errors"
	"strconv"
)

type parser interface {
	parse([]byte) (float64, error)
}

type floatParser struct {
}

func (f *floatParser) parse(result []byte) (amount float64, err error) {
	amount, err = strconv.ParseFloat(string(result), 10)
	return
}

func newFloatParse(ctx Context) parser {
	return &floatParser{}
}

type jsonParser struct {
	path string
}

func (f *jsonParser) parse(result []byte) (amount float64, err error) {
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

func newJsonParse(ctx Context) parser {
	return &jsonParser{
		path: ctx.(string),
	}
}
