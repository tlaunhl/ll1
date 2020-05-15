package tables

import (
	"github.com/tlaunhl/ll1/internal/errors"
	"github.com/tlaunhl/ll1/internal/symbols"
)

type FirstFollowTable struct {
	Firsts    symbols.SymbolSetMap
	Follows   symbols.SymbolSetMap
	Nullables symbols.SymbolSet
}

func NewFirstFollowTable(
	symt *SymbolTable,
	prdt ProductionTable,
	idxt IndexTable) (*FirstFollowTable, error) {

	builder := newFirstFollowBuilder(symt, prdt, idxt)
	return builder.build()
}

type firstFollowBuilder struct {
	symt      *SymbolTable
	prdt      ProductionTable
	idxt      IndexTable
	firsts    symbols.SymbolSetMap
	follows   symbols.SymbolSetMap
	nullables symbols.SymbolSet
}

func newFirstFollowBuilder(
	symt *SymbolTable,
	prdt ProductionTable,
	idxt IndexTable) *firstFollowBuilder {

	return &firstFollowBuilder{
		symt:      symt,
		prdt:      prdt,
		idxt:      idxt,
		firsts:    make(symbols.SymbolSetMap),
		follows:   make(symbols.SymbolSetMap),
		nullables: make(symbols.SymbolSet),
	}
}

func (b *firstFollowBuilder) build() (*FirstFollowTable, error) {
	// Firsts and Nullables
	if err := b.symt.Nonterminals.Apply(func(nt string) error {
		visited := make(symbols.SymbolSet)

		isNullable, err := b.buildFirsts(nt, visited)
		if err != nil {
			return err
		}

		if isNullable {
			b.nullables.Insert(nt)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	// Follows
	b.follows.Insert(b.symt.Ref.Start(), b.symt.Ref.Dollar())

	if err := b.symt.Nonterminals.Apply(func(nt string) error {
		return b.buildFollows(nt)
	}); err != nil {
		return nil, err
	}

	return &FirstFollowTable{
		Firsts:    b.firsts,
		Follows:   b.follows,
		Nullables: b.nullables,
	}, nil
}

func (b *firstFollowBuilder) buildFirsts(nt string, visited symbols.SymbolSet) (bool, error) {
	visited.Insert(nt)
	var isNullable bool

	if err := b.prdt.Apply(nt, func(pb ProductionBody) error {
		it := pb.Begin()

		for !it.IsEnd() {
			symbol := it.Get()

			if visited.Contains(symbol) {
				return &errors.LeftRecursiveGrammarError{Nonterminal: symbol}
			}

			// Epsilon
			if b.symt.Ref.IsEpsilon(symbol) {
				isNullable = true
				break
			}

			// Terminal
			if b.symt.Terminals.Contains(symbol) {
				b.firsts.Insert(nt, symbol)
				break
			}

			// Nonterminal
			isNullableRec, err := b.buildFirsts(symbol, visited)
			if err != nil {
				return err
			}

			b.firsts.ApplyNoError(symbol, func(first string) {
				b.firsts.Insert(nt, first)
			})

			if !isNullableRec {
				break
			}

			it.Inc()
		}

		if it.IsEnd() {
			isNullable = true
		}

		return nil
	}); err != nil {
		return false, err
	}

	return isNullable, nil
}

func (b *firstFollowBuilder) buildFollows(nt string) error {
	if err := b.idxt.Apply(nt, func(head string) error {
		if err := b.prdt.Apply(head, func(pb ProductionBody) error {
			var isLast bool

			it := pb.FindForwardFirst(nt)

			for !it.IsEnd() {
				it.Inc()
				if it.IsEnd() {
					isLast = true
					break
				}

				symbol := it.Get()

				// Terminal
				if b.symt.Terminals.Contains(symbol) {
					b.follows.Insert(nt, symbol)
					pb.FindForwardNext(it, nt)
					continue
				}

				// Nonterminal
				b.firsts.ApplyNoError(symbol, func(first string) {
					b.follows.Insert(nt, first)
				})

				// Nullable nonterminal
				if b.nullables.Contains(symbol) {
					continue
				}

				break
			}

			if isLast && head != nt {
				b.buildFollows(head)
				b.follows.ApplyNoError(head, func(follow string) {
					b.follows.Insert(nt, follow)
				})
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
