package ll1

import "fmt"

type productionEpsilonError struct {
	head string
}

func (e *productionEpsilonError) Error() string {
	return fmt.Sprintf("production [%s] contains epsilon", e.head)
}

type productionNotFoundError struct {
	head string
}

func (e *productionNotFoundError) Error() string {
	return fmt.Sprintf("production [%s] not found", e.head)
}

type grammarLeftRecursiveError struct {
	nt string
}

func (e *grammarLeftRecursiveError) Error() string {
	return fmt.Sprintf("grammar is left-recursive in production [%s]", e.nt)
}

type grammarAmbiguousError struct{}

func (e *grammarAmbiguousError) Error() string {
	return "grammar is ambiguous"
}

type parserUnmatchedSymbolsError struct {
	st string
	ss string
}

func (e *parserUnmatchedSymbolsError) Error() string {
	return fmt.Sprintf("umatched symbols: %s != %s", e.st, e.ss)
}

type parsingTableEntryNotFoundError struct {
	nt string
	t  string
}

func (e *parsingTableEntryNotFoundError) Error() string {
	return fmt.Sprintf("parsing table entry [%s, %s] not found", e.nt, e.t)
}

type stackNotEmptyError struct{}

func (e *stackNotEmptyError) Error() string {
	return "stack not empty"
}

type unparsedSymbolsError struct{}

func (e *unparsedSymbolsError) Error() string {
	return "remained unparsed symbols in input"
}
