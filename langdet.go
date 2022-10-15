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

func DetectLanguageOptions(b []byte, scripts []*unicode.RangeTable,
	langset map[*unicode.RangeTable]Languages) []Result {
	// detect the script
	script := DetectScript(b, scripts)
	if script == nil {
		return nil
	}

	// get the language set
	langs, ok := langset[script]
	if !ok {
		return nil
	}

	// some languages are defined by their script
	if len(langs.Languages) == 0 {
		return []Result{
			{
				Tag:        langs.DefaultTag,
				Confidence: 1,
			},
		}
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
	return DetectLanguageOptions(b, DefaultScripts, DefaultLanguages)
}

func Train(b []byte) []Trigram {
	trigrams := make(map[Trigram]int, 300)
	countTrigrams(b, trigrams)
	return rankTrigrams(trigrams)
}
