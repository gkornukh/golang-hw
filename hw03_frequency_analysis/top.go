package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const MaxTopSize = 10

const allowedPunct = "!\"#$%&'()*+,./:;<=>?@[\\]^_`{|}~"

type wordCount struct {
	word  string
	count int
}

func Top10(str string) []string {
	words := cleanInput(strings.Fields(str))
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	counts := make([]wordCount, 0, len(frequency))
	for word, count := range frequency {
		counts = append(counts, wordCount{word, count})
	}
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})
	top := make([]string, 0, MaxTopSize)
	for i := 0; i < len(counts) && i < MaxTopSize; i++ {
		top = append(top, counts[i].word)
	}
	return top
}

func cleanInput(words []string) []string {
	result := make([]string, 0, len(words))
	for _, word := range words {
		runes := []rune(word)
		start := 0
		end := len(runes) - 1
		isEdgePunctuation := func(r rune) bool {
			return strings.ContainsRune(allowedPunct, r)
		}
		for ; start < len(runes); start++ {
			if !isEdgePunctuation(runes[start]) {
				break
			}
		}
		for ; end >= start; end-- {
			if !isEdgePunctuation(runes[end]) {
				break
			}
		}
		cleaned := string(runes[start : end+1])
		cleaned = strings.ToLower(cleaned)
		if cleaned != "" && cleaned != "-" {
			result = append(result, cleaned)
		}
	}
	return result
}
