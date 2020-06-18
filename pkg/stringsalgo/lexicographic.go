package stringsalgo

import "strings"

func smallestString(substrings []string) string {
	// Write your code here
	lexisort(substrings, len(substrings))
	s := ""
	n := len(substrings)
	for i := 0; i < n; i++ {
		s += substrings[i]
	}
	return s
}
func lexisort(a []string, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			str1 := a[i] + a[j]
			str2 := a[j] + a[j]
			x := strings.Compare(str1, str2) > 0
			if x {
				s := a[i]
				a[i] = a[j]
				a[j] = s
			}
		}
	}
}
func main() {
	x := []string{"a", "bc", "ad"}
	s := smallestString(x)
	println(s)
}
