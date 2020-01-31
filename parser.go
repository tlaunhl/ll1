package ll1

import (
	"encoding/json"
	"unicode/utf8"
)

// Parser is a direct implementation of LL(1) parser based
// on algorithms presented in the famous Dragon Book
// by Alfred V. Aho et al.
type Parser struct {
	st *symbolTable
	pt parsingTable
}

// Parse builds left derivation while reading
// input string symbols one by one from left to right
// It returns an error if parsing failed
func (p *Parser) Parse(str string) error {
	str += p.st.getDollar()

	var ss stringStack
	ss.push(p.st.getDollar())
	ss.push(p.st.getStart())

	var idx int

	for !ss.isEmpty() {
		st := ss.top()
		rv, w := utf8.DecodeRuneInString(str[idx:])
		sv := string(rv)

		if p.st.isNonterminal(st) {
			pr := p.pt.get(st, sv)

			if pr == nil {
				return &parsingTableEntryNotFoundError{st, sv}
			}

			ss.pop()

			for pos := 0; pos < len(pr); pos++ {
				ss.push(pr[len(pr)-pos-1])
			}
		} else if p.st.isTerminal(st) {
			if st != sv {
				return &parserUnmatchedSymbolsError{st, sv}
			}

			ss.pop()
			idx += w
		} else if p.st.isEpsilon(st) {
			ss.pop()
		} else if p.st.isDollar(st) {
			ss.pop()
			idx += w
			break
		}
	}

	if !ss.isEmpty() {
		return &stackNotEmptyError{}
	}

	if idx != len(str) {
		return &unparsedSymbolsError{}
	}

	return nil
}

// NewParser constructs Parser object from given LL1 grammar
// It returns an error once grammar is erroneous, e.g. left-recursive
// Only most important checks are pefromed
func NewParser(configJSON string) (*Parser, error) {
	var cfg config

	err := json.Unmarshal([]byte(configJSON), &cfg)
	if err != nil {
		return nil, err
	}

	st := newSymbolTable(cfg)
	prt := productionTable(cfg.Productions)

	index, err := buildIndex(prt, st)
	if err != nil {
		return nil, err
	}

	ptb := newParsingTableBuilder(st, prt, index)
	pt, err := ptb.build()
	if err != nil {
		return nil, err
	}

	return &Parser{st, pt}, nil
}
