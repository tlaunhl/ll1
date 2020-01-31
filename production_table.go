package ll1

type productionTable map[string][][]string

func (prt productionTable) get(nt string) [][]string {
	pr, ok := prt[nt]

	if !ok {
		return nil
	}

	return pr
}

func buildIndex(prt productionTable, st *symbolTable) (symbolSetMap, error) {
	index := make(symbolSetMap)

	for nt := range st.getNonterminals() {
		prs := prt.get(nt)

		if prs == nil {
			return nil, &productionNotFoundError{nt}
		}
		for _, pr := range prs {
			for _, s := range pr {
				if st.isEpsilon(s) && len(pr) > 1 {
					return nil, &productionEpsilonError{nt}
				}

				if st.isNonterminal(s) {
					index.insert(s, nt)
				}
			}
		}
	}

	return index, nil
}
