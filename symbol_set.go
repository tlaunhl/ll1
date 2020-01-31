package ll1

type symbolSet map[string]bool

func (ss symbolSet) insert(s string) {
	ss[s] = true
}

func (ss symbolSet) contains(s string) bool {
	_, ok := ss[s]
	return ok
}
