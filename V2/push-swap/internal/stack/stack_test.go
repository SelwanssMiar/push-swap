package stack

import "testing"

// Esempio della traccia: 2 1 3 6 8 5 con le istruzioni indicate.
func TestSubjectExample(t *testing.T) {
	st := New([]int{2, 1, 3, 6, 8, 5})
	for _, op := range []string{"sa", "pb", "pb", "pb", "rb", "rrr", "pa", "pa", "pa"} {
		if err := st.Do(op); err != nil {
			t.Fatal(err)
		}
	}
	if !st.Sorted() {
		t.Errorf("stato non ordinato: %v %v", st.A, st.B)
	}
}

func TestUnknownInstruction(t *testing.T) {
	if err := New([]int{1}).Do("xx"); err == nil {
		t.Error("atteso errore per istruzione sconosciuta")
	}
}
