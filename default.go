package langdet

import (
	"unicode"

	"golang.org/x/text/language"
)

var (
	BelarusianTag       = language.Make("be")
	BosnianTag          = language.Make("bs")
	IrishTag            = language.Make("ga")
	LatinTag            = language.Make("la")
	LuxembourgishTag    = language.Make("lb")
	MalteseTag          = language.Make("mt")
	NorwegianBokmalTag  = language.Make("nb")
	NorwegianNyNorskTag = language.Make("nn")
)

var DefaultOptions = Options{
	Languages: map[*unicode.RangeTable]Languages{
		unicode.Armenian: {
			DefaultTag: language.Armenian,
		},
		unicode.Cyrillic: {
			Languages: []Language{
				Belarusian,
				Bulgarian,
				Macedonian,
				Russian,
				Serbian,
				Ukrainian,
			},
		},
		unicode.Latin: {
			Languages: []Language{
				Albanian,
				Bosnian,
				Croatian,
				Czech,
				Danish,
				Dutch,
				English,
				Estonian,
				Finnish,
				French,
				German,
				Hungarian,
				Icelandic,
				Irish,
				Italian,
				Latin,
				Latvian,
				Lithuanian,
				Luxembourgish,
				Maltese,
				NorwegianBokmal,
				NorwegianNyNorsk,
				Polish,
				Portuguese,
				Romanian,
				Spanish,
				Slovak,
				Slovenian,
				Swedish,
				Turkish,
			},
		},
		unicode.Georgian: {
			DefaultTag: language.Georgian,
		},
		unicode.Greek: {
			DefaultTag: language.Greek,
		},
	},
	// https://www.worldatlas.com/articles/the-world-s-most-popular-writing-scripts.html
	Scripts: []*unicode.RangeTable{
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
	},
}
