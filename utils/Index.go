package utils

type Index map[string][]int

func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		words := tokenize(doc.Title + " " + doc.CONTENT)
		for _, word := range words {
			idx[word] = append(idx[word], doc.ID)
		}
	}
}
func (idx Index) Search(query string) []int {
	words := tokenize(query)
	var results []int
	for _, word := range words {
		results = append(results, idx[word]...)
	}
	return results
}
