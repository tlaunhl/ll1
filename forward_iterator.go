package ll1

type forwardIterator struct {
	slice []string
	index int
	end   int
}

func newforwardIterator(slice []string) *forwardIterator {
	return &forwardIterator{slice, 0, len(slice)}
}

func (it *forwardIterator) isEnd() bool {
	return it.index == it.end
}

func (it *forwardIterator) inc() {
	if it.index == it.end {
		panic("forwardIterator: incrementing at end")
	}

	it.index++
}

func (it *forwardIterator) get() string {
	if it.index == it.end {
		panic("forwardIterator: dereferencing at end")
	}

	return it.slice[it.index]
}

func findFrom(it *forwardIterator, str string) {
	for it.index < it.end {
		if it.slice[it.index] == str {
			break
		}

		it.index++
	}
}

func find(slice []string, str string) *forwardIterator {
	it := newforwardIterator(slice)
	findFrom(it, str)

	return it
}
