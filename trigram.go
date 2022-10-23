package langdet

import (
	"sort"
	"unicode"
	"unicode/utf8"
)

// Trigram is a tuple of three unicode runes.
type Trigram [3]rune

// MarshalText implements encoding.TextMarshaler.
func (t Trigram) MarshalText() ([]byte, error) {
	n0 := utf8.RuneLen(t[0])
	n1 := utf8.RuneLen(t[1])
	n2 := utf8.RuneLen(t[2])
	b := make([]byte, n0+n1+n2)
	utf8.EncodeRune(b, t[0])
	utf8.EncodeRune(b[n0:], t[1])
	utf8.EncodeRune(b[n0+n1:], t[2])
	return b, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Trigram) UnmarshalText(b []byte) error {
	var n int
	t[0], n = utf8.DecodeRune(b)
	b = b[n:]
	t[1], n = utf8.DecodeRune(b)
	b = b[n:]
	t[2], _ = utf8.DecodeRune(b)
	return nil
}

// String implements fmt.Stringer.
func (t Trigram) String() string {
	b, _ := t.MarshalText()
	return string(b)
}

func (t *Trigram) shift(r rune) {
	t[0] = t[1]
	t[1] = t[2]
	t[2] = r
}

func trigramLess(a, b Trigram) bool {
	if a[0] == b[0] {
		if a[1] == b[1] {
			return a[2] < b[2]
		}
		return a[1] < b[1]
	}
	return a[0] < b[0]
}

func countTrigrams(s string, counts map[Trigram]int) {
	var prev rune
	t := Trigram{'_', '_', '_'}

	for _, r := range s {
		switch {
		case unicode.IsPunct(r):
			if r == '.' {
				r = '_'
			}
		case unicode.IsSpace(r):
			r = '_'
		case unicode.IsLetter(r):
			if prev == '_' {
				t.shift('_')
				counts[t]++
			}

			r = unicode.ToLower(r)
			t.shift(r)
			counts[t]++
		}

		prev = r
	}

	t.shift('_')
	counts[t]++
}

type trigramSorter struct {
	counts map[Trigram]int
	ranked []Trigram
}

func (ts trigramSorter) Len() int {
	return len(ts.ranked)
}

func (ts trigramSorter) Less(i, j int) bool {
	ti := ts.ranked[i]
	tj := ts.ranked[j]
	ci := ts.counts[ti]
	cj := ts.counts[tj]

	// sort equal counts lexicographically
	if ci == cj {
		return trigramLess(ti, tj)
	}

	// sort unequal counts descending
	return cj < ci
}

func (ts trigramSorter) Swap(i, j int) {
	ts.ranked[i], ts.ranked[j] = ts.ranked[j], ts.ranked[i]
}

func rankTrigrams(counts map[Trigram]int) []Trigram {
	sorter := trigramSorter{
		counts: counts,
		ranked: make([]Trigram, 0, len(counts)),
	}

	for k := range counts {
		sorter.ranked = append(sorter.ranked, k)
	}

	sort.Sort(sorter)

	return sorter.ranked
}

func distance(language []Trigram, document map[Trigram]int) int {
	var dist int
	for rank, tg := range language {
		if docrank, ok := document[tg]; ok {
			dist += abs(rank - min(docrank, len(language)))
		} else {
			dist += len(language)
		}
	}
	return dist
}

func confidence(language []Trigram, dist int) float64 {
	n := len(language)
	return float64(n*n-dist) / float64(n*n)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Train counts all trigrams in s and orders them by frequency.
func Train(s string) []Trigram {
	n := utf8.RuneCountInString(s)
	trigrams := make(map[Trigram]int, n/2)
	countTrigrams(s, trigrams)
	return rankTrigrams(trigrams)
}
