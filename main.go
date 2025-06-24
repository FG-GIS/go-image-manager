package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/FG-GIS/go-image-manager/internal"
)

func main() {
	verbose := flag.Bool("verbose", false, "Abilita l'output dettagliato")

	flist := flag.Bool("flist", false, "Stampa i formati accettati")

	path := flag.String("path", ".", "tha path to scan for duplicate images")
	exclude := flag.String("exclude", "", "name folders you want to exclude")

	flag.Parse()

	if *flist {
		fmt.Println("Available extensions are:\n", internal.Extensions)
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Implementare procedura guidata")
		os.Exit(0)
	}

	if *verbose {
		fmt.Printf("Path da verificare: %s\n", *path)
	}

	if len(args) > 0 {
		if *verbose {
			fmt.Printf("Argomenti rimanenti: %v\n", args)
		}
		exSlice := []string{}
		if *exclude != "" {
			exSlice = strings.Split(strings.ToLower(*exclude), ",")
		}

		fileList := internal.FileScanner(*path, exSlice)
		if fileList != nil {
			fmt.Println(fileList)
		} else {
			fmt.Println("Nessun file trovato o errore durante la scansione.")
		}
	} else {
		fmt.Println("Nessun path da verificare fornito.")
	}
}
