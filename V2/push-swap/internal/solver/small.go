package solver

import "push-swap/internal/stack"

// sortThree ordina tre elementi con al massimo due istruzioni.
func (v *solver) sortThree() {
	a := v.st.A
	x, y, z := a[0], a[1], a[2]
	switch {
	case x > y && y < z && x < z:
		v.do("sa")
	case x > y && y > z:
		v.do("sa", "rra")
	case x > y && y < z && x > z:
		v.do("ra")
	case x < y && y > z && x < z:
		v.do("sa", "ra")
	case x < y && y > z && x > z:
		v.do("rra")
	}
}

// selection gestisce 4 e 5 elementi: mette i più piccoli in b.
func (v *solver) selection() {
	for len(v.st.A) > 3 {
		a := v.st.A
		idx := 0
		for i, x := range a {
			if x < a[idx] {
				idx = i
			}
		}
		// Si sceglie il percorso più corto verso la cima.
		if idx <= len(a)/2 {
			for ; idx > 0; idx-- {
				v.do("ra")
			}
		} else {
			for ; idx < len(a); idx++ {
				v.do("rra")
			}
		}
		v.do("pb")
	}
	if !stack.IsAscending(v.st.A) {
		v.sortThree()
	}
	for len(v.st.B) > 0 {
		v.do("pa")
	}
}
