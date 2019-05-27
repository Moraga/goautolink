package main

import (
	"fmt"
	"regexp"
	"strings"
)

var reEntities = regexp.MustCompile(`([0-9]+(?:[,.-:][0-9]+|[dhms%])*|[a-z]+[a-z0-9$]*(?:[-@.][0-9a-z]+)*)|([.!?]+)`)

func splitWords(s string) (words []string, index [][]int) {
	index = reEntities.FindAllStringIndex(s, -1)
	for i, match := range index {
		words = append(words, s[match[0]:match[1]])
		// size instead position
		index[i][1] = match[1] - match[0]
	}
	return words, index
}

func sumInt(input ...int) (sum int) {
	for _, item := range input {
		sum += item
	}
	return sum
}

func sumFloat32(list ...float32) (sum float32) {
	for _, item := range list {
		sum += item
	}
	return sum
}

func sumFloat64(input ...float64) (total float64) {
	for _, v := range input {
		total += v
	}
	return total
}

func max(input ...int) (max int) {
	for _, item := range input {
		if item > max {
			max = item
		}
	}
	return max
}

func maxFloat32(input ...float32) (max float32) {
	for _, item := range input {
		if item > max {
			max = item
		}
	}
	return max
}

func maxFloat64(input ...float64) (max float64) {
	for _, item := range input {
		if item > max {
			max = item
		}
	}
	return max
}

func min(input ...int) (min int) {
	for i, item := range input {
		if i == 0 || item < min {
			min = item
		}
	}
	return min
}

func minFloat32(input ...float32) (min float32) {
	for i, item := range input {
		if i == 0 || item < min {
			min = item
		}
	}
	return min
}

func inArray(item int, list []int) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}

func removeSliceDupes(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, exists := seen[v]; exists || v == "" {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func arrayIntToString(a []int, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delimiter, -1), "[]")
}

func indexesFromSliceString(s []string) map[string][]int {
	m := make(map[string][]int, len(s))
	for i, v := range s {
		if _, exists := m[v]; exists {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}
	return m
}

func isOneOf(c byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
