package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	visited := map[string]int{}

	for _, v := range s {
		ii, letterExists := visited[string(v)]

		if letterExists == true {
			visited[string(v)] = ii + 1
			continue
		}

		visited[string(v)] = 1
	}

	for i, v := range s {
		if visited[string(v)] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
