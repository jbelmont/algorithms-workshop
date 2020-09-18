package set

import "testing"

func TestSetOperations(t *testing.T) {
	set := &Set{}
	set.New()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	if !set.Contains(2) {
		t.Error("Expected set to contain 2!")
	}

	set.Delete(2)

	if set.Contains(2) {
		t.Error("The element 2 should be deleted")
	}

	set2 := &Set{}
	set2.New()

	set2.Add(1)
	set2.Add(3)

	intersection := set.Intersect(set2)
	if !intersection.Contains(3) {
		t.Error("The element 3 should exist")
	}

	set3 := &Set{}
	set3.New()

	set3.Add(4)
	set3.Add(5)

	set2.Union(set3)
}
