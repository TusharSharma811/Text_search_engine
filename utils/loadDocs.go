package utils

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"os"
)

type Document struct {
	ID      int    `xml:"id,attr"`
	Title   string `xml:"title"`
	URL     string `xml:"url"`
	CONTENT string `xml:"abstract"`
}

func LoadDocuments(filePath string) ([]Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var docs []Document
	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)

	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil && err != io.EOF {
		return nil, err
	}
	docs = dump.Documents

	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
