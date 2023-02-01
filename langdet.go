// Package langdet detects natural languages in text.
package langdet

import (
	"sort"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/language"
)

//go:generate go run ./cmd/train/main.go

// Language profiles a natural language.
type Language struct {
	// Tag is the BCP 47 language tag.
	Tag language.Tag

	// Trigrams is the trigrams profile created by Train.
	Trigrams []Trigram
}

// Languages is a set of languages that share the same writing script.
type Languages struct {
	// DefaultTag is the default language tag used if Languages is empty.
	DefaultTag language.Tag

	// Languages is the set of languages sharing the same writing script.
	// If this is empty or nil, the detected language is always DefaultTag.
	Languages []Language
}

// Options configures the language detector.
type Options struct {
	// Scripts is the set of writing scripts to detect.
	Scripts []*unicode.RangeTable

	// Languages maps writing systems to a set of languages.
	Languages map[*unicode.RangeTable]Languages

	// MinConfidence is the minimum confidence that must be met
	// before DetectLanguage returns the detected language.
	MinConfidence float64

	// MinRelConfidence is the minimum confidence difference
	// that must be met between detected languages.
	// Languages that do not meet the minimum are filtered from the result.
	MinRelConfidence float64
}

// Result holds a detected language and confidence.
type Result struct {
	// Tag is the detected language.
	Tag language.Tag

	// Confidence is the probability that this language is correct, between 0 and 1.
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

// DetectLanguageWithOptions detects the language of s configured by options.
// It returns a set of candidate languages ordered by confidence level.
// At least one result is always returned.
func DetectLanguageWithOptions(s string, options Options) []Result {
	// detect the script
	script := DetectScript(s, options.Scripts)
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
	trigrams := make(map[Trigram]int, utf8.RuneCountInString(s)/2)
	countTrigrams(s, trigrams)
	ranked := rankTrigrams(trigrams)
	for i, r := range ranked { // map the rankings
		trigrams[r] = i
	}

	// compute the distance to each language in the set
	res := make([]Result, len(langs.Languages))
	for i, lang := range langs.Languages {
		dist := distance(lang.Trigrams, trigrams)
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

// DetectLanguage is a shorthand that calls DetectLanguageWithOptions
// with the default options and returns the best detected language.
func DetectLanguage(s string) language.Tag {
	return DetectLanguageWithOptions(s, DefaultOptions)[0].Tag
}
