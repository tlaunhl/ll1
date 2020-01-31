package ll1

type symbolSetMap map[string]symbolSet

func (ssm symbolSetMap) insert(k, v string) {
	if _, ok := ssm[k]; !ok {
		ssm[k] = make(symbolSet)
	}

	ssm[k].insert(v)
}

func (ssm symbolSetMap) insertSet(k string, ss symbolSet) {
	if _, ok := ssm[k]; !ok {
		ssm[k] = make(symbolSet)
	}

	for s := range ss {
		ssm[k].insert(s)
	}
}

func (ssm symbolSetMap) get(k string) symbolSet {
	ss, ok := ssm[k]

	if !ok {
		return nil
	}

	return ss
}
