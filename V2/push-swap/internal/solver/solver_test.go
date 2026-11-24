package solver

import (
	"math/rand"
	"testing"

	"push-swap/internal/stack"
)

// Applica le istruzioni prodotte da Solve a molte permutazioni casuali.
func TestSolveSorts(t *testing.T) {
	for n := 0; n <= 120; n++ {
		for trial := 0; trial < 20; trial++ {
			nums := rand.Perm(n * 3)[:n]
			st := stack.New(nums)
			for _, op := range Solve(nums) {
				if err := st.Do(op); err != nil {
					t.Fatal(err)
				}
			}
			if !st.Sorted() {
				t.Fatalf("non ordinato per %v: %v %v", nums, st.A, st.B)
			}
		}
	}
}

func TestAlreadySorted(t *testing.T) {
	if ops := Solve([]int{1, 2, 3, 4}); len(ops) != 0 {
		t.Errorf("attese 0 istruzioni, trovate %d", len(ops))
	}
}
