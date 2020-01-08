package analyzer

import "fmt"

type Pair struct {
	Word  string
	Count int
}

func (p Pair) String() string {
	return fmt.Sprintf("Word <%s> occurs <%d> times", p.Word, p.Count)
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].Count < p[j].Count
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pairs) Append(word string) {
	for index, pair := range *p {
		if pair.Word == word {
			(*p)[index].Count++
			return
		}
	}

	*p = append(*p, Pair{
		Word:  word,
		Count: 1,
	})
}

func (p *Pairs) DeleteUniqueWords() {
	uniqueWordExist := true

	for uniqueWordExist {
		uniqueWordExist = false
		for index, pair := range *p {
			if pair.Count == 1 {
				*p = append((*p)[:index], (*p)[index+1:]...)
				uniqueWordExist = true
				break
			}
		}
	}
}
