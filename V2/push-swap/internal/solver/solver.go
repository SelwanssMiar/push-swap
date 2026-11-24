// Il pacchetto solver calcola le istruzioni che ordinano lo stack a.
package solver

import "push-swap/internal/stack"

// solver registra ogni istruzione che applica.
type solver struct {
	st  *stack.State
	ops []string
}

// do applica le istruzioni allo stato e le aggiunge alla lista.
func (v *solver) do(ops ...string) {
	for _, op := range ops {
		_ = v.st.Do(op)
		v.ops = append(v.ops, op)
	}
}

// Solve restituisce la lista di istruzioni che ordina nums.
func Solve(nums []int) []string {
	v := &solver{st: stack.New(nums)}
	n := len(nums)
	if stack.IsAscending(v.st.A) {
		return nil
	}
	switch {
	case n == 2:
		v.do("sa")
	case n == 3:
		v.sortThree()
	case n <= 5:
		v.selection()
	default:
		v.radix()
	}
	return v.ops
}
