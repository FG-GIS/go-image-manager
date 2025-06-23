package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FG-GIS/go-image-manager/internal"
)

func main() {
	// 1. Definizione dei flag
	// flag.Bool(nome, valore_default, "descrizione")
	verbose := flag.Bool("verbose", false, "Abilita l'output dettagliato")

	flist := flag.Bool("flist", false, "Stampa i formati accettati")

	path := flag.String("path", ".", "tha path to scan for duplicate images")
	// Altri esempi di flag:
	// port := flag.Int("port", 8080, "Porta del server")
	// host := flag.String("host", "localhost", "Host del server")

	// 2. Parsing degli argomenti
	// È fondamentale chiamarlo dopo aver definito tutti i flag.
	flag.Parse()

	// 3. Accesso agli argomenti posizionali
	// flag.Args() restituisce una slice di stringhe con gli argomenti non-flag.
	args := flag.Args()

	// Check full procedure case
	if len(args) < 1 {
		fmt.Println("Implementare procedura guidata")
		os.Exit(0)
	}

	if *flist {
		fmt.Println("Available extensions are:\n", internal.Extensions)
		os.Exit(0)
	}

	// Utilizzo del flag
	if *verbose {
		fmt.Println("Modalità verbose: ON")
	} else {
		fmt.Println("Modalità verbose: OFF")
	}

	// Utilizzo degli altri argomenti
	if len(args) > 1 {
		fmt.Printf("Argomenti rimanenti: %v\n", args)
		// Nel tuo caso, qui avresti il path "C:\programs"
		fmt.Printf("Path da verificare: %s\n", *path)
	} else {
		fmt.Println("Nessun path da verificare fornito.")
	}
}
