// Il pacchetto stack implementa i due stack e il set di istruzioni.
package stack

import "fmt"

// State contiene entrambi gli stack. L'indice 0 è la cima dello stack.
type State struct {
	A, B []int
}

// New crea uno stato in cui lo stack a contiene nums.
func New(nums []int) *State {
	return &State{A: append([]int(nil), nums...)}
}

// Do esegue una singola istruzione.
func (s *State) Do(op string) error {
	switch op {
	case "pa":
		push(&s.B, &s.A)
	case "pb":
		push(&s.A, &s.B)
	case "sa":
		swap(s.A)
	case "sb":
		swap(s.B)
	case "ss":
		swap(s.A)
		swap(s.B)
	case "ra":
		rotate(s.A)
	case "rb":
		rotate(s.B)
	case "rr":
		rotate(s.A)
		rotate(s.B)
	case "rra":
		reverseRotate(s.A)
	case "rrb":
		reverseRotate(s.B)
	case "rrr":
		reverseRotate(s.A)
		reverseRotate(s.B)
	default:
		return fmt.Errorf("unknown instruction %q", op)
	}
	return nil
}

// Sorted indica se a è in ordine crescente e b è vuoto.
func (s *State) Sorted() bool {
	return len(s.B) == 0 && IsAscending(s.A)
}

// IsAscending indica se la slice è in ordine crescente.
func IsAscending(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
