package ll1

import (
	"testing"

	"github.com/tlaunhl/ll1/internal/errors"
	"github.com/tlaunhl/ll1/test"
)

func TestEpsilonProductionErrorError(t *testing.T) {
	_, err := NewParser(test.EpsilonProductionConfg)
	if _, ok := err.(*errors.EpsilonProductionError); !ok {
		t.Error(err)
	}
}

func TestProductionNotFound(t *testing.T) {
	_, err := NewParser(test.ProductionNotFoundConfig)
	if _, ok := err.(*errors.ProductionNotFoundError); !ok {
		t.Error(err)
	}
}

func TestLeftRecursiveGrammar(t *testing.T) {
	_, err := NewParser(test.LeftRecursiveGrammarConfig)
	if _, ok := err.(*errors.LeftRecursiveGrammarError); !ok {
		t.Error(err)
	}
}

func TestAmbiguousGrammar(t *testing.T) {
	_, err := NewParser(test.AmbiguousGrammarConfig)
	if _, ok := err.(*errors.AmbiguousGrammarError); !ok {
		t.Error(err)
	}
}

func TestUnmatchedSymbols(t *testing.T) {
	if p, err := NewParser(test.UnmatchedSymbolsConfig); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("(")
		if _, ok := err.(*errors.UnmatchedSymbolsError); !ok {
			t.Error(err)
		}
	}
}

func TestUnparsedSymbols(t *testing.T) {
	if p, err := NewParser(test.BracketsSequenceConfig); err != nil {
		t.Error(err)
	} else {
		err := p.Parse("()())")
		if _, ok := err.(*errors.UnparsedSymbolsError); !ok {
			t.Error(err)
		}
	}
}

func TestInvalidInput(t *testing.T) {
	if p, err := NewParser(test.BracketsSequenceConfig); err != nil {
		t.Error(err)
	} else {
		err := p.Parse("\x80")
		if _, ok := err.(*errors.InvalidInputError); !ok {
			t.Error(err)
		}
	}
}

func TestParsingTableEntry(t *testing.T) {
	if p, err := NewParser(test.ParsingTableEntryConfig); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("β")
		if _, ok := err.(*errors.ParsingTableEntryError); !ok {
			t.Error(err)
		}
	}
}

func TestBracketsSequence(t *testing.T) {
	if p, err := NewParser(test.BracketsSequenceConfig); err != nil {
		t.Error(err)
	} else {
		bs := test.NewBracketSequence(5)
		if err := bs.ForEach(func(seq string) error {
			return p.Parse(string([]byte(seq)))
		}); err != nil {
			t.Error(err)
		}
	}
}

func TestArithmeticExpressions(t *testing.T) {
	if p, err := NewParser(test.ArithmeticExpressionsConfig); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("(((α+β)*(β+α))+(α*β+β*α))")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestExtendedArithmeticExpressions(t *testing.T) {
	if p, err := NewParser(test.ExtendedArithmeticExpressionsConfig); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("((α+β)*(α-β)/(α^β+β^γ+γ^α))*(α*β*γ)")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNumericArithmeticExpressionsConfig(t *testing.T) {
	if p, err := NewParser(test.NumericArithmeticExpressionsConfig); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("((1.618+2.718)*(3.141-0.577))^((1.732-1.202)/(2.503+3.359))")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestABCD(t *testing.T) {
	if p, err := NewParser(test.ABCDConfig); err != nil {
		t.Error(err)
	} else {
		t.Run("ab", func(t *testing.T) {
			if err := p.Parse("ab"); err != nil {
				t.Error(err)
			}
		})

		t.Run("ba", func(t *testing.T) {
			if err := p.Parse("ba"); err != nil {
				t.Error(err)
			}
		})

		t.Run("dadb", func(t *testing.T) {
			if err := p.Parse("dadb"); err != nil {
				t.Error(err)
			}
		})

		t.Run("cbca", func(t *testing.T) {
			if err := p.Parse("cbca"); err != nil {
				t.Error(err)
			}
		})

	}
}
