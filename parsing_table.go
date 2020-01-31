package ll1

type parsingTable map[string]map[string][]string

func (pt parsingTable) insertOnce(nt, t string, prod []string) bool {
	if _, ok := pt[nt]; !ok {
		pt[nt] = make(map[string][]string)
	}

	if _, ok := pt[nt][t]; ok {
		return false
	}

	pt[nt][t] = prod
	return true
}

func (pt parsingTable) get(nt, t string) []string {
	pr, ok := pt[nt][t]

	if !ok {
		return nil
	}

	return pr
}

type parsingTableBuilder struct {
	st     *symbolTable
	prt    productionTable
	index  symbolSetMap
	first  symbolSetMap
	follow symbolSetMap
	nulls  symbolSet
}

func newParsingTableBuilder(st *symbolTable, prt productionTable, idx symbolSetMap) *parsingTableBuilder {
	return &parsingTableBuilder{st, prt, idx, nil, nil, nil}
}

func (b *parsingTableBuilder) build() (parsingTable, error) {
	b.first = make(symbolSetMap)
	b.nulls = make(symbolSet)

	for nt := range b.st.getNonterminals() {
		fs := make(symbolSet)
		vs := make(symbolSet)

		hasEps, err := b.fillFirst(nt, &fs, &vs)
		if err != nil {
			return nil, err
		}

		if hasEps {
			b.nulls.insert(nt)
		}

		b.first.insertSet(nt, fs)
	}

	b.follow = make(symbolSetMap)
	b.follow.insert(b.st.getStart(), b.st.getDollar())

	for nt := range b.st.getNonterminals() {
		b.fillFollow(nt)
	}

	pt := make(parsingTable)

	for nt := range b.st.getNonterminals() {
		for _, pr := range b.prt.get(nt) {
			for _, s := range pr {
				if b.st.isTerminal(s) {
					if !pt.insertOnce(nt, s, pr) {
						return nil, &grammarAmbiguousError{}
					}

					break
				} else if b.st.isNonterminal(s) {
					for fs := range b.first.get(s) {
						if !pt.insertOnce(nt, fs, pr) {
							return nil, &grammarAmbiguousError{}
						}
					}

					if b.nulls.contains(s) {
						continue
					}

					break
				}
			}
		}

		if b.nulls.contains(nt) {
			for fs := range b.follow.get(nt) {
				if !pt.insertOnce(nt, fs, []string{b.st.getEpsilon()}) {
					return nil, &grammarAmbiguousError{}
				}
			}
		}
	}

	return pt, nil
}

func (b *parsingTableBuilder) fillFirst(nt string, fs *symbolSet, vs *symbolSet) (bool, error) {
	vs.insert(nt)

	var hasEps bool

	for _, pr := range b.prt.get(nt) {
		cnt := 0

		for _, s := range pr {
			if b.st.isEpsilon(s) {
				hasEps = true
				break
			}

			if b.st.isTerminal(s) {
				fs.insert(s)
				break
			}

			if vs.contains(s) {
				return false, &grammarLeftRecursiveError{s}
			}

			hasRecEps, err := b.fillFirst(s, fs, vs)
			if err != nil {
				return false, err
			}

			if hasRecEps {
				cnt++
				continue
			}

			break
		}

		if cnt == len(pr) {
			hasEps = true
		}
	}

	return hasEps, nil
}

func (b *parsingTableBuilder) fillFollow(nt string) {
	for h := range b.index.get(nt) {
		for _, pr := range b.prt.get(h) {
			var isLast bool

			it := find(pr, nt)

			for !it.isEnd() {
				it.inc()
				if it.isEnd() {
					isLast = true
					break
				}

				s := it.get()

				if b.st.isTerminal(s) {
					b.follow.insert(nt, s)
					findFrom(it, nt)
					continue
				}

				b.follow.insertSet(nt, b.first.get(s))

				if b.nulls.contains(s) {
					continue
				}

				break
			}

			if isLast && h != nt {
				b.fillFollow(h)
				b.follow.insertSet(nt, b.follow.get(h))
			}
		}
	}
}
