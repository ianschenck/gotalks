package generate

// Automatically generated.

// Add elements to the set.
func (lhs StringSet) Add(elements ...string) {
	for _, element := range elements {
		lhs[element] = struct{}{}
	}
}

// Remove elements from the set.
func (lhs StringSet) Remove(elements ...string) {
	for _, element := range elements {
		if _, ok := lhs[element]; ok {
			delete(lhs, element)
		}
	}
}

// Union returns the union of both sets.
func (lhs StringSet) Union(rhs ...StringSet) (unioned StringSet) {
	unioned = make(StringSet)
	for k := range lhs {
		unioned[k] = struct{}{}
	}
	for _, other := range rhs {
		for k := range other {
			unioned[k] = struct{}{}
		}
	}
	return
}

// Intersect returns the intersection of all sets.
func (lhs StringSet) Intersect(rhs ...StringSet) (intersected StringSet) {
	intersected = make(StringSet)
	for k := range lhs {
		ok := len(rhs) > 0
		for _, other := range rhs {
			_, contains := other[k]
			ok = ok && contains
			if !ok {
				break
			}
		}
		if ok {
			intersected[k] = struct{}{}
		}
	}
	return
}

// Equals returns true if the sets are equal.
func (lhs StringSet) Equals(rhs StringSet) bool {
	intersect := lhs.Intersect(rhs)
	union := lhs.Union(rhs)
	for e := range union {
		if _, ok := intersect[e]; !ok {
			return false
		}
	}
	return true
}
