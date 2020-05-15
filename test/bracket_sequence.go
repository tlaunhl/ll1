package test

type BracketSequence struct {
	bytes []byte
}

func NewBracketSequence(size int) BracketSequence {
	bs := BracketSequence{make([]byte, 2*size)}

	for cnt := 0; cnt < size; cnt++ {
		bs.bytes[cnt] = '('
		bs.bytes[cnt+size] = ')'
	}

	return bs
}

func (bs BracketSequence) ForEach(fnc func(seq string) error) error {
	for {
		if err := fnc(string(bs.bytes)); err != nil {
			return err
		}

		if !bs.next() {
			break
		}
	}

	return nil
}

func (bs BracketSequence) next() bool {
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
