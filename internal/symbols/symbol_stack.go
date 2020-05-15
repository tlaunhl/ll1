package symbols

type SymbolStack struct {
	symbols []string
}

func (ss *SymbolStack) Push(symbol string) {
	ss.symbols = append(ss.symbols, symbol)
}

func (ss *SymbolStack) Pop() {
	if len(ss.symbols) == 0 {
		panic("SymbolStack: calling Pop() on empty stack")
	}

	ss.symbols = ss.symbols[:len(ss.symbols)-1]
}

func (ss *SymbolStack) Top() string {
	if len(ss.symbols) == 0 {
		panic("SymbolStack: calling Top() on empty stack")
	}

	return ss.symbols[len(ss.symbols)-1]
}

func (ss *SymbolStack) IsEmpty() bool {
	return len(ss.symbols) == 0
}
