package symbols

import (
	"github.com/tlaunhl/ll1/internal/config"
)

type SymbolRef struct {
	start   string
	dollar  string
	epsilon string
}

func NewSymbolRef(cfg config.Config) SymbolRef {
	return SymbolRef{
		start:   cfg.Start,
		dollar:  cfg.Dollar,
		epsilon: cfg.Epsilon,
	}
}

func (sr SymbolRef) Start() string {
	return sr.start
}

func (sr SymbolRef) IsEpsilon(symbol string) bool {
	return symbol == sr.epsilon
}

func (sr SymbolRef) Epsilon() string {
	return sr.epsilon
}

func (sr SymbolRef) IsDollar(symbol string) bool {
	return symbol == sr.dollar
}

func (sr SymbolRef) Dollar() string {
	return sr.dollar
}
