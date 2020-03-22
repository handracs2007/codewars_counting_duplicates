package main

import "strings"

func duplicate_count(s1 string) int {
	var duplicateCount = 0
	var charCountMap = make(map[rune]int)

	for _, chr := range strings.ToLower(s1) {
		charCountMap[chr] = charCountMap[chr] + 1
	}

	for _, val := range charCountMap {
		if val > 1 {
			duplicateCount++
		}
	}

	return duplicateCount
}

func main() {

}
