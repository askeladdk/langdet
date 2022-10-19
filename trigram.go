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

// Train counts all trigrams in b and orders them by frequency.
// Use parameter nTrigramsEstimate to preallocate an expected number of trigrams.
func Train(b []byte, nTrigramsEstimate int) []Trigram {
	trigrams := make(map[Trigram]int, nTrigramsEstimate)
	countTrigrams(b, trigrams)
	return rankTrigrams(trigrams)
}
