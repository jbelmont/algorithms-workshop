package sorting

import (
	"fmt"
)

type Soldier struct {
	Name           string
	Rank           string
	Age            int
	YearsOfService int
}

func (s Soldier) ToString() string {
	return fmt.Sprintf(
		"Name: %s\nRank: %s\nAge: %d\nYearsOfService: %d\n",
		s.Name,
		s.Rank,
		s.Age,
		s.YearsOfService,
	)
}

type SortByYearsOfService []Soldier

func (s SortByYearsOfService) Len() int {
	return len(s)
}

func (s SortByYearsOfService) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByYearsOfService) Less(i int, j int) bool {
	return s[i].YearsOfService < s[j].YearsOfService
}
