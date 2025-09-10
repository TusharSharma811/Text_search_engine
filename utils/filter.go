package utils

import (
	"strings"

	snowball "github.com/kljensen/snowball"
)

func lowercaseFilter(tokens []string) []string {

	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopwordfilter(tokens []string) []string {
	stopwords := map[string]struct{}{
		"the": {}, "is": {}, "at": {}, "which": {}, "on": {},
	}
	var filtered []string
	for _, token := range tokens {
		if _, found := stopwords[token]; !found {
			filtered = append(filtered, token)
		}
	}
	return filtered
}

func stemmerFilter(tokens []string) []string {
	var filtered []string
	for _, token := range tokens {
		stemmed, err := snowball.Stem(token, "english", true)
		if err != nil {
			stemmed = token // Fallback to original token if stemming fails
		}
		filtered = append(filtered, stemmed)
	}
	return filtered
}
