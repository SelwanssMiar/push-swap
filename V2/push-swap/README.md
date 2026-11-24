# push-swap

Ordina una lista di interi con due stack (`a` e `b`) e un set di istruzioni.

## Struttura

```
push-swap/
├── go.mod
├── README.md
├── cmd/
│   ├── checker/main.go        programma checker (OK / KO)
│   └── push-swap/main.go      programma push-swap (stampa le istruzioni)
└── internal/
    ├── parser/                lettura e validazione degli argomenti
    │   ├── parser.go
    │   └── parser_test.go
    ├── stack/                 stack a/b e istruzioni
    │   ├── stack.go
    │   ├── ops.go
    │   └── stack_test.go
    └── solver/                algoritmi di ordinamento
        ├── solver.go          scelta della strategia
        ├── small.go           2-5 elementi
        ├── radix.go           radix sort per 6 o più elementi
        └── solver_test.go
```

## Compilazione e uso

```console
mkdir -p bin
go build -o bin/push-swap ./cmd/push-swap
go build -o bin/checker ./cmd/checker

ARG="4 67 3 87 23"
./bin/push-swap "$ARG" | ./bin/checker "$ARG"
./bin/push-swap "$ARG" | wc -l
```

Su Windows aggiungi `.exe` ai nomi dei binari. Per vedere gli stack passo
per passo: `PS_DEBUG=1 ./bin/push-swap "2 1 3 6 5 8"`.

## Test

```console
go test ./...
```
