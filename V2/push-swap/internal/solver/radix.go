package solver

import (
	"math/bits"
	"sort"

	"push-swap/internal/stack"
)

// radix è un radix sort LSD non comparativo, applicato ai ranghi dei valori.
func (v *solver) radix() {
	n := len(v.st.A)
	sorted := append([]int(nil), v.st.A...)
	sort.Ints(sorted)
	rank := make(map[int]int, n)
	for i, x := range sorted {
		rank[x] = i
	}
	// Si lavora sui ranghi; i valori originali vengono rimessi alla fine.
	for i, x := range v.st.A {
		v.st.A[i] = rank[x]
	}
	for bit := 0; bit < bits.Len(uint(n-1)); bit++ {
		if stack.IsAscending(v.st.A) {
			break
		}
		for i := 0; i < n; i++ {
			if (v.st.A[0]>>bit)&1 == 0 {
				v.do("pb")
			} else {
				v.do("ra")
			}
		}
		for len(v.st.B) > 0 {
			v.do("pa")
		}
	}
	for i, r := range v.st.A {
		v.st.A[i] = sorted[r]
	}
}
