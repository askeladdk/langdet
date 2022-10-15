package langdet

import (
	"testing"
)

func newTrigram(s string) Trigram {
	r := []rune(s)
	return Trigram{r[0], r[1], r[2]}
}

func TestCountTrigrams(t *testing.T) {
	test := "De kat krabt de krullen van de trap."

	expected := map[Trigram]int{
		newTrigram("__d"): 1,
		newTrigram("_de"): 3,
		newTrigram("_ka"): 1,
		newTrigram("_kr"): 2,
		newTrigram("_tr"): 1,
		newTrigram("_va"): 1,
		newTrigram("abt"): 1,
		newTrigram("an_"): 1,
		newTrigram("at_"): 1,
		newTrigram("bt_"): 1,
		newTrigram("de_"): 3,
		newTrigram("e_k"): 2,
		newTrigram("e_t"): 1,
		newTrigram("en_"): 1,
		newTrigram("kat"): 1,
		newTrigram("kra"): 1,
		newTrigram("kru"): 1,
		newTrigram("len"): 1,
		newTrigram("lle"): 1,
		newTrigram("n_d"): 1,
		newTrigram("n_v"): 1,
		newTrigram("rab"): 1,
		newTrigram("rap"): 1,
		newTrigram("rul"): 1,
		newTrigram("t_d"): 1,
		newTrigram("t_k"): 1,
		newTrigram("tra"): 1,
		newTrigram("ull"): 1,
		newTrigram("van"): 1,
		newTrigram("ap_"): 1,
	}

	counts := make(map[Trigram]int, 300)

	countTrigrams([]byte(test), counts)

	assertFalse(t, len(counts) != len(expected), "wrong number of trigrams")

	for k, v := range counts {
		assertFalse(t, expected[k] != v, "wrong count")
	}
}

func TestTrigramDistance(t *testing.T) {
	trigrams := []Trigram{
		{'_', 'h', 'e'},
		{'h', 'e', 'l'},
		{'e', 'l', 'l'},
		{'l', 'l', 'o'},
		{'l', 'o', '_'},
	}

	for _, testcase := range []struct {
		Name     string
		Distance int
		Document map[Trigram]int
	}{
		{
			Name:     "Zero",
			Distance: 0,
			Document: map[Trigram]int{
				{'_', 'h', 'e'}: 0,
				{'h', 'e', 'l'}: 1,
				{'e', 'l', 'l'}: 2,
				{'l', 'l', 'o'}: 3,
				{'l', 'o', '_'}: 4,
			},
		},
		{
			Name:     "Reverse",
			Distance: 12,
			Document: map[Trigram]int{
				{'_', 'h', 'e'}: 4,
				{'h', 'e', 'l'}: 3,
				{'e', 'l', 'l'}: 2,
				{'l', 'l', 'o'}: 1,
				{'l', 'o', '_'}: 0,
			},
		},
		{
			Name:     "French",
			Distance: len(trigrams) * len(trigrams),
			Document: map[Trigram]int{
				{'_', 'b', 'o'}: 0,
				{'b', 'o', 'n'}: 1,
				{'o', 'n', 'j'}: 2,
				{'n', 'j', 'o'}: 3,
				{'j', 'o', 'u'}: 4,
				{'o', 'u', 'r'}: 5,
				{'u', 'r', '_'}: 6,
			},
		},
	} {
		t.Run(testcase.Name, func(t *testing.T) {
			dist := distance(trigrams, testcase.Document)
			assertFalse(t, dist != testcase.Distance, "incorrect distance")
		})
	}
}
