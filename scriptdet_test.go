package langdet

import (
	"testing"
	"unicode"
)

func TestDetectScript(t *testing.T) {
	for _, testcase := range []struct {
		Name   string
		Text   string
		Script *unicode.RangeTable
	}{
		{
			Name:   "Garbage",
			Text:   "123456789-=?",
			Script: nil,
		},
		{
			Name:   "Latin",
			Text:   "Hello, world!",
			Script: unicode.Latin,
		},
		{
			Name:   "Cyrillic",
			Text:   "Привет всем!",
			Script: unicode.Cyrillic,
		},
		{
			Name:   "Georgian",
			Text:   "ქართული ენა მსოფლიო ",
			Script: unicode.Georgian,
		},
		{
			Name:   "Han",
			Text:   "県見夜上温国阪題富販",
			Script: unicode.Han,
		},
		{
			Name:   "Arabic",
			Text:   " ككل حوالي 1.6، ومعظم الناس ",
			Script: unicode.Arabic,
		},
		{
			Name:   "Devanagari",
			Text:   "हिमालयी वन चिड़िया (जूथेरा सालिमअली) चिड़िया की एक प्रजाति है",
			Script: unicode.Devanagari,
		},
		{
			Name:   "Hebrew",
			Text:   "היסטוריה והתפתחות של האלפבית העברי",
			Script: unicode.Hebrew,
		},
		{
			Name:   "Ethiopic",
			Text:   "የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ",
			Script: unicode.Ethiopic,
		},
		{
			Name:   "Cyrillic and English",
			Text:   "Привет! Текст на русском with some English.",
			Script: unicode.Cyrillic,
		},
		{
			Name:   "Russian and English",
			Text:   "Russian word любовь means love.",
			Script: unicode.Latin,
		},
		{
			Name:   "Bengali",
			Text:   "আমি ভালো আছি, ধন্যবাদ!",
			Script: unicode.Bengali,
		},
		{
			Name:   "Japanese",
			Text:   "ハローワールド",
			Script: HiraganaKatakana,
		},
	} {
		t.Run(testcase.Name, func(t *testing.T) {
			got := DetectScript(testcase.Text, DefaultOptions.Scripts)
			if got != testcase.Script {
				t.Error("got", got, "but expected", testcase.Script)
			}
		})
	}
}
