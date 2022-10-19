package langdet

import (
	"unicode"

	"golang.org/x/text/language"
)

// Tags for languages missing from golang.org/x/text/language.
var (
	BelarusianTag       = language.Make("be")
	BosnianTag          = language.Make("bs")
	IrishTag            = language.Make("ga")
	JavaneseTag         = language.Make("jv")
	LatinTag            = language.Make("la")
	LuxembourgishTag    = language.Make("lb")
	MalteseTag          = language.Make("mt")
	MyanmarTag          = language.Make("my")
	NorwegianBokmålTag  = language.Make("nb")
	NorwegianNynorskTag = language.Make("nn")
	OriyaTag            = language.Make("or")
	PunjabiTag          = language.Make("pa")
	SinhaleseTag        = language.Make("si")
	SundaneseTag        = language.Make("su")
	TibetanTag          = language.Make("bo")
)

// DefaultOptions is a default set of options that detects
// the most commonly used languages worldwide.
var DefaultOptions = Options{
	Languages: map[*unicode.RangeTable]Languages{
		unicode.Arabic: {
			DefaultTag: language.Arabic,
		},
		unicode.Armenian: {
			DefaultTag: language.Armenian,
		},
		unicode.Bengali: {
			DefaultTag: language.Bengali,
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
		unicode.Devanagari: {
			DefaultTag: language.Hindi,
		},
		unicode.Ethiopic: {
			DefaultTag: language.Amharic,
		},
		unicode.Javanese: {
			DefaultTag: JavaneseTag,
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
				NorwegianBokmål,
				NorwegianNynorsk,
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
		unicode.Gujarati: {
			DefaultTag: language.Gujarati,
		},
		unicode.Gurmukhi: {
			DefaultTag: PunjabiTag,
		},
		unicode.Han: {
			DefaultTag: language.Chinese,
		},
		unicode.Hangul: {
			DefaultTag: language.Korean,
		},
		unicode.Hebrew: {
			DefaultTag: language.Hebrew,
		},
		HiraganaKatakana: {
			DefaultTag: language.Japanese,
		},
		unicode.Kannada: {
			DefaultTag: language.Kannada,
		},
		unicode.Khmer: {
			DefaultTag: language.Khmer,
		},
		unicode.Lao: {
			DefaultTag: language.Lao,
		},
		unicode.Malayalam: {
			DefaultTag: language.Malayalam,
		},
		unicode.Myanmar: {
			DefaultTag: MyanmarTag,
		},
		unicode.Oriya: {
			DefaultTag: OriyaTag,
		},
		unicode.Sinhala: {
			DefaultTag: SinhaleseTag,
		},
		unicode.Sundanese: {
			DefaultTag: SundaneseTag,
		},
		unicode.Tamil: {
			DefaultTag: language.Tamil,
		},
		unicode.Telugu: {
			DefaultTag: language.Telugu,
		},
		unicode.Thai: {
			DefaultTag: language.Thai,
		},
		unicode.Tibetan: {
			DefaultTag: TibetanTag,
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
		unicode.Tibetan,
		unicode.Georgian,
	},
}
