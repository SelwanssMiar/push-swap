package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func fail() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}

func main() {
	// Senza argomenti il programma non stampa nulla.
	if len(os.Args) < 2 {
		return
	}
	nums, err := parser.Parse(os.Args[1:])
	if err != nil {
		fail()
	}

	// Prima si leggono tutte le istruzioni, così una istruzione errata
	// viene segnalata come Error prima di eseguire qualsiasi mossa.
	var ops []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		ops = append(ops, line)
	}

	st := stack.New(nums)
	for _, op := range ops {
		if err := st.Do(op); err != nil {
			fail()
		}
	}

	if st.Sorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
