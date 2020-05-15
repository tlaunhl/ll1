package tables

import (
	"github.com/tlaunhl/ll1/internal/config"
	"github.com/tlaunhl/ll1/internal/errors"
)

type ProductionTable map[string][]ProductionBody

func NewProductionTable(st *SymbolTable, cfg config.Config) (ProductionTable, error) {
	pt := make(ProductionTable)

	if err := st.Nonterminals.Apply(func(nt string) error {
		bodies, ok := cfg.Productions[nt]

		if !ok {
			return &errors.ProductionNotFoundError{Nonterminal: nt}
		}

		for _, body := range bodies {
			pb := ProductionBody(body)

			if err := pb.ApplyForward(func(symbol string) error {
				if st.Ref.IsEpsilon(symbol) && len(pb) > 1 {
					return &errors.EpsilonProductionError{Nonterminal: nt}
				}

				return nil
			}); err != nil {
				return err
			}

			pt.addEntry(nt, ProductionBody(pb))
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return pt, nil
}

func (prdt ProductionTable) addEntry(nt string, pb ProductionBody) {
	_, ok := prdt[nt]

	if !ok {
		prdt[nt] = make([]ProductionBody, 0)
	}

	prdt[nt] = append(prdt[nt], pb)
}

func (prdt ProductionTable) Apply(nt string, fnc func(pb ProductionBody) error) error {
	pbs, ok := prdt[nt]

	if !ok {
		return &errors.ProductionNotFoundError{Nonterminal: nt}
	}

	for _, pb := range pbs {
		if err := fnc(pb); err != nil {
			return err
		}
	}

	return nil
}
