package langdet

// based on https://github.com/abadojack/whatlanggo

import (
	"unicode"
)

// HiraganaKatakana is the unicode set of Japanese characters.
var HiraganaKatakana = &unicode.RangeTable{
	R16: append(unicode.Hiragana.R16, unicode.Katakana.R16...),
	R32: append(unicode.Hiragana.R32, unicode.Katakana.R32...),
}

// DetectScript detects the dominant writing script of s.
func DetectScript(s string, scripts []*unicode.RangeTable) *unicode.RangeTable {
	counts := make([]int, len(scripts))

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		for i := range counts {
			if unicode.Is(scripts[i], r) {
				counts[i]++
			}
		}
	}

	var best int
	var jp bool
	for i := 1; i < len(counts); i++ {
		if counts[i] > counts[best] {
			best = i
			jp = jp || scripts[i] == HiraganaKatakana
		}
	}

	if counts[best] == 0 {
		return nil
	}

	// If Hiragana or Katakana is included, even if judged as Mandarin,
	// it is regarded as Japanese. Japanese uses Kanji (unicode.Han)
	// in addition to Hiragana and Katakana.
	if jp && scripts[best] == unicode.Han {
		return HiraganaKatakana
	}

	return scripts[best]
}
