# push-swap

Progetto in Go che ordina una lista di interi usando solo due stack (`a` e `b`)
e un set limitato di istruzioni. Il repository contiene due programmi:

- **`push-swap`** calcola e stampa, una per riga, le istruzioni che ordinano la lista.
- **`checker`** legge le istruzioni da stdin, le esegue sulla lista e stampa `OK` se `a` è ordinato e `b` è vuoto, altrimenti `KO`.

Usa solo la libreria standard di Go.

## Indice

- [Requisiti](#requisiti)
- [Compilazione](#compilazione)
- [Utilizzo](#utilizzo)
- [Istruzioni disponibili](#istruzioni-disponibili)
- [Gestione degli errori](#gestione-degli-errori)
- [Come funziona l'algoritmo](#come-funziona-lalgoritmo)
- [Struttura del progetto](#struttura-del-progetto)
- [Test](#test)
- [Modalità debug](#modalità-debug)
- [Limiti noti](#limiti-noti)

## Requisiti

- Go 1.21 o successivo.
- Nessuna dipendenza esterna.

## Compilazione

Dalla cartella che contiene `go.mod`:

```console
mkdir -p bin
go build -o bin/push-swap ./cmd/push-swap
go build -o bin/checker ./cmd/checker
```

Su Windows aggiungi `.exe` ai nomi dei binari (`bin/push-swap.exe`, `bin/checker.exe`).
Il percorso di output deve essere diverso dalle cartelle dei sorgenti, altrimenti Go scrive il binario dentro la cartella.

## Utilizzo

La lista si passa come un unico argomento tra virgolette (oppure come più argomenti). Il **primo numero è la cima dello stack `a`**.

```console
$ ./bin/push-swap "2 1 3 6 5 8"
ra
pb
...

$ ARG="4 67 3 87 23"
$ ./bin/push-swap "$ARG" | ./bin/checker "$ARG"
OK
$ ./bin/push-swap "$ARG" | wc -l      # numero di istruzioni
```

Il checker può anche essere usato a mano:

```console
$ echo -e "rra\npb\nsa\nrra\npa" | ./bin/checker "3 2 1 0"
OK
```

Se non ci sono argomenti, entrambi i programmi terminano senza stampare nulla. Se la lista è già ordinata, `push-swap` non stampa nessuna istruzione.

## Istruzioni disponibili

| Istruzione | Effetto |
|---|---|
| `sa` / `sb` / `ss` | scambia i primi due elementi di `a` / `b` / entrambi |
| `pa` | sposta il primo elemento di `b` in cima ad `a` |
| `pb` | sposta il primo elemento di `a` in cima a `b` |
| `ra` / `rb` / `rr` | ruota verso l'alto: il primo elemento diventa l'ultimo |
| `rra` / `rrb` / `rrr` | ruota verso il basso: l'ultimo elemento diventa il primo |

Ogni istruzione deve stare su una riga e non sono ammessi spazi o altri caratteri.

## Gestione degli errori

Su stderr viene stampato `Error` (seguito da un a capo) e il programma esce con codice 1 quando:

- un argomento non è un intero valido;
- ci sono numeri duplicati;
- il checker riceve una istruzione inesistente o scritta male.

Il checker ignora le righe vuote e controlla tutte le istruzioni prima di eseguire qualsiasi mossa.

## Come funziona l'algoritmo

`solver.Solve` sceglie la strategia in base alla lunghezza della lista:

| Elementi | Strategia |
|---|---|
| già ordinata | nessuna istruzione |
| 2 | un solo `sa` |
| 3 | `sortThree`: sequenza minima per ciascuna permutazione |
| 4 – 5 | `selection`: porta i minimi in `b` col percorso più breve, ordina i 3 rimasti, poi `pa` |
| 6 o più | `radix`: radix sort binario sui ranghi |

**Radix sort sui ranghi.** Ogni valore viene sostituito dalla posizione che avrebbe nella lista ordinata (`2 1 3 6 5 8` → `1 0 2 4 3 5`). Così i numeri negativi o molto grandi non creano problemi e bastano pochi bit. Per ogni bit, dal meno significativo, un elemento con bit 0 va in `b` (`pb`), uno con bit 1 viene ruotato (`ra`); a fine passata tutto torna in `a` con `pa`. Dopo l'ultimo bit la lista è ordinata e i valori originali vengono rimessi al posto dei ranghi.

## Struttura del progetto

```
push-swap/
├── go.mod
├── README.md
├── cmd/
│   ├── checker/main.go        programma checker
│   └── push-swap/main.go      programma push-swap
└── internal/
    ├── parser/                lettura e validazione degli argomenti
    │   ├── parser.go
    │   └── parser_test.go
    ├── stack/                 stack a/b e istruzioni
    │   ├── stack.go           stato, Do, Sorted
    │   ├── ops.go             operazioni elementari (swap, rotate, push)
    │   └── stack_test.go
    └── solver/                algoritmi di ordinamento
        ├── solver.go          scelta della strategia
        ├── small.go           da 2 a 5 elementi
        ├── radix.go           radix sort per 6 o più elementi
        └── solver_test.go
```

`push-swap` e `checker` condividono gli stessi pacchetti, quindi le regole delle istruzioni sono identiche nei due programmi.

## Test

```console
go test ./...
```

I test controllano, tra l'altro:

- il rifiuto di valori non interi e duplicati;
- l'esempio della traccia (`2 1 3 6 8 5` con `sa pb pb pb rb rrr pa pa pa`);
- che `Solve` ordini correttamente molte permutazioni casuali, da 0 a 120 elementi;
- che una lista già ordinata produca 0 istruzioni.

## Modalità debug

Con la variabile d'ambiente `PS_DEBUG` impostata, `push-swap` stampa su stderr lo stato degli stack dopo ogni istruzione e, alla fine, `OK` o `KO`. L'output su stdout resta invariato.

```console
$ PS_DEBUG=1 ./bin/push-swap "2 1 3 6 5 8"
```

È utile dove il checker non può essere eseguito (per esempio con policy di sicurezza che bloccano i binari non firmati).

## Limiti noti

- Il numero di istruzioni **non è il minimo possibile**. La radix sort è semplice e affidabile, ma su liste piccole produce più mosse di approcci basati su inserzione con calcolo del costo. Per esempio `2 1 3 6 5 8` richiede 29 istruzioni invece delle 7 dell'esempio della traccia.
- Possibili miglioramenti: inserzione con scelta del costo minimo (usando `rr` e `rrr`) per liste piccole e approccio a blocchi (chunk) per liste di 100 o 500 elementi.
