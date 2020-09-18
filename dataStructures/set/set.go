package set

// Set Class
type Set struct {
	IMap map[int]bool
}

// New is a set method that creates an integer map
func (s *Set) New() {
	s.IMap = make(map[int]bool)
}

// Add is a set method that adds element to set.
func (s *Set) Add(number int) {
	if !s.Contains(number) {
		s.IMap[number] = true
	}
}

// Delete is a set method that deletes element from set.
func (s *Set) Delete(number int) {
	delete(s.IMap, number)
}

// Intersect is a set method that returns the intersection of 2 sets
func (s *Set) Intersect(comparisonSet *Set) *Set {
	intersectionSet := &Set{}
	intersectionSet.New()

	for v := range s.IMap {
		if comparisonSet.Contains(v) {
			intersectionSet.Add(v)
		}
	}
	return intersectionSet
}

// Union is a set method that returns the union of 2 sets
func (s *Set) Union(unionSet *Set) *Set {
	union := &Set{}
	union.New()

	for v := range s.IMap {
		union.Add(v)
	}

	for v := range unionSet.IMap {
		union.Add(v)
	}
	return union
}

// Contains is a set method that checks whether
// an element is a member of the set
func (s *Set) Contains(number int) bool {
	return s.IMap[number]
}
