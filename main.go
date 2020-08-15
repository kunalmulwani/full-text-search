package main

import (
  "fmt"
)

func main() {
  filePath := "/home/kunal/projects/full-text-search/enwiki-latest-abstract1.xml"
  docs, _ := loadDocuments(filePath)
  fmt.Println("Documents loaded")
  fmt.Println("Corupus length", len(docs))
  /*
  * Testing part 
  catDocs := search(docs, "Caterpillar")
  fmt.Println("search completed. Results size", len(catDocs))
  for _, doc := range catDocs {
    fmt.Println(doc.Title)
    if doc.Title == "Wikipedia: Caterpillar" {
      fmt.Println("Found")
      break
    }
  }
  
  fmt.Println(analyze("A donut on a glass plate. Only the donuts."))
  */
  idx := make(index)
  idx.add(docs)
  res := idx.search("Small wild cat")
  for _, id := range res {
    fmt.Println(docs[id].Title)
  }
}
