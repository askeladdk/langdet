package langdet

import (
	"sort"
	"unicode"

	"golang.org/x/text/language"
)

type Language struct {
	Tag      language.Tag
	Trigrams []Trigram
}

type Languages struct {
	DefaultTag language.Tag
	Languages  []Language
}

type Options struct {
	Scripts   []*unicode.RangeTable
	Languages map[*unicode.RangeTable]Languages
}

type Result struct {
	Tag        language.Tag
	Confidence float64
}

type resultSorter []Result

func (rs resultSorter) Len() int {
	return len(rs)
}

func (rs resultSorter) Less(i, j int) bool {
	return rs[j].Confidence < rs[i].Confidence
}

func (rs resultSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func DetectLanguageWithOptions(b []byte, options Options) []Result {
	// detect the script
	script := DetectScript(b, options.Scripts)
	if script == nil {
		return []Result{{}}
	}

	// get the language set
	langs, ok := options.Languages[script]
	if !ok {
		return []Result{{}}
	}

	// some languages are defined by their script
	if len(langs.Languages) == 0 {
		return []Result{{Tag: langs.DefaultTag, Confidence: 1}}
	}

	// build and rank the trigrams of the document
	doctri := make(map[Trigram]int, len(b)/2)
	countTrigrams(b, doctri)
	ranked := rankTrigrams(doctri)
	for i, r := range ranked { // map the rankings
		doctri[r] = i
	}

	// compute the distance to each language in the set
	res := make([]Result, len(langs.Languages))
	for i, lang := range langs.Languages {
		dist := distance(lang.Trigrams, doctri)
		res[i].Confidence = confidence(lang.Trigrams, dist)
		res[i].Tag = lang.Tag
	}

	sort.Sort(resultSorter(res))

	return res
}

func DetectLanguage(b []byte) []Result {
	return DetectLanguageWithOptions(b, DefaultOptions)
}

func Train(b []byte) []Trigram {
	trigrams := make(map[Trigram]int, 300)
	countTrigrams(b, trigrams)
	return rankTrigrams(trigrams)
}
