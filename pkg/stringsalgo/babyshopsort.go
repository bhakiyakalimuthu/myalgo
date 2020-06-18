package stringsalgo

import (
	"sort"
)

type option struct {
	list       []string
	booster    string
	mulbooster []string
}

type sortBy option

func (s sortBy) Len() int { return len(s.list) }

func (s sortBy) Swap(i, j int) {

	s.list[i], s.list[j] = s.list[j], s.list[i]

}

func (s sortBy) Less(i, j int) bool {

	return len(s.list[i]) < len(s.list[j])
}

func BabySorting(s []string, booster string, mulbooster []string) option {
	o := option{
		list:       s,
		booster:    "",
		mulbooster: nil,
	}
	sort.Sort(sortBy(o))
	return o
}
