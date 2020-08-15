package main

import (
  "encoding/xml"
  "os"
  "regexp"
  "strings"
  "unicode"
)

type Document struct {
    Title string `xml:"title"`
    URL string `xml:"url"`
    Text string `xml:"abstract"`
    Id int
}

func loadDocuments(path string) ([]Document, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer f.Close()

    decoder := xml.NewDecoder(f)
    // creating anonymous struct and its object
    dump := struct{
      Documents []Document `xml:"doc"`
    }{}
    if err := decoder.Decode(&dump); err != nil {
      return nil, err
    }

    docs := dump.Documents
    // assigning Id to each of the created document
    for i := range docs {
      docs[i].Id = i
    }

    return docs, nil
}

func search(docs []Document, term string) []Document {
  var result []Document
  // regExStr := fmt.Sprintf("(?i)\\b%s\\b", term)
  // fmt.Println("Term:", regExStr)
  // regEx := regexp.MustCompile(`(?i)\b` + term + `\b`)
  regEx := regexp.MustCompile(`(?i)\b` + term + `\b`)
  for _, doc := range docs {
    // if strings.Contains(doc.Text, term) {
    if regEx.MatchString(doc.Text) || regEx.MatchString(doc.Title) {
      result = append(result, doc)
    }
  }
  return result
}

func tokenize(text string) []string {
  // tokenize given text into set of words.
  // FieldsFunc is used to split texts on non-alphanumeric characters
  fn := func(c rune) bool {
    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
  }
  return strings.FieldsFunc(text, fn)
}

func analyze(text string) []string {
  tokens := tokenize(text)
  tokens = lowercaseFilter(tokens)
  tokens = stopwordFilter(tokens)
  tokens = stemmerFilter(tokens)
  return tokens
}

