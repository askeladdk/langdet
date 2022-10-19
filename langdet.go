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
	Scripts          []*unicode.RangeTable
	Languages        map[*unicode.RangeTable]Languages
	MinConfidence    float64
	MinRelConfidence float64
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

	// return default tag if confidence is too low
	if res[0].Confidence < options.MinConfidence {
		return []Result{{Tag: langs.DefaultTag}}
	}

	// filter languages that have less than the min relative confidence difference
	if len(res) > 1 && options.MinRelConfidence > 0 {
		res2 := res[:1]
		conf := res[0].Confidence
		for _, r := range res[1:] {
			if options.MinRelConfidence < conf-r.Confidence {
				res2 = append(res2, r)
				conf = r.Confidence
			}
		}
		return res2
	}

	return res
}

func DetectLanguage(b []byte) language.Tag {
	return DetectLanguageWithOptions(b, DefaultOptions)[0].Tag
}
