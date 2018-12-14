package main

import "fmt"

const (
	Add = iota
	Occupy
)

type assembler interface {
	combine(baseUrl, address string) string
}

type addAssembler struct {
}

func (a *addAssembler) combine(baseUrl, address string) string {
	return baseUrl + address
}

func newAddAssemble(ctx Context) assembler {
	return &addAssembler{}
}

type occupyAssembler struct {
}

func (a *occupyAssembler) combine(baseUrl, address string) string {
	return fmt.Sprintf(baseUrl, address)
}

func newOccupyAssemble(ctx Context) assembler {
	return &occupyAssembler{}
}
