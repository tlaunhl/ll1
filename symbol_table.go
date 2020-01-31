package ll1

type symbolTable struct {
	terminals    symbolSet
	nonterminals symbolSet
	start        string
	dollar       string
	epsilon      string
}

func newSymbolTable(cfg config) *symbolTable {
	st := &symbolTable{make(symbolSet), make(symbolSet), cfg.Start, cfg.Dollar, cfg.Epsilon}

	for _, s := range cfg.Terminals {
		st.terminals.insert(s)
	}

	for _, s := range cfg.Nonterminals {
		st.nonterminals.insert(s)
	}

	return st
}

func (st *symbolTable) isTerminal(s string) bool {
	return st.terminals.contains(s)
}

func (st *symbolTable) getNonterminals() symbolSet {
	return st.nonterminals
}

func (st *symbolTable) isNonterminal(s string) bool {
	return st.nonterminals.contains(s)
}

func (st *symbolTable) getStart() string {
	return st.start
}

func (st *symbolTable) isEpsilon(s string) bool {
	return s == st.epsilon
}

func (st *symbolTable) getEpsilon() string {
	return st.epsilon
}

func (st *symbolTable) isDollar(s string) bool {
	return s == st.dollar
}

func (st *symbolTable) getDollar() string {
	return st.dollar
}
