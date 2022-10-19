package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/askeladdk/langdet"
)

var corpora = []struct {
	Name       string
	Tag        string
	CorpusFile string
	OutFile    string
}{
	{
		Name:       "German",
		Tag:        "language.German",
		CorpusFile: "corpora/deu.txt",
		OutFile:    "lang_deu.go",
	},
	{
		Name:       "English",
		Tag:        "language.English",
		CorpusFile: "corpora/eng.txt",
		OutFile:    "lang_eng.go",
	},
	{
		Name:       "Dutch",
		Tag:        "language.Dutch",
		CorpusFile: "corpora/nld.txt",
		OutFile:    "lang_nld.go",
	},
	{
		Name:       "French",
		Tag:        "language.French",
		CorpusFile: "corpora/fra.txt",
		OutFile:    "lang_fra.go",
	},
	{
		Name:       "Spanish",
		Tag:        "language.Spanish",
		CorpusFile: "corpora/spa.txt",
		OutFile:    "lang_spa.go",
	},
	{
		Name:       "Portuguese",
		Tag:        "language.Portuguese",
		CorpusFile: "corpora/por.txt",
		OutFile:    "lang_por.go",
	},
	{
		Name:       "Italian",
		Tag:        "language.Italian",
		CorpusFile: "corpora/ita.txt",
		OutFile:    "lang_ita.go",
	},
	{
		Name:       "Swedish",
		Tag:        "language.Swedish",
		CorpusFile: "corpora/swe.txt",
		OutFile:    "lang_swe.go",
	},
	{
		Name:       "Finnish",
		Tag:        "language.Finnish",
		CorpusFile: "corpora/fin.txt",
		OutFile:    "lang_fin.go",
	},
	{
		Name:       "NorwegianBokmål",
		Tag:        "NorwegianBokmålTag",
		CorpusFile: "corpora/nob.txt",
		OutFile:    "lang_nob.go",
	},
	{
		Name:       "NorwegianNynorsk",
		Tag:        "NorwegianNynorskTag",
		CorpusFile: "corpora/nno.txt",
		OutFile:    "lang_nno.go",
	},
	{
		Name:       "Danish",
		Tag:        "language.Danish",
		CorpusFile: "corpora/dan.txt",
		OutFile:    "lang_dan.go",
	},
	{
		Name:       "Icelandic",
		Tag:        "language.Icelandic",
		CorpusFile: "corpora/isl.txt",
		OutFile:    "lang_isl.go",
	},
	{
		Name:       "Irish",
		Tag:        "IrishTag",
		CorpusFile: "corpora/gle.txt",
		OutFile:    "lang_gle.go",
	},
	{
		Name:       "Latin",
		Tag:        "LatinTag",
		CorpusFile: "corpora/lat.txt",
		OutFile:    "lang_lat.go",
	},
	{
		Name:       "Hungarian",
		Tag:        "language.Hungarian",
		CorpusFile: "corpora/hun.txt",
		OutFile:    "lang_hun.go",
	},
	{
		Name:       "Bulgarian",
		Tag:        "language.Bulgarian",
		CorpusFile: "corpora/bul.txt",
		OutFile:    "lang_bul.go",
	},
	{
		Name:       "Romanian",
		Tag:        "language.Romanian",
		CorpusFile: "corpora/ron.txt",
		OutFile:    "lang_ron.go",
	},
	{
		Name:       "Czech",
		Tag:        "language.Czech",
		CorpusFile: "corpora/ces.txt",
		OutFile:    "lang_ces.go",
	},
	{
		Name:       "Slovak",
		Tag:        "language.Slovak",
		CorpusFile: "corpora/slk.txt",
		OutFile:    "lang_slk.go",
	},
	{
		Name:       "Slovenian",
		Tag:        "language.Slovenian",
		CorpusFile: "corpora/slv.txt",
		OutFile:    "lang_slv.go",
	},
	{
		Name:       "Estonian",
		Tag:        "language.Estonian",
		CorpusFile: "corpora/est.txt",
		OutFile:    "lang_est.go",
	},
	{
		Name:       "Latvian",
		Tag:        "language.Latvian",
		CorpusFile: "corpora/lav.txt",
		OutFile:    "lang_lav.go",
	},
	{
		Name:       "Lithuanian",
		Tag:        "language.Lithuanian",
		CorpusFile: "corpora/lit.txt",
		OutFile:    "lang_lit.go",
	},
	{
		Name:       "Polish",
		Tag:        "language.Polish",
		CorpusFile: "corpora/pol.txt",
		OutFile:    "lang_pol.go",
	},
	{
		Name:       "Russian",
		Tag:        "language.Russian",
		CorpusFile: "corpora/rus.txt",
		OutFile:    "lang_rus.go",
	},
	{
		Name:       "Ukrainian",
		Tag:        "language.Ukrainian",
		CorpusFile: "corpora/ukr.txt",
		OutFile:    "lang_ukr.go",
	},
	{
		Name:       "Maltese",
		Tag:        "MalteseTag",
		CorpusFile: "corpora/mlt.txt",
		OutFile:    "lang_mlt.go",
	},
	{
		Name:       "Croatian",
		Tag:        "language.Croatian",
		CorpusFile: "corpora/hrv.txt",
		OutFile:    "lang_hrv.go",
	},
	{
		Name:       "Albanian",
		Tag:        "language.Albanian",
		CorpusFile: "corpora/sqi.txt",
		OutFile:    "lang_sqi.go",
	},
	{
		Name:       "Bosnian",
		Tag:        "BosnianTag",
		CorpusFile: "corpora/bos.txt",
		OutFile:    "lang_bos.go",
	},
	{
		Name:       "Macedonian",
		Tag:        "language.Macedonian",
		CorpusFile: "corpora/mkd.txt",
		OutFile:    "lang_mkd.go",
	},
	{
		Name:       "Serbian",
		Tag:        "language.Serbian",
		CorpusFile: "corpora/srp.txt",
		OutFile:    "lang_srp.go",
	},
	{
		Name:       "Belarusian",
		Tag:        "BelarusianTag",
		CorpusFile: "corpora/bel.txt",
		OutFile:    "lang_bel.go",
	},
	{
		Name:       "Turkish",
		Tag:        "language.Turkish",
		CorpusFile: "corpora/tur.txt",
		OutFile:    "lang_tur.go",
	},
	{
		Name:       "Luxembourgish",
		Tag:        "LuxembourgishTag",
		CorpusFile: "corpora/ltz.txt",
		OutFile:    "lang_ltz.go",
	},
}

func main() {
	for _, corpus := range corpora {
		f, err := os.Open(corpus.CorpusFile)
		if err != nil {
			log.Println(corpus.Name, err)
			continue
		}
		defer f.Close()

		b, _ := io.ReadAll(f)
		tgs := langdet.Train(b, len(b)/2)

		w, err := os.Create(corpus.OutFile)
		if err != nil {
			log.Println(corpus.Name, err)
			continue
		}
		defer w.Close()

		fmt.Fprintln(w, "// Code generated with cmd/train.go. DO NOT EDIT.")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "package langdet")
		fmt.Fprintln(w)
		if strings.HasPrefix(corpus.Tag, "language.") {
			fmt.Fprintln(w, "import (")
			fmt.Fprintln(w, "\t\"golang.org/x/text/language\"")
			fmt.Fprintln(w, ")")
			fmt.Fprintln(w)
		}

		fmt.Fprintf(w, "var _%sTrigrams = []Trigram{\n", corpus.Name)
		for _, tg := range tgs[:400] {
			fmt.Fprintf(w, "\t{'%s', '%s', '%s'},\n", string(tg[0]), string(tg[1]), string(tg[2]))
		}
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w)

		fmt.Fprintf(w, "// %s is a language profile.\n", corpus.Name)
		fmt.Fprintf(w, "var %s = Language {\n", corpus.Name)
		fmt.Fprintf(w, "\tTag: %s,\n", corpus.Tag)
		fmt.Fprintf(w, "\tTrigrams: _%sTrigrams,\n", corpus.Name)
		fmt.Fprintln(w, "}")
	}
}
