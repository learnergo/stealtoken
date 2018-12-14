package main

type token interface {
	generage() (string, string, error)
	balance(string) (amount float64, err error)
}

type Context interface {
}
