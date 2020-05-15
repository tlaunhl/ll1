package ll1

import (
	"encoding/json"
	"unicode/utf8"

	"github.com/tlaunhl/ll1/internal/config"
	"github.com/tlaunhl/ll1/internal/symbols"
	"github.com/tlaunhl/ll1/internal/tables"

	"github.com/tlaunhl/ll1/internal/errors"
)

// Parser is a direct implementation of LL(1) parser based
// on algorithms presented in the famous Dragon Book
// by Alfred V. Aho et al.
type Parser struct {
	symt *tables.SymbolTable
	prst tables.ParsingTable
}

// Parse constructs left derivation while reading
// input string symbols one by one from left to right
// Returns an error if parsing failed
func (p *Parser) Parse(str string) error {
	str += p.symt.Ref.Dollar()

	var ss symbols.SymbolStack
	ss.Push(p.symt.Ref.Dollar())
	ss.Push(p.symt.Ref.Start())

	var idx int

	for !ss.IsEmpty() {
		symbol := ss.Top()
		r, size := utf8.DecodeRuneInString(str[idx:])
		if r == utf8.RuneError {
			return &errors.InvalidInputError{}
		}
		input := string(r)

		// Nonterminal
		if p.symt.Nonterminals.Contains(symbol) {
			ss.Pop()

			if err := p.prst.Apply(symbol, input, func(pb tables.ProductionBody) error {
				pb.ApplyReverseNoError(func(symbol string) {
					ss.Push(symbol)
				})

				return nil
			}); err != nil {
				return err
			}
		}

		// Terminal
		if p.symt.Terminals.Contains(symbol) {
			if symbol != input {
				return &errors.UnmatchedSymbolsError{
					StackSymbol: symbol,
					InputSymbol: input}
			}

			ss.Pop()
			idx += size
		}

		// Epsilon
		if p.symt.Ref.IsEpsilon(symbol) {
			ss.Pop()
		}

		// Dollar
		if p.symt.Ref.IsDollar(symbol) {
			ss.Pop()
			idx += size
			break
		}
	}

	if idx != len(str) {
		return &errors.UnparsedSymbolsError{}
	}

	return nil
}

// NewParser constructs Parser object from given LL1 grammar
// Returns an error once grammar is erroneous, e.g. left-recursive
// Only most important checks are pefromed
func NewParser(configJSON string) (*Parser, error) {
	var cfg config.Config

	err := json.Unmarshal([]byte(configJSON), &cfg)
	if err != nil {
		return nil, err
	}

	symt := tables.NewSymbolTable(cfg)
	prdt, err := tables.NewProductionTable(symt, cfg)
	if err != nil {
		return nil, err
	}

	indt, err := tables.NewIndexTable(symt, prdt)
	if err != nil {
		return nil, err
	}

	fft, err := tables.NewFirstFollowTable(symt, prdt, indt)
	if err != nil {
		return nil, err
	}

	prst, err := tables.NewParsingTable(symt, prdt, fft)
	if err != nil {
		return nil, err
	}

	return &Parser{symt, prst}, nil
}
