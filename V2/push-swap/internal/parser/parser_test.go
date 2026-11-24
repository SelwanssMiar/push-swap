package parser

import "testing"

func TestParse(t *testing.T) {
	if _, err := Parse([]string{"0 one 2 3"}); err == nil {
		t.Error("atteso errore per un valore non intero")
	}
	if _, err := Parse([]string{"1 2 1"}); err == nil {
		t.Error("atteso errore per un duplicato")
	}
	got, err := Parse([]string{"2 1", "3"})
	if err != nil || len(got) != 3 {
		t.Errorf("risultato inatteso: %v %v", got, err)
	}
}
