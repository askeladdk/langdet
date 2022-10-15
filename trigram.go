package langdet

import (
	"sort"
	"unicode"
	"unicode/utf8"
)

type Trigram [3]rune

func (t *Trigram) shift(r rune) {
	t[0] = t[1]
	t[1] = t[2]
	t[2] = r
}

func (t Trigram) String() string {
	return string([]rune(t[:]))
}

func trigramLess(a, b Trigram) bool {
	if a[0] < b[0] {
		return true
	} else if a[1] < b[1] {
		return true
	}
	return a[2] < b[2]
}

func countTrigrams(in []byte, freqs map[Trigram]int) {
	var prev rune
	t := Trigram{'_', '_', '_'}

	for {
		r, sz := utf8.DecodeRune(in)
		in = in[sz:]

		if sz == 0 {
			t.shift('_')
			freqs[t]++
			return
		}

		if unicode.IsDigit(r) {
			continue
		}

		if unicode.IsPunct(r) {
			if r == '.' {
				r = '_'
			} else {
				continue
			}
		}

		if unicode.IsSpace(r) {
			r = '_'
		}

		if r != '_' {
			if prev == '_' {
				t.shift('_')
				freqs[t]++
			}

			r = unicode.ToLower(r)
			t.shift(r)
			freqs[t]++
		}

		prev = r
	}
}

type trigramSorter struct {
	freqs  map[Trigram]int
	ranked []Trigram
}

func (ts trigramSorter) Len() int {
	return len(ts.ranked)
}

func (ts trigramSorter) Less(i, j int) bool {
	ti := ts.ranked[i]
	tj := ts.ranked[j]
	ci := ts.freqs[ti]
	cj := ts.freqs[tj]

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

func rankTrigrams(freqs map[Trigram]int) []Trigram {
	sorter := trigramSorter{
		freqs:  freqs,
		ranked: make([]Trigram, 0, len(freqs)),
	}

	for k := range freqs {
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
