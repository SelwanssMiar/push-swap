package stack

// Operazioni elementari su una singola slice (indice 0 = cima dello stack).

func swap(s []int) {
	if len(s) > 1 {
		s[0], s[1] = s[1], s[0]
	}
}

func rotate(s []int) {
	if len(s) > 1 {
		first := s[0]
		copy(s, s[1:])
		s[len(s)-1] = first
	}
}

func reverseRotate(s []int) {
	if len(s) > 1 {
		last := s[len(s)-1]
		copy(s[1:], s[:len(s)-1])
		s[0] = last
	}
}

func push(from, to *[]int) {
	if len(*from) == 0 {
		return
	}
	*to = append([]int{(*from)[0]}, *to...)
	*from = (*from)[1:]
}
