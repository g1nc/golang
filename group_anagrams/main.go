package main

import (
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	dict := make(map[string][]string)
	for _, w := range strs {
		alpha := strings.Split(w, "")
		sort.Strings(alpha)
		dict[strings.Join(alpha, "")] = append(dict[strings.Join(alpha, "")], w)
	}
	mapValues(dict)
}

func mapValues(d map[string][]string) [][]string {
	ans := [][]string{}
	for _, value := range d {
		ans = append(ans, value)
	}
	return ans
}
