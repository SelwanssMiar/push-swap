// Il pacchetto parser converte gli argomenti della riga di comando in interi.
package parser

import (
	"errors"
	"strconv"
	"strings"
)

// ErrInvalid viene restituito per argomenti non interi o duplicati.
var ErrInvalid = errors.New("invalid input")

// Parse trasforma gli argomenti (ognuno può essere una lista separata da
// spazi) in una lista di interi senza duplicati.
func Parse(args []string) ([]int, error) {
	var nums []int
	seen := map[int]bool{}
	for _, arg := range args {
		for _, f := range strings.Fields(arg) {
			n, err := strconv.Atoi(f)
			if err != nil || seen[n] {
				return nil, ErrInvalid
			}
			seen[n] = true
			nums = append(nums, n)
		}
	}
	return nums, nil
}
