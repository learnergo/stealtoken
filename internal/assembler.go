package internal

import "fmt"

type Context interface {
}

const (
	Add = iota
	Occupy
)

type Assembler interface {
	Combine(baseUrl, address string) string
}

type addAssembler struct {
}

func (a *addAssembler) Combine(baseUrl, address string) string {
	return baseUrl + address
}

func NewAddAssemble(ctx Context) Assembler {
	return &addAssembler{}
}

type occupyAssembler struct {
}

func (a *occupyAssembler) Combine(baseUrl, address string) string {
	return fmt.Sprintf(baseUrl, address)
}

func NewOccupyAssemble(ctx Context) Assembler {
	return &occupyAssembler{}
}
