package main

import (
  "strings"
  snowballeng "github.com/kljensen/snowball/english"
)

var stopwords = map[string]struct{}{
  "a": {}, "and": {}, "be": {}, "have": {}, "i": {},
  "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func lowercaseFilter(tokens []string) []string {
  result := make([]string, len(tokens))
  for i, token := range tokens {
    result[i] = strings.ToLower(token)
  }
  return result
}

func stopwordFilter(tokens []string) []string {
  result := make([]string, 0, len(tokens))
  for _, token := range tokens {
    if _, ok := stopwords[token]; !ok {
      result = append(result, token)
    }
  }
  return result
}

func stemmerFilter(tokens []string) []string {
  result := make([]string, len(tokens))
  for i, token := range tokens {
    result[i] = snowballeng.Stem(token, false)
  }
  return result
}

