package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/parser"
	"push-swap/internal/solver"
	"push-swap/internal/stack"
)

func main() {
	// Senza argomenti il programma non stampa nulla.
	if len(os.Args) < 2 {
		return
	}
	nums, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	ops := solver.Solve(nums)

	// Una istruzione per riga su stdout.
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, op := range ops {
		fmt.Fprintln(out, op)
	}

	// Modalità debug: con PS_DEBUG=1 mostra gli stack passo per passo su stderr.
	if os.Getenv("PS_DEBUG") != "" {
		st := stack.New(nums)
		fmt.Fprintf(os.Stderr, "%-6s a=%v b=%v\n", "start", st.A, st.B)
		for _, op := range ops {
			_ = st.Do(op)
			fmt.Fprintf(os.Stderr, "%-6s a=%v b=%v\n", op, st.A, st.B)
		}
		if st.Sorted() {
			fmt.Fprintln(os.Stderr, "OK")
		} else {
			fmt.Fprintln(os.Stderr, "KO")
		}
	}
}
