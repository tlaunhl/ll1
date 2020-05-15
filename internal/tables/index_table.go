package tables

import "github.com/tlaunhl/ll1/internal/symbols"

type IndexTable symbols.SymbolSetMap

func NewIndexTable(symt *SymbolTable, prdt ProductionTable) (IndexTable, error) {

	idxt := make(IndexTable)

	if err := symt.Nonterminals.Apply(func(nt string) error {
		if err := prdt.Apply(nt, func(pb ProductionBody) error {
			pb.ApplyForwardNoError(func(symbol string) {
				if symt.Nonterminals.Contains(symbol) {
					symbols.SymbolSetMap(idxt).Insert(symbol, nt)
				}
			})

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return idxt, nil
}

func (idxt IndexTable) Apply(key string, fnc func(symbol string) error) error {
	ss, ok := idxt[key]

	if !ok {
		return nil
	}

	for symbol := range ss {
		if err := fnc(symbol); err != nil {
			return err
		}
	}

	return nil
}
