package ll1

type config struct {
	Terminals    []string
	Nonterminals []string
	Start        string
	Dollar       string
	Epsilon      string
	Productions  map[string][][]string
}
