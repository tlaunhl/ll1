package tables

import (
	"github.com/tlaunhl/ll1/internal/errors"
)

type ParsingTable map[string]map[string]ProductionBody

func NewParsingTable(symt *SymbolTable, prdt ProductionTable, fft *FirstFollowTable) (ParsingTable, error) {

	prst := make(ParsingTable)

	if err := symt.Nonterminals.Apply(func(nt string) error {
		return prdt.Apply(nt, func(pb ProductionBody) error {
			it := pb.Begin()

			for !it.IsEnd() {
				symbol := it.Get()

				// Terminal
				if symt.Terminals.Contains(symbol) {
					if err := prst.addEntry(nt, symbol, pb); err != nil {
						return err
					}
					break
				}

				// Nonterminal
				if symt.Nonterminals.Contains(symbol) {
					if err := fft.Firsts.Apply(symbol, func(first string) error {
						return prst.addEntry(nt, first, pb)
					}); err != nil {
						return err
					}

					// Nullable nonterminal
					if !fft.Nullables.Contains(symbol) {
						break
					}
				}

				it.Inc()
			}

			if it.IsEnd() {
				return fft.Follows.Apply(nt, func(follow string) error {
					return prst.addEntry(nt, follow, ProductionBody{symt.Ref.Epsilon()})
				})
			}

			return nil
		})

	}); err != nil {
		return nil, err
	}

	return prst, nil
}

func (prst ParsingTable) addEntry(nt, t string, pb ProductionBody) error {
	if _, ok := prst[nt]; !ok {
		prst[nt] = make(map[string]ProductionBody)
	}

	if _, ok := prst[nt][t]; ok {
		return &errors.AmbiguousGrammarError{
			Nonterminal: nt,
			Terminal:    t,
		}
	}

	prst[nt][t] = pb
	return nil
}

func (prst ParsingTable) Apply(nt, t string, fnc func(pb ProductionBody) error) error {
	pb, ok := prst[nt][t]

	if !ok {
		return &errors.ParsingTableEntryError{
			Nonterminal: nt,
			Terminal:    t,
		}
	}

	return fnc(pb)
}
