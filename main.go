package main

import (
	utils "Text_Search_Engine/utils"
	"flag"
	"log"
	"time"
)

func main() {
	println("Hello, World!")

	var filePath, query string
	flag.StringVar(&filePath, "file", "", "Path to the file")
	flag.StringVar(&query, "query", "", "Search query")
	flag.Parse()

	log.Println("Searching for:", query, "in file:", filePath)
	start := time.Now()
	docs, err := utils.LoadDocuments(filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexing completed in %s", time.Since(start))
	start = time.Now()
	matched := idx.Search(query)
	log.Printf("Found %d documents matching query in %s", len(matched), time.Since(start))
	for _, id := range matched {
		doc := docs[id]
		log.Printf("Document ID: %d, Content: %s", id, doc)
	}
}
