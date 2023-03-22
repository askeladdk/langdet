# langdet - Language Detection for Go

[![GoDoc](https://godoc.org/github.com/askeladdk/langdet?status.png)](https://godoc.org/github.com/askeladdk/langdet)
[![Go Report Card](https://goreportcard.com/badge/github.com/askeladdk/langdet)](https://goreportcard.com/report/github.com/askeladdk/langdet)
[![Coverage Status](https://coveralls.io/repos/github/askeladdk/langdet/badge.svg?branch=master)](https://coveralls.io/github/askeladdk/langdet?branch=master)

## Overview

Package langdet detects natural languages in text using a straightforward implementation of [trigram based text categorization](https://web.archive.org/web/20180509095530/http://odur.let.rug.nl/~vannoord/TextCat/textcat.pdf). The most commonly used languages worldwide are supported out of the box, but the code is flexible enough to accept any set of languages.

Langdet first detects the writing script in order to narrow down the number of languages to test against. Some writing scripts are used by only a single language (Korean, Greek, etc). In that case the language is returned directly without needing to do trigram analysis. Otherwise, it matches each language profile under the detected writing script against the input text and returns a result set listing the languages ordered by confidence.

## Install

```
go get -u github.com/askeladdk/langdet
```

## Quickstart

Use `DetectLanguage` to detect the language of a string. It returns the BCP 47 language tag of the language with the highest probability. If no language was detected, the function returns `language.Und`.

```go
detectedLanguage := langdet.DetectLanguage(s)
```

Use `DetectLanguageWithOptions` if you need more control. `DetectLanguage` is a shorthand for this function using `DefaultOptions`. Unlike `DetectLanguage`, `DetectLanguageWithOptions` returns a slice of `Result`s listing the probabilities of all languages using the detected writing script ordered by probability.

```go
results := langdet.DetectLanguageWithOptions(s, DefaultOptions)
```

Use `Options` to configure the detector. Any number of writing scripts and languages can be detected by setting the `Scripts` and `Languages` fields. Use the `Train` function to build language profiles. Use `MinConfidence` and `MinRelConfidence` to filter languages by confidence.

```go
myLang := langdet.Language {
    Tag: language.Make("zz"),
    Trigrams: langdet.Train(trainingSet),
}

options := langdet.Options {
    Scripts: []*unicode.RangeTable{
        unicode.Latin,
    },
    Languages: map[*unicode.RangeTable]langdet.Languages {
        unicode.Latin: {
            Languages: []langdet.Languge {
                langdet.Dutch,
                langdet.French,
                myLang,
            },
        },
    },
}

results := langdet.DetectLanguageWithOptions(s, options)
```

Read the rest of the [documentation on pkg.go.dev](https://pkg.go.dev/github.com/askeladdk/langdet). It's easy-peasy!

## License

Package langdet is released under the terms of the ISC license.
