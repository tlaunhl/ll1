package config

type Config struct {
	Terminals    []string
	Nonterminals []string
	Start        string
	Dollar       string
	Epsilon      string
	Productions  map[string][][]string
}
