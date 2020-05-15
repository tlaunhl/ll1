package symbols

type SymbolSetMap map[string]SymbolSet

func (ssm SymbolSetMap) Insert(key, value string) {
	if _, ok := ssm[key]; !ok {
		ssm[key] = make(SymbolSet)
	}

	ssm[key].Insert(value)
}

func (ssm SymbolSetMap) Apply(key string, fnc func(symbol string) error) error {
	ss, ok := ssm[key]

	if !ok {
		return nil
	}

	return ss.Apply(fnc)
}

func (ssm SymbolSetMap) ApplyNoError(key string, fnc func(symbol string)) {
	ss, ok := ssm[key]

	if !ok {
		return
	}

	ss.ApplyNoError(fnc)
}
