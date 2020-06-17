package stringsalgo

import "strings"

// "single word",
func WordCounter(words string) int {
	counter := 0
	for i := 0; i < len(words); i++ {
		if words[i] == ' ' && i != len(words)-1 {
			counter++
			continue
		}
		if i == len(words)-1 {
			counter++
		}
	}
	return counter
}

func WordCounter1(words string) int {
	counter := 0
	for i := 0; i < len(words); i++ {
		if words[i] == ' ' {
			counter++
			continue
		}
		if i == len(words)-1 {
			counter++
		}
	}
	return counter
}

func WordOccCounter(words string) map[string]int {
	counter := make(map[string]int, 0)
	s := strings.Split(words, " ")
	for _, item := range s {
		if val, ok := counter[item]; ok {
			counter[item] = val + 1
		}
		println(counter[item])
	}
	return counter
}
