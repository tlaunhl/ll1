package tables

import (
	"github.com/tlaunhl/ll1/internal/config"
	"github.com/tlaunhl/ll1/internal/symbols"
)

type SymbolTable struct {
	Terminals    symbols.SymbolSet
	Nonterminals symbols.SymbolSet
	Ref          symbols.SymbolRef
}

func NewSymbolTable(cfg config.Config) *SymbolTable {
	return &SymbolTable{
		Terminals:    symbols.NewSymbolSet(cfg.Terminals),
		Nonterminals: symbols.NewSymbolSet(cfg.Nonterminals),
		Ref:          symbols.NewSymbolRef(cfg),
	}
}
