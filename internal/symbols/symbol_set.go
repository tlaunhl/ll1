package symbols

type SymbolSet map[string]bool

func NewSymbolSet(symbols []string) SymbolSet {
	ss := make(SymbolSet)

	for _, symbol := range symbols {
		ss[symbol] = true
	}

	return ss
}

func (ss SymbolSet) Contains(symbol string) bool {
	_, ok := ss[symbol]
	return ok
}

func (ss SymbolSet) Insert(symbol string) {
	ss[symbol] = true
}

func (ss SymbolSet) Apply(fnc func(symbol string) error) error {
	for symbol := range ss {
		if err := fnc(symbol); err != nil {
			return err
		}
	}

	return nil
}

func (ss SymbolSet) ApplyNoError(fnc func(symbol string)) {
	for symbol := range ss {
		fnc(symbol)
	}
}
