package main

import (
	"fmt"
	"io"
	"log"
	"os"

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
		tgs := langdet.Train(b)

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
		fmt.Fprintln(w, "import (")
		fmt.Fprintln(w, "\t\"golang.org/x/text/language\"")
		fmt.Fprintln(w, ")")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "var %s = Language {\n", corpus.Name)
		fmt.Fprintf(w, "\tTag: %s,\n", corpus.Tag)
		fmt.Fprintln(w, "\tTrigrams: []Trigram{")
		for _, tg := range tgs[:400] {
			fmt.Fprintf(w, "\t\t{'%s', '%s', '%s'},\n", string(tg[0]), string(tg[1]), string(tg[2]))
		}
		fmt.Fprintln(w, "\t},")
		fmt.Fprintln(w, "}")
	}
}
