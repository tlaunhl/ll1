package errors

import "fmt"

type EpsilonProductionError struct {
	Nonterminal string
}

func (e *EpsilonProductionError) Error() string {
	return fmt.Sprintf("production [%s] contains epsilon", e.Nonterminal)
}

type ProductionNotFoundError struct {
	Nonterminal string
}

func (e *ProductionNotFoundError) Error() string {
	return fmt.Sprintf("production [%s] not found", e.Nonterminal)
}

type LeftRecursiveGrammarError struct {
	Nonterminal string
}

func (e *LeftRecursiveGrammarError) Error() string {
	return fmt.Sprintf("grammar is left-recursive in production [%s]", e.Nonterminal)
}

type AmbiguousGrammarError struct {
	Nonterminal string
	Terminal    string
}

func (e *AmbiguousGrammarError) Error() string {
	return fmt.Sprintf("grammar is ambiguous: entry [%s][%s] exists", e.Nonterminal, e.Terminal)
}

type UnmatchedSymbolsError struct {
	StackSymbol string
	InputSymbol string
}

func (e *UnmatchedSymbolsError) Error() string {
	return fmt.Sprintf("umatched symbols: %s != %s", e.StackSymbol, e.InputSymbol)
}

type ParsingTableEntryError struct {
	Nonterminal string
	Terminal    string
}

func (e *ParsingTableEntryError) Error() string {
	return fmt.Sprintf("parsing table entry [%s, %s] not found", e.Nonterminal, e.Terminal)
}

type UnparsedSymbolsError struct{}

func (e *UnparsedSymbolsError) Error() string {
	return "remained unparsed symbols in input"
}

type InvalidInputError struct{}

func (e *InvalidInputError) Error() string {
	return "invalid input"
}
