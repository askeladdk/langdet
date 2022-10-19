package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/askeladdk/langdet"
)

var corpora = []struct {
	Name       string
	Tag        string
	SourceFile string
	CorpusFile string
	OutFile    string
}{
	{
		Name:       "German",
		Tag:        "language.German",
		SourceFile: "deu_news_2020_10K.tar.gz",
		CorpusFile: "corpora/deu.txt",
		OutFile:    "lang_deu.go",
	},
	{
		Name:       "English",
		Tag:        "language.English",
		SourceFile: "eng_news_2020_10K.tar.gz",
		CorpusFile: "corpora/eng.txt",
		OutFile:    "lang_eng.go",
	},
	{
		Name:       "Dutch",
		Tag:        "language.Dutch",
		SourceFile: "nld_news_2020_10K.tar.gz",
		CorpusFile: "corpora/nld.txt",
		OutFile:    "lang_nld.go",
	},
	{
		Name:       "French",
		Tag:        "language.French",
		SourceFile: "fra_news_2020_10K.tar.gz",
		CorpusFile: "corpora/fra.txt",
		OutFile:    "lang_fra.go",
	},
	{
		Name:       "Spanish",
		Tag:        "language.Spanish",
		SourceFile: "spa_news_2020_10K.tar.gz",
		CorpusFile: "corpora/spa.txt",
		OutFile:    "lang_spa.go",
	},
	{
		Name:       "Portuguese",
		Tag:        "language.Portuguese",
		SourceFile: "por_news_2020_10K.tar.gz",
		CorpusFile: "corpora/por.txt",
		OutFile:    "lang_por.go",
	},
	{
		Name:       "Italian",
		Tag:        "language.Italian",
		SourceFile: "ita_news_2020_10K.tar.gz",
		CorpusFile: "corpora/ita.txt",
		OutFile:    "lang_ita.go",
	},
	{
		Name:       "Swedish",
		Tag:        "language.Swedish",
		SourceFile: "swe_news_2020_10K.tar.gz",
		CorpusFile: "corpora/swe.txt",
		OutFile:    "lang_swe.go",
	},
	{
		Name:       "Finnish",
		Tag:        "language.Finnish",
		SourceFile: "fin_news_2020_10K.tar.gz",
		CorpusFile: "corpora/fin.txt",
		OutFile:    "lang_fin.go",
	},
	{
		Name:       "NorwegianBokmål",
		Tag:        "NorwegianBokmålTag",
		SourceFile: "nob_newscrawl_2019_10K.tar.gz",
		CorpusFile: "corpora/nob.txt",
		OutFile:    "lang_nob.go",
	},
	{
		Name:       "NorwegianNynorsk",
		Tag:        "NorwegianNynorskTag",
		SourceFile: "nno-no_web_2020_10K.tar.gz",
		CorpusFile: "corpora/nno.txt",
		OutFile:    "lang_nno.go",
	},
	{
		Name:       "Danish",
		Tag:        "language.Danish",
		SourceFile: "dan_news_2020_10K.tar.gz",
		CorpusFile: "corpora/dan.txt",
		OutFile:    "lang_dan.go",
	},
	{
		Name:       "Icelandic",
		Tag:        "language.Icelandic",
		SourceFile: "isl_news_2020_10K.tar.gz",
		CorpusFile: "corpora/isl.txt",
		OutFile:    "lang_isl.go",
	},
	{
		Name:       "Irish",
		Tag:        "IrishTag",
		SourceFile: "gle_wikipedia_2021_10K.tar.gz",
		CorpusFile: "corpora/gle.txt",
		OutFile:    "lang_gle.go",
	},
	{
		Name:       "Latin",
		Tag:        "LatinTag",
		SourceFile: "lat_wikipedia_2021_10K.tar.gz",
		CorpusFile: "corpora/lat.txt",
		OutFile:    "lang_lat.go",
	},
	{
		Name:       "Hungarian",
		Tag:        "language.Hungarian",
		SourceFile: "hun_wikipedia_2021_10K.tar.gz",
		CorpusFile: "corpora/hun.txt",
		OutFile:    "lang_hun.go",
	},
	{
		Name:       "Bulgarian",
		Tag:        "language.Bulgarian",
		SourceFile: "bul_news_2020_10K.tar.gz",
		CorpusFile: "corpora/bul.txt",
		OutFile:    "lang_bul.go",
	},
	{
		Name:       "Romanian",
		Tag:        "language.Romanian",
		SourceFile: "ron_news_2020_10K.tar.gz",
		CorpusFile: "corpora/ron.txt",
		OutFile:    "lang_ron.go",
	},
	{
		Name:       "Czech",
		Tag:        "language.Czech",
		SourceFile: "ces_news_2020_10K.tar.gz",
		CorpusFile: "corpora/ces.txt",
		OutFile:    "lang_ces.go",
	},
	{
		Name:       "Slovak",
		Tag:        "language.Slovak",
		SourceFile: "slk_news_2020_10K.tar.gz",
		CorpusFile: "corpora/slk.txt",
		OutFile:    "lang_slk.go",
	},
	{
		Name:       "Slovenian",
		Tag:        "language.Slovenian",
		SourceFile: "slv_news_2020_10K.tar.gz",
		CorpusFile: "corpora/slv.txt",
		OutFile:    "lang_slv.go",
	},
	{
		Name:       "Estonian",
		Tag:        "language.Estonian",
		SourceFile: "est_news_2020_10K.tar.gz",
		CorpusFile: "corpora/est.txt",
		OutFile:    "lang_est.go",
	},
	{
		Name:       "Latvian",
		Tag:        "language.Latvian",
		SourceFile: "lav_news_2020_10K.tar.gz",
		CorpusFile: "corpora/lav.txt",
		OutFile:    "lang_lav.go",
	},
	{
		Name:       "Lithuanian",
		Tag:        "language.Lithuanian",
		SourceFile: "lit_news_2020_10K.tar.gz",
		CorpusFile: "corpora/lit.txt",
		OutFile:    "lang_lit.go",
	},
	{
		Name:       "Polish",
		Tag:        "language.Polish",
		SourceFile: "pol_news_2020_10K.tar.gz",
		CorpusFile: "corpora/pol.txt",
		OutFile:    "lang_pol.go",
	},
	{
		Name:       "Russian",
		Tag:        "language.Russian",
		SourceFile: "rus_news_2020_10K.tar.gz",
		CorpusFile: "corpora/rus.txt",
		OutFile:    "lang_rus.go",
	},
	{
		Name:       "Ukrainian",
		Tag:        "language.Ukrainian",
		SourceFile: "ukr_news_2020_10K.tar.gz",
		CorpusFile: "corpora/ukr.txt",
		OutFile:    "lang_ukr.go",
	},
	{
		Name:       "Maltese",
		Tag:        "MalteseTag",
		SourceFile: "mlt_news_2020_10K.tar.gz",
		CorpusFile: "corpora/mlt.txt",
		OutFile:    "lang_mlt.go",
	},
	{
		Name:       "Croatian",
		Tag:        "language.Croatian",
		SourceFile: "hrv_news_2020_10K.tar.gz",
		CorpusFile: "corpora/hrv.txt",
		OutFile:    "lang_hrv.go",
	},
	{
		Name:       "Albanian",
		Tag:        "language.Albanian",
		SourceFile: "sqi_news_2020_10K.tar.gz",
		CorpusFile: "corpora/sqi.txt",
		OutFile:    "lang_sqi.go",
	},
	{
		Name:       "Bosnian",
		Tag:        "BosnianTag",
		SourceFile: "bos_news_2020_10K.tar.gz",
		CorpusFile: "corpora/bos.txt",
		OutFile:    "lang_bos.go",
	},
	{
		Name:       "Macedonian",
		Tag:        "language.Macedonian",
		SourceFile: "mkd_news_2020_10K.tar.gz",
		CorpusFile: "corpora/mkd.txt",
		OutFile:    "lang_mkd.go",
	},
	{
		Name:       "Serbian",
		Tag:        "language.Serbian",
		SourceFile: "srp_news_2020_10K.tar.gz",
		CorpusFile: "corpora/srp.txt",
		OutFile:    "lang_srp.go",
	},
	{
		Name:       "Belarusian",
		Tag:        "BelarusianTag",
		SourceFile: "bel_news_2020_10K.tar.gz",
		CorpusFile: "corpora/bel.txt",
		OutFile:    "lang_bel.go",
	},
	{
		Name:       "Turkish",
		Tag:        "language.Turkish",
		SourceFile: "tur_news_2020_10K.tar.gz",
		CorpusFile: "corpora/tur.txt",
		OutFile:    "lang_tur.go",
	},
	{
		Name:       "Luxembourgish",
		Tag:        "LuxembourgishTag",
		SourceFile: "ltz-lu_web_2020_10K.tar.gz",
		CorpusFile: "corpora/ltz.txt",
		OutFile:    "lang_ltz.go",
	},
}

func extractSentences(w io.Writer, r io.Reader) (bool, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return false, err
	}

	tr := tar.NewReader(gz)

	for {
		h, err := tr.Next()
		if errors.Is(err, io.EOF) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		if strings.HasSuffix(h.Name, "sentences.txt") {
			scan := bufio.NewScanner(tr)
			for scan.Scan() {
				s := scan.Text()
				if i := strings.IndexByte(s, '\t'); i != -1 {
					s = s[i+1:]
				}
				_, err := fmt.Fprintln(w, s)
				if err != nil {
					return true, err
				}
			}

			return true, nil
		}
	}
}

func fetch(c *http.Client, outfile, corpus string) (bool, error) {
	u := "https://pcai056.informatik.uni-leipzig.de/downloads/corpora/" + corpus

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", u, nil)
	res, err := c.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, errors.New(http.StatusText(res.StatusCode))
	}

	f, err := os.Create(outfile)
	if err != nil {
		return false, err
	}

	defer f.Close()

	return extractSentences(f, res.Body)
}

func main() {
	_ = os.Mkdir("corpora", 0755)
	c := &http.Client{}

	for _, corpus := range corpora {
		if _, err := os.Stat(corpus.CorpusFile); err != nil {
			if found, err := fetch(c, corpus.CorpusFile, corpus.SourceFile); err != nil {
				log.Println(corpus.SourceFile, err)
				continue
			} else if !found {
				log.Println(corpus.SourceFile, "NOT FOUND")
				continue
			}

			log.Println(corpus.SourceFile, "SAVED")
		}

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
