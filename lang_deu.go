// Code generated with cmd/train.go. DO NOT EDIT.

package langdet

import (
	"golang.org/x/text/language"
)

var _GermanTrigrams = []Trigram{
	{'e', 'n', '_'},
	{'e', 'r', '_'},
	{'_', 'd', 'e'},
	{'d', 'e', 'r'},
	{'i', 'e', '_'},
	{'e', 'i', 'n'},
	{'i', 'c', 'h'},
	{'c', 'h', '_'},
	{'s', 'c', 'h'},
	{'d', 'i', 'e'},
	{'_', 'd', 'i'},
	{'n', '_', 'd'},
	{'i', 'n', '_'},
	{'n', 'd', '_'},
	{'_', 'e', 'i'},
	{'_', 'u', 'n'},
	{'d', 'e', 'n'},
	{'u', 'n', 'd'},
	{'c', 'h', 'e'},
	{'_', 'a', 'u'},
	{'_', 'd', 'a'},
	{'t', 'e', 'n'},
	{'_', 'b', 'e'},
	{'_', 'i', 'n'},
	{'_', 'g', 'e'},
	{'g', 'e', 'n'},
	{'i', 'n', 'e'},
	{'n', 'd', 'e'},
	{'t', 'e', '_'},
	{'c', 'h', 't'},
	{'u', 'n', 'g'},
	{'t', 'e', 'r'},
	{'n', '_', 's'},
	{'e', 's', '_'},
	{'_', 's', 'i'},
	{'t', '_', 'd'},
	{'_', 'z', 'u'},
	{'e', 'i', 't'},
	{'s', 't', 'e'},
	{'_', 'v', 'o'},
	{'v', 'e', 'r'},
	{'b', 'e', 'r'},
	{'i', 't', '_'},
	{'h', 'e', 'n'},
	{'_', 'w', 'e'},
	{'d', 'a', 's'},
	{'_', 'w', 'i'},
	{'r', '_', 'd'},
	{'n', '_', 'a'},
	{'n', 'g', '_'},
	{'_', 'm', 'i'},
	{'o', 'n', '_'},
	{'e', '_', 'd'},
	{'h', 't', '_'},
	{'a', 's', '_'},
	{'e', 'm', '_'},
	{'s', 't', '_'},
	{'n', 'e', 'n'},
	{'n', 'g', 'e'},
	{'m', 'i', 't'},
	{'n', '_', 'e'},
	{'_', 'h', 'a'},
	{'e', '_', 's'},
	{'a', 'c', 'h'},
	{'_', 'a', 'n'},
	{'_', 'e', 'r'},
	{'i', 's', 't'},
	{'a', 'u', 'f'},
	{'e', 'r', 'e'},
	{'n', '_', 'w'},
	{'_', 's', 'e'},
	{'r', 'e', 'n'},
	{'_', 'v', 'e'},
	{'a', 'n', 'd'},
	{'e', 'r', 's'},
	{'_', 's', 't'},
	{'n', 'e', '_'},
	{'a', 'u', 's'},
	{'l', 'i', 'c'},
	{'n', 't', 'e'},
	{'l', 'l', 'e'},
	{'_', 's', 'c'},
	{'r', 'd', 'e'},
	{'i', 'e', 'r'},
	{'r', 'e', 'i'},
	{'e', 'n', 't'},
	{'r', '_', 's'},
	{'e', 'r', 't'},
	{'n', '_', 'i'},
	{'r', '_', 'a'},
	{'e', '_', 'a'},
	{'n', '_', 'u'},
	{'_', 's', 'o'},
	{'m', 'e', 'n'},
	{'s', 't', 'a'},
	{'b', 'e', 'n'},
	{'e', 'n', 'd'},
	{'_', 'a', 'l'},
	{'e', 'i', '_'},
	{'s', 'i', 'c'},
	{'n', 'e', 'r'},
	{'b', 'e', 'i'},
	{'s', 'e', 'n'},
	{'i', 'g', 'e'},
	{'n', '_', 'b'},
	{'u', 'c', 'h'},
	{'g', 'e', 's'},
	{'_', 'i', 'm'},
	{'z', 'u', '_'},
	{'d', 'e', 's'},
	{'e', 'r', 'n'},
	{'t', '_', 'e'},
	{'i', 'm', '_'},
	{'n', '_', 'm'},
	{'e', '_', 'e'},
	{'a', 'b', 'e'},
	{'w', 'e', 'i'},
	{'_', 'f', 'ü'},
	{'r', 't', '_'},
	{'s', 'e', 'i'},
	{'_', 'm', 'a'},
	{'d', 'e', 'm'},
	{'v', 'o', 'n'},
	{'s', '_', 'd'},
	{'u', 'f', '_'},
	{'t', '_', 's'},
	{'a', 'n', '_'},
	{'f', 'ü', 'r'},
	{'s', 's', 'e'},
	{'h', 'r', 'e'},
	{'d', 'e', '_'},
	{'a', 's', 's'},
	{'i', 'n', 'd'},
	{'ü', 'r', '_'},
	{'l', 'e', 'n'},
	{'_', 'i', 's'},
	{'t', '_', 'a'},
	{'v', 'o', 'r'},
	{'t', '_', 'w'},
	{'h', 'e', 'r'},
	{'g', 'e', 'r'},
	{'_', 'n', 'i'},
	{'g', 'e', '_'},
	{'n', '_', 'g'},
	{'h', 'e', '_'},
	{'w', 'e', 'r'},
	{'r', '_', 'e'},
	{'i', 'o', 'n'},
	{'_', 'm', 'e'},
	{'o', 'c', 'h'},
	{'_', 'w', 'a'},
	{'e', 'g', 'e'},
	{'r', 't', 'e'},
	{'n', '_', 'k'},
	{'_', 'n', 'a'},
	{'w', 'i', 'e'},
	{'h', '_', 'd'},
	{'i', 't', 'e'},
	{'a', 'l', 'l'},
	{'n', '_', 'v'},
	{'w', 'i', 'r'},
	{'_', 'a', 'b'},
	{'e', '_', 'i'},
	{'u', 'm', '_'},
	{'t', 't', 'e'},
	{'e', '_', 'b'},
	{'l', 'e', '_'},
	{'a', 'n', 'g'},
	{'_', 'e', 's'},
	{'s', 'i', 'e'},
	{'e', 'l', 'l'},
	{'e', 's', 't'},
	{'i', 's', 'c'},
	{'n', 'i', 'c'},
	{'i', 'e', 'l'},
	{'e', 's', 'e'},
	{'d', '_', 'd'},
	{'e', 'n', 's'},
	{'n', 'a', 'c'},
	{'r', 'u', 'n'},
	{'e', '_', 'm'},
	{'u', 's', '_'},
	{'a', 'u', 'c'},
	{'s', 's', '_'},
	{'i', 'n', 'g'},
	{'h', 'r', '_'},
	{'l', 'a', 'n'},
	{'e', 'r', 'd'},
	{'_', 'k', 'o'},
	{'n', 'n', '_'},
	{'a', 'n', 'n'},
	{'_', 'r', 'e'},
	{'r', '_', 'w'},
	{'m', 'm', 'e'},
	{'a', 'h', 'r'},
	{'n', '_', 'f'},
	{'e', '_', 'g'},
	{'s', '_', 's'},
	{'_', 'p', 'r'},
	{'l', 't', 'e'},
	{'n', '_', 'z'},
	{'r', '_', 'b'},
	{'t', '_', 'i'},
	{'r', 'e', '_'},
	{'n', '_', 'n'},
	{'s', 'e', '_'},
	{'l', 's', '_'},
	{'a', 'm', '_'},
	{'z', 'e', 'i'},
	{'r', '_', 'i'},
	{'e', '_', 'v'},
	{'e', 'n', 'e'},
	{'o', 'l', 'l'},
	{'e', 'i', 's'},
	{'n', '_', 'h'},
	{'r', 's', 't'},
	{'e', '_', 'w'},
	{'h', 'e', 'i'},
	{'a', 'l', 's'},
	{'a', 'r', '_'},
	{'_', 'l', 'a'},
	{'h', 'a', 't'},
	{'l', 'e', 'i'},
	{'t', 'r', 'a'},
	{'r', '_', 'm'},
	{'s', '_', 'a'},
	{'r', 'n', '_'},
	{'e', 't', '_'},
	{'i', 'e', 's'},
	{'e', 'l', '_'},
	{'_', 's', 'p'},
	{'a', 'g', 'e'},
	{'t', 'i', 'o'},
	{'e', 'r', 'i'},
	{'n', 'n', 'e'},
	{'w', 'a', 'r'},
	{'e', 'd', 'e'},
	{'e', 'b', 'e'},
	{'c', 'h', 'a'},
	{'i', 'e', 'n'},
	{'_', 'l', 'e'},
	{'l', 'i', 'e'},
	{'ü', 'b', 'e'},
	{'f', 'e', 'n'},
	{'e', '_', 'f'},
	{'e', 'h', 'r'},
	{'l', 'e', 'r'},
	{'_', 'a', 'm'},
	{'n', 't', '_'},
	{'n', 'g', 's'},
	{'a', 't', '_'},
	{'e', 't', 'z'},
	{'s', '_', 'e'},
	{'t', 'e', 'l'},
	{'_', 'k', 'a'},
	{'t', 'i', 'g'},
	{'c', 'h', 'l'},
	{'t', 'a', 'g'},
	{'g', 't', '_'},
	{'h', 't', 'e'},
	{'l', 't', '_'},
	{'t', 'e', 'i'},
	{'r', '_', 'g'},
	{'p', 'r', 'o'},
	{'g', '_', 'd'},
	{'_', 'f', 'r'},
	{'u', 'r', '_'},
	{'l', 'l', '_'},
	{'_', 's', 'a'},
	{'_', 'g', 'r'},
	{'o', 'n', 'a'},
	{'a', 'l', 't'},
	{'_', 'b', 'i'},
	{'z', 'e', 'n'},
	{'n', 's', 't'},
	{'r', '_', 'k'},
	{'e', 'i', 'c'},
	{'s', '_', 'i'},
	{'u', 'n', 't'},
	{'r', '_', 'v'},
	{'e', '_', 'k'},
	{'m', 'a', 'n'},
	{'o', 'r', '_'},
	{'_', 'n', 'e'},
	{'e', 't', 'e'},
	{'_', 'ü', 'b'},
	{'t', '_', 'm'},
	{'_', 'n', 'o'},
	{'e', '_', 'z'},
	{'r', 'a', 'n'},
	{'b', 'e', 's'},
	{'e', 'i', 'l'},
	{'_', 'w', 'o'},
	{'t', 's', 'c'},
	{'_', 'h', 'e'},
	{'e', 'c', 'h'},
	{'e', 'l', 'e'},
	{'r', '_', 'n'},
	{'e', '_', 'u'},
	{'_', 'u', 'm'},
	{'a', 't', 'i'},
	{'e', '_', 'n'},
	{'t', '_', 'u'},
	{'i', 't', 't'},
	{'e', 'u', 't'},
	{'f', '_', 'd'},
	{'_', 'e', 'n'},
	{'a', 'r', 't'},
	{'e', 'r', 'l'},
	{'t', 'u', 'n'},
	{'e', 'r', 'g'},
	{'_', 'i', 'h'},
	{'s', 't', 'r'},
	{'i', 'e', 'd'},
	{'i', 'n', 's'},
	{'n', '_', 'l'},
	{'m', 'e', 'r'},
	{'r', '_', 'f'},
	{'s', '_', 'g'},
	{'s', 't', 'i'},
	{'g', 'e', 'b'},
	{'s', 'i', 'n'},
	{'e', 'r', 'k'},
	{'s', 'e', 'r'},
	{'r', '_', 'u'},
	{'e', 'h', 'e'},
	{'i', 's', '_'},
	{'u', 's', 's'},
	{'e', 'n', 'n'},
	{'g', 'e', 'l'},
	{'t', 'l', 'i'},
	{'m', '_', 's'},
	{'_', 'j', 'a'},
	{'a', 'g', '_'},
	{'n', '_', 'p'},
	{'_', 'v', 'i'},
	{'o', 'm', 'm'},
	{'r', 'g', 'e'},
	{'e', '_', 'h'},
	{'_', 'n', 'u'},
	{'h', 'a', 'b'},
	{'o', 'r', 't'},
	{'_', 't', 'r'},
	{'e', 's', 'c'},
	{'r', 'i', 'n'},
	{'n', '_', 't'},
	{'t', 'z', 't'},
	{'t', 'e', 't'},
	{'r', '_', 'h'},
	{'k', 'e', 'n'},
	{'r', 'i', 'e'},
	{'n', 'i', 's'},
	{'t', '_', 'n'},
	{'e', 'l', 't'},
	{'r', 'i', 'c'},
	{'h', 'a', 'l'},
	{'f', 't', '_'},
	{'r', 'd', '_'},
	{'s', '_', 'w'},
	{'i', 'g', '_'},
	{'_', 'c', 'o'},
	{'n', 'e', 'm'},
	{'e', 'r', 'u'},
	{'n', 's', '_'},
	{'r', 'b', 'e'},
	{'s', 'o', 'l'},
	{'n', 'n', 't'},
	{'r', 'a', 'u'},
	{'e', 'r', 'a'},
	{'d', '_', 's'},
	{'c', 'h', 'i'},
	{'v', 'i', 'e'},
	{'t', '_', 'g'},
	{'m', '_', 'd'},
	{'e', '_', 'p'},
	{'k', 'e', 'i'},
	{'k', 'o', 'm'},
	{'a', 't', 't'},
	{'g', 'e', 'g'},
	{'h', 'a', 'f'},
	{'c', 'h', 's'},
	{'_', 'h', 'i'},
	{'h', '_', 'a'},
	{'_', 'f', 'a'},
	{'s', 'o', 'n'},
	{'s', 'p', 'i'},
	{'t', '_', 'b'},
	{'s', 'o', '_'},
	{'k', 'o', 'n'},
	{'e', 's', 's'},
	{'t', '_', 'v'},
	{'e', 'r', 'h'},
	{'_', 'b', 'a'},
	{'j', 'a', 'h'},
	{'_', 'z', 'e'},
	{'_', 't', 'e'},
	{'i', 'd', 'e'},
	{'n', 's', 'c'},
	{'r', '_', 'p'},
	{'r', '_', 'z'},
}

// German profiles the German language.
var German = Language {
	Tag: language.German,
	Trigrams: _GermanTrigrams,
}
