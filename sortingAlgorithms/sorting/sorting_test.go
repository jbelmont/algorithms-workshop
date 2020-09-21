package sorting

import (
	"sort"
	"testing"
)

func TestSorting(t *testing.T) {
	soldiers := []Soldier{
		Soldier{
			Name:           "John Smith",
			Age:            18,
			Rank:           "PVT",
			YearsOfService: 1,
		},
		Soldier{
			Name:           "Jane Carter",
			Age:            21,
			Rank:           "CPT",
			YearsOfService: 4,
		},
		Soldier{
			Name:           "Billy Bob",
			Age:            22,
			Rank:           "SGT",
			YearsOfService: 3,
		},
		Soldier{
			Name:           "Daniel LaRusso",
			Age:            24,
			Rank:           "SSG",
			YearsOfService: 6,
		},
	}

	sort.Sort(SortByYearsOfService(soldiers))
	if soldiers[0].Name != "John Smith" {
		t.Error("Should be sorted by the YearsOfService starting with least experienced")
	}

	sort.Slice(soldiers, func(i, j int) bool {
		return soldiers[i].Age > soldiers[j].Age
	})

	if soldiers[0].Name != "Daniel LaRusso" {
		t.Error("Should be sorted by the oldest soldier")
	}
}
