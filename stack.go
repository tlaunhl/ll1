package ll1

type stringStack struct {
	data []string
}

func (ss *stringStack) push(s string) {
	ss.data = append(ss.data, s)
}

func (ss *stringStack) pop() {
	if len(ss.data) == 0 {
		panic("stringStack: calling pop() on empty stack")
	}

	ss.data = ss.data[:len(ss.data)-1]
}

func (ss *stringStack) top() string {
	if len(ss.data) == 0 {
		panic("stringStack: calling top() on empty stack")
	}

	return ss.data[len(ss.data)-1]
}

func (ss *stringStack) isEmpty() bool {
	return len(ss.data) == 0
}
