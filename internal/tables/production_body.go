package tables

type ProductionBody []string

func (pb ProductionBody) Begin() *ForwardIterator {
	return NewForwardIterator(pb)
}

func (pb ProductionBody) FindForwardFirst(symbol string) *ForwardIterator {
	it := NewForwardIterator(pb)

	for !it.IsEnd() && it.Get() != symbol {
		it.Inc()
	}

	return it
}

func (pb ProductionBody) FindForwardNext(it *ForwardIterator, symbol string) *ForwardIterator {
	for !it.IsEnd() && it.Get() != symbol {
		it.Inc()
	}

	return it
}

func (pb ProductionBody) ApplyForward(fnc func(symbol string) error) error {
	for it := NewForwardIterator(pb); !it.IsEnd(); it.Inc() {
		if err := fnc(it.Get()); err != nil {
			return err
		}
	}

	return nil
}

func (pb ProductionBody) ApplyForwardNoError(fnc func(symbol string)) {
	for it := NewForwardIterator(pb); !it.IsEnd(); it.Inc() {
		fnc(it.Get())
	}
}

func (pb ProductionBody) ApplyReverseNoError(fnc func(symbol string)) {
	for it := NewReverseIterator(pb); !it.IsEnd(); it.Dec() {
		fnc(it.Get())
	}
}
