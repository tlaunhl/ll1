package ll1

import (
	"testing"
)

type BracketSequence struct {
	bytes []byte
}

func NewBracketSequence(size int) *BracketSequence {
	bs := &BracketSequence{make([]byte, 2*size)}

	for cnt := 0; cnt < size; cnt++ {
		bs.bytes[cnt] = '('
		bs.bytes[cnt+size] = ')'
	}

	return bs
}

func (bs *BracketSequence) Next() bool {
	var openCount, closeCount int

	idx := len(bs.bytes) - 1

	for idx > 0 {
		if bs.bytes[idx] == ')' {
			closeCount++
		} else {
			openCount++
			if closeCount > openCount {
				break
			}
		}

		idx--
	}

	if idx == 0 {
		return false
	}

	bs.bytes[idx] = ')'
	idx++
	closeCount--

	for openCount > 0 {
		bs.bytes[idx] = '('
		idx++
		openCount--
	}

	for closeCount > 0 {
		bs.bytes[idx] = ')'
		idx++
		closeCount--
	}

	return true
}

func (bs *BracketSequence) String() string {
	return string(bs.bytes)
}

func TestProductionEpsilonError(t *testing.T) {
	config := `
	{
		"Nonterminals": ["S"],
		"Epsilon":      "ε",
		"Productions":	{
			"S":	[["S", "ε"]]
		}
	}`

	_, err := NewParser(config)
	if _, ok := err.(*productionEpsilonError); !ok {
		t.Error(err)
	}
}

func TestProductionNotFound(t *testing.T) {
	config := `
	{
		"Nonterminals":	["S", "T"],
		"Productions":	{
			"S":	[["T"]]
		}
	}`

	_, err := NewParser(config)
	if _, ok := err.(*productionNotFoundError); !ok {
		t.Error(err)
	}
}

func TestGrammarLeftRecursive(t *testing.T) {
	config := `
	{
		"Nonterminals": ["S", "T", "R"],
		"Productions":	{
			"S":	[["T"]],
			"T":	[["R"]],
			"R":	[["S"]]
		}
	}`

	_, err := NewParser(config)

	if _, ok := err.(*grammarLeftRecursiveError); !ok {
		t.Error(err)
	}
}

func TestGrammarAmbiguous(t *testing.T) {
	config := `
	{
		"Terminals":	["α"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Productions":	{
			"S":	[["α"], ["α"]]
		}
	}`

	_, err := NewParser(config)

	if _, ok := err.(*grammarAmbiguousError); !ok {
		t.Error(err)
	}
}

func TestParsingTableEntryNotFoundError(t *testing.T) {
	config := `
	{
		"Terminals":	["(", ")"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["(", "S", ")", "S"], ["ε"]]
		}
	}`

	if p, err := NewParser(config); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("(")
		if _, ok := err.(*parserUnmatchedSymbolsError); !ok {
			t.Error(err)
		}
	}
}

func TestParserUnmatchedSymbolsError(t *testing.T) {
	config := `
	{
		"Terminals":	["α"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["α"]]
		}
	}`

	if p, err := NewParser(config); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("β")
		if _, ok := err.(*parsingTableEntryNotFoundError); !ok {
			t.Error(err)
		}
	}
}

func TestBracketsSequence(t *testing.T) {
	config := `
	{
		"Terminals":	["(", ")"],
		"Nonterminals": ["S"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["(", "S", ")", "S"], ["ε"]]
		}
	}`

	if p, err := NewParser(config); err != nil {
		t.Error(err)
	} else {
		bs := NewBracketSequence(5)

		for {
			err = p.Parse(bs.String())
			if err != nil {
				t.Error(err)
				break
			}

			if !bs.Next() {
				break
			}
		}
	}
}

func TestArithmetic(t *testing.T) {
	config := `
	{
		"Terminals":	["+", "-", "*", "/", "^", "(", ")", "α", "β", "γ"],
		"Nonterminals":	["E", "E'", "T", "T'", "F", "F'", "R"],
		"Start":		"E",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"E":	[["T", "E'"]],
			"E'":	[["+", "T", "E'"], ["-", "T", "E'"], ["ε"]],
			"T":	[["F", "T'"]],
			"T'":	[["*", "F", "T'"], ["/", "F", "T'"], ["ε"]],
			"F":	[["R", "F'"]],
			"F'":	[["^", "R", "F'"], ["ε"]],
			"R":	[["(", "E", ")"], ["α"], ["β"], ["γ"]]
		}
	}`

	if p, err := NewParser(config); err != nil {
		t.Error(err)
	} else {
		err = p.Parse("((α+β)*(α-β)/(α^β+β^γ+γ^α))*(α*β*γ)")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestABBA(t *testing.T) {
	config := `
	{
		"Terminals":	["a", "b", "c", "d"],
		"Nonterminals":	["S", "A", "B", "C", "D"],
		"Start":		"S",
		"Dollar":		"$",
		"Epsilon":		"ε",
		"Productions":	{
			"S":	[["A", "a", "A", "b"], ["B", "b", "B", "a"]],
			"A":	[["C", "d"], ["ε"]],
			"B":	[["D", "c"], ["ε"]],
			"C":	[["ε"]],
			"D":	[["ε"]]
		}
	}`

	if p, err := NewParser(config); err != nil {
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
