package main

type index map[string][]int

func (idx index) add(docs []Document) {
  for _, doc := range docs {
    for _, token := range analyze(doc.Text) {
      ids := idx[token]
      if ids != nil && ids[len(ids)-1] == doc.Id {
        continue
      }
      idx[token] = append(ids, doc.Id)
    }
  }
}

func (idx index) search(text string) []int {
  var r []int
  for _, token := range analyze(text) {
    if ids, ok := idx[token]; ok {
      if r == nil {
        r = ids
      } else {
        r = intersection(r, ids)
      }
    } else {
      return nil
    }
  }
  return r
}

func intersection(a []int, b []int) []int {
  maxLen := len(a)
  if len(b) > maxLen {
    maxLen = len(b)
  }
  result := make([]int, 0, maxLen)
  var i, j int
  for i < len(a) && j <len(b) {
    if a[i] < b[j] {
      i++
    } else if a[i] > b[j] {
      j++
    } else {
      result = append(result, a[i])
      i++
      j++
    }
  }
  return result
}

