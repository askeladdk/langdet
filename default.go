package langdet

import (
	"unicode"

	"golang.org/x/text/language"
)

var DefaultLanguages = map[*unicode.RangeTable]Languages{
	unicode.Latin: {
		Languages: []Language{
			Dutch,
			English,
			German,
		},
	},
	unicode.Greek: {
		DefaultTag: language.Greek,
	},
}

// https://www.worldatlas.com/articles/the-world-s-most-popular-writing-scripts.html

var DefaultScripts = []*unicode.RangeTable{
	unicode.Latin,
	unicode.Han,
	unicode.Arabic,
	unicode.Devanagari,
	unicode.Bengali,
	unicode.Cyrillic,
	HiraganaKatakana,
	unicode.Javanese,
	unicode.Hangul,
	unicode.Telugu,
	unicode.Tamil,
	unicode.Gujarati,
	unicode.Kannada,
	unicode.Myanmar,
	unicode.Malayalam,
	unicode.Thai,
	unicode.Sundanese,
	unicode.Gurmukhi,
	unicode.Lao,
	unicode.Oriya,
	unicode.Ethiopic,
	unicode.Sinhala,
	unicode.Hebrew,
	unicode.Armenian,
	unicode.Khmer,
	unicode.Greek,
	// Lontara
	unicode.Tibetan,
	unicode.Georgian,
}
