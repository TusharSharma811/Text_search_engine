package utils

import "strings"

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
	})
}

func analyze(text string) []string {
	// Simple whitespace tokenizer, can be improved with more complex logic
	token := tokenize(text)
	token = lowercaseFilter(token)
	token = stopwordfilter(token)
	token = stemmerFilter(token)
	return token

}
