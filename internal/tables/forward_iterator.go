package tables

type ForwardIterator struct {
	pb    ProductionBody
	index int
	end   int
}

func NewForwardIterator(pb ProductionBody) *ForwardIterator {
	return &ForwardIterator{pb, 0, len(pb)}
}

func (it *ForwardIterator) IsEnd() bool {
	return it.index == it.end
}

func (it *ForwardIterator) Inc() {
	if it.index == it.end {
		panic("ForwardIterator: incrementing at end")
	}

	it.index++
}

func (it *ForwardIterator) Get() string {
	if it.index == it.end {
		panic("ForwardIterator: dereferencing at end")
	}

	return it.pb[it.index]
}
