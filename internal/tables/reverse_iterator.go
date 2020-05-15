package tables

type ReverseIterator struct {
	pb    ProductionBody
	index int
	end   int
}

func NewReverseIterator(pb ProductionBody) *ReverseIterator {
	return &ReverseIterator{pb, len(pb) - 1, -1}
}

func (it *ReverseIterator) IsEnd() bool {
	return it.index == it.end
}

func (it *ReverseIterator) Dec() {
	if it.index == it.end {
		panic("ReverseIterator: decrementing at end")
	}

	it.index--
}

func (it *ReverseIterator) Get() string {
	if it.index == it.end {
		panic("ReverseIterator: dereferencing at end")
	}

	return it.pb[it.index]
}
