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
