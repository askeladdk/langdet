package langdet

import (
	"reflect"
	"testing"
)

func trigram(s string) Trigram {
	r := []rune(s)
	return Trigram{r[0], r[1], r[2]}
}

func TestCountTrigrams(t *testing.T) {
	test := "De kat krabt de krullen van de trap."

	expected := map[Trigram]int{
		trigram("__d"): 1,
		trigram("_de"): 3,
		trigram("_ka"): 1,
		trigram("_kr"): 2,
		trigram("_tr"): 1,
		trigram("_va"): 1,
		trigram("abt"): 1,
		trigram("an_"): 1,
		trigram("at_"): 1,
		trigram("bt_"): 1,
		trigram("de_"): 3,
		trigram("e_k"): 2,
		trigram("e_t"): 1,
		trigram("en_"): 1,
		trigram("kat"): 1,
		trigram("kra"): 1,
		trigram("kru"): 1,
		trigram("len"): 1,
		trigram("lle"): 1,
		trigram("n_d"): 1,
		trigram("n_v"): 1,
		trigram("rab"): 1,
		trigram("rap"): 1,
		trigram("rul"): 1,
		trigram("t_d"): 1,
		trigram("t_k"): 1,
		trigram("tra"): 1,
		trigram("ull"): 1,
		trigram("van"): 1,
		trigram("ap_"): 1,
	}

	counts := make(map[Trigram]int, 300)

	countTrigrams(test, counts)

	if len(counts) != len(expected) {
		t.Error("wrong number of trigrams")
	}

	for k, v := range counts {
		if expected[k] != v {
			t.Error("wrong count for", k)
		}
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
			if dist != testcase.Distance {
				t.Error("got", dist, "but expected", testcase.Distance)
			}
		})
	}
}

func TestTrain(t *testing.T) {
	text := "How much would a woodchunk chunk if a woodchunk could chunk wood?"
	got := Train(text)

	expected := []Trigram{
		trigram("_wo"),
		trigram("chu"),
		trigram("hun"),
		trigram("nk_"),
		trigram("unk"),
		trigram("ood"),
		trigram("woo"),
		trigram("_a_"),
		trigram("_ch"),
		trigram("a_w"),
		trigram("dch"),
		trigram("k_c"),
		trigram("ld_"),
		trigram("odc"),
		trigram("oul"),
		trigram("uld"),
		trigram("__h"),
		trigram("_co"),
		trigram("_ho"),
		trigram("_if"),
		trigram("_mu"),
		trigram("ch_"),
		trigram("cou"),
		trigram("d_a"),
		trigram("d_c"),
		trigram("f_a"),
		trigram("h_w"),
		trigram("how"),
		trigram("if_"),
		trigram("k_i"),
		trigram("k_w"),
		trigram("muc"),
		trigram("od_"),
		trigram("ow_"),
		trigram("uch"),
		trigram("w_m"),
		trigram("wou"),
	}

	if !reflect.DeepEqual(got, expected) {
		t.Error()
	}
}

func TestTrigramTextMarshal(t *testing.T) {
	tr0 := Trigram{'a', 'b', 'c'}
	b, _ := tr0.MarshalText()
	var tr1 Trigram
	_ = tr1.UnmarshalText(b)

	if tr0 != tr1 {
		t.Error()
	}
}
