package main

import (
	"fmt"
	"math"
	"sort"
)

type MatchInScore struct {
	Rank       float64 `json:"rank"`
	Some       int     `json:"some"`
	TermsCount int
	Size       int
	Some2      float32
	Some3      int
	Some4      int
	Group      []int
	Slice      [2]int `json:"slice"`
}

type Cluster struct {
	TermsFound   int
	Group        []int
	TermsMatched [][]int
	Hash         string
}

type Match struct {
	Title   string         `json:"title"`
	Stem    []string       `json:"stem"`
	URL     string         `json:"url"`
	Matches []MatchInScore `json:"matches"`
}

func findMatches(text string) (matches []Match) {
	text = sanitizeString(text)
	stemmedText, offset := stemText(text)
	textCount := len(stemmedText)
	sourceMap := indexesFromSliceString(stemmedText)

	fmt.Println(stemmedText)

	for _, row := range db {
		match := matchScore(&row.Ustem, &stemmedText, sourceMap, textCount, &offset)
		if match != nil {
			matches = append(matches, Match{row.Title, row.Stem, row.URL, match})
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Matches[0].Rank > matches[j].Matches[0].Rank
	})
	return matches
}

func matchScore(terms *[]string, source *[]string, sourceMap map[string][]int, textCount int, offset *[][]int) (scores []MatchInScore) {
	var termsPositions [][]int
	var termsWeight []float64
	var positions []int

	indexTermsMap := make(map[int]int)

	for termIndex, term := range *terms {
		termPositions, _ := sourceMap[term]
		termsPositions = append(termsPositions, termPositions)
		positions = append(positions, termPositions...)
		if termWeight, weightExists := vectors[term]; weightExists {
			termsWeight = append(termsWeight, termWeight)
		} else {
			termsWeight = append(termsWeight, 0.)
		}
		for _, termPosition := range termPositions {
			indexTermsMap[termPosition] = termIndex
		}
	}

	if positions == nil {
		return scores
	}

	sort.Ints(positions)

	termsMaxWeight := maxFloat64(termsWeight...)

	for i, v := range termsWeight {
		termsWeight[i] = v / termsMaxWeight
	}

	scores = matchIn(&termsPositions, &termsWeight, &positions, len(termsPositions), textCount, &indexTermsMap, offset)

	return scores
}

func matchIn(termsPositions *[][]int, termsWeight *[]float64, positions *[]int, termsCount int, textCount int, indexTermsMap *map[int]int, offset *[][]int) (scores []MatchInScore) {
	stack := [][]int{*positions}
	var item []int
	hashes := make(map[string]struct{})
	termsBase := math.Log2(sumFloat64(*termsWeight...) + 1)
	termsPosBase := math.Log2(float64(termsCount + 1))

	for len(stack) > 0 {
		item, stack = stack[len(stack)-1], stack[:len(stack)-1]
		clusters := clusterizeAndEvaluates(termsPositions, item, indexTermsMap)
		for _, cluster := range clusters {
			if _, exists := hashes[cluster.Hash]; exists {
				continue
			}
			hashes[cluster.Hash] = struct{}{}

			var termsD float64
			var wPos float64 = 1

			for k, v := range cluster.TermsMatched {
				if v != nil {
					termsD += (*termsWeight)[k] * wPos
				} else {
					wPos = math.Log2(float64(termsCount-k)) / termsPosBase
				}
			}

			termsFoundWeight := math.Log2(termsD+1) / termsBase

			if cluster.TermsFound <= 1 {
				continue
			}

			stack = append(stack, cluster.Group)
			size := len(cluster.Group)
			first := cluster.Group[0]
			last := cluster.Group[size-1]

			bias := math.Log2(float64(cluster.TermsFound) / float64(last-first+1))

			if bias <= -1 {
				continue
			}

			rank := math.Pow(math.Log2(float64(cluster.TermsFound))/math.Log2(float64(termsCount)), 1-bias)
			rank *= termsFoundWeight

			if rank < .3 {
				continue
			}

			slice := [2]int{cluster.Group[0], sumInt((*offset)[cluster.Group[len(cluster.Group)-1]]...)}

			scores = append(scores, MatchInScore{rank, -1, termsCount, size, float32(size) / float32(textCount) * 10, 0, 0, cluster.Group, slice})
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Rank > scores[j].Rank
	})

	return scores
}

func clusterizeAndEvaluates(termsPositions *[][]int, positions []int, indexTermsMap *map[int]int) (clusters []Cluster) {
	for _, group := range clusterizeScalar2(positions...) {
		termsFound := 0
		termsMatched := make([][]int, len(*termsPositions))
		for _, position := range group {
			if v, exists := (*indexTermsMap)[position]; exists {
				termsMatched[v] = append(termsMatched[v], position)
				termsFound++
			}
		}
		clusters = append(clusters, Cluster{termsFound, group, termsMatched, arrayIntToString(group, ",")})
	}
	return clusters
}
