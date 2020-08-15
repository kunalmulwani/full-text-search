package main

import (
  "fmt"
)

func main() {
  filePath := "enwiki-latest-abstract1.xml"
  docs, _ := loadDocuments(filePath)
  fmt.Println("Documents loaded")
  fmt.Println("Corupus length", len(docs))
  idx := make(index)
  idx.add(docs)
  res := idx.search("Small wild cat")
  for _, id := range res {
    fmt.Println(docs[id].Title)
  }
}
