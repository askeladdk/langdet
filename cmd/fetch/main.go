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
)

var filenames = []string{
	"eng_news_2020_10K.tar.gz", "eng.txt",
	"deu_news_2020_10K.tar.gz", "deu.txt",
	"nld_news_2020_10K.tar.gz", "nld.txt",
	"fra_news_2020_10K.tar.gz", "fra.txt",
	"spa_news_2020_10K.tar.gz", "spa.txt",
	"por_news_2020_10K.tar.gz", "por.txt",
	"ita_news_2020_10K.tar.gz", "ita.txt",
	"swe_news_2020_10K.tar.gz", "swe.txt",
	"fin_news_2020_10K.tar.gz", "fin.txt",
	"nob_newscrawl_2019_10K.tar.gz", "nob.txt",
	"nno-no_web_2020_10K.tar.gz", "nno.txt",
	"dan_news_2020_10K.tar.gz", "dan.txt",
	"isl_news_2020_10K.tar.gz", "isl.txt",
	"gle_wikipedia_2021_10K.tar.gz", "gle.txt",
	"lat_wikipedia_2021_10K.tar.gz", "lat.txt",
	"hun_news_2020_10K.tar.gz", "hun.txt",
	"bul_news_2020_10K.tar.gz", "bul.txt",
	"ron_news_2020_10K.tar.gz", "ron.txt",
	"ces_news_2020_10K.tar.gz", "ces.txt",
	"slk_news_2020_10K.tar.gz", "slk.txt",
	"slv_news_2020_10K.tar.gz", "slv.txt",
	"est_news_2020_10K.tar.gz", "est.txt",
	"lav_news_2020_10K.tar.gz", "lav.txt",
	"lit_news_2020_10K.tar.gz", "lit.txt",
	"pol_news_2020_10K.tar.gz", "pol.txt",
	"rus_news_2020_10K.tar.gz", "rus.txt",
	"ukr_news_2020_10K.tar.gz", "ukr.txt",
	"mlt_news_2020_10K.tar.gz", "mlt.txt",
	"hrv_news_2020_10K.tar.gz", "hrv.txt",
	"sqi_news_2020_10K.tar.gz", "sqi.txt",
	"bos_news_2020_10K.tar.gz", "bos.txt",
	"mkd_news_2020_10K.tar.gz", "mkd.txt",
	"srp_news_2020_10K.tar.gz", "srp.txt",
	"bel_news_2020_10K.tar.gz", "bel.txt",
	"tur_news_2020_10K.tar.gz", "tur.txt",
	"ltz-lu_web_2020_10K.tar.gz", "ltz.txt",
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
	if _, err := os.Stat("corpora/" + outfile); err == nil {
		return false, errors.New("EXISTS")
	}

	u := fmt.Sprintf("https://pcai056.informatik.uni-leipzig.de/downloads/corpora/%s", corpus)

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

	f, err := os.Create("corpora/" + outfile)
	if err != nil {
		return false, err
	}

	defer f.Close()

	return extractSentences(f, res.Body)
}

func main() {
	_ = os.Mkdir("corpora", 0755)

	c := http.Client{}

	for i := 0; i < len(filenames); i += 2 {
		targzfile := filenames[i]

		if found, err := fetch(&c, filenames[i+1], targzfile); err != nil {
			log.Println(targzfile, err)
		} else if !found {
			log.Println(targzfile, "NOT FOUND")
		} else {
			log.Println(targzfile, "SAVED")
		}
	}
}
