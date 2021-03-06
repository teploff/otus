package analyzer

import "fmt"

// Pair stores the word and its number of occurrences in the text.
type Pair struct {
	Word  string
	Count int
}

func (p Pair) String() string {
	return fmt.Sprintf("Word <%s> occurs <%d> times", p.Word, p.Count)
}

// Pairs are the sequence of Pair
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

// Append Pair. If the word is not new - increases filed "Count"; if word is unique create new Pair with word and
// count = 1.
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
