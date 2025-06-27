package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/FG-GIS/go-image-manager/internal"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(v string) error {
	*s = strings.Split(v, ",")
	return nil
}

func main() {
	verbose := flag.Bool("verbose", false, "Abilita l'output dettagliato")

	flist := flag.Bool("flist", false, "Stampa le estensioni accettate")

	path := flag.String("path", ".", "tha path to scan for duplicate images")
	// exclude := flag.String("exclude", "", "name folders you want to exclude")
	var exclude stringSlice
	var ext stringSlice = internal.Extensions

	flag.Var(&exclude, "exclude", "Nome di cartelle da escludere nella scansione, separate da una virgola (ex.: whatsapp,video)")
	flag.Var(&ext, "ext", "Estensioni immagini da verificare, separate da una virgola (ex.: .jpg,.png)")

	flag.Parse()

	if *flist {
		fmt.Println("Available extensions are:\n", ext)
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

		fileList := internal.FileScanner(*path, exclude)
		if fileList != nil {
			fmt.Println("Lista di file trovati:\n", fileList)
			filteredList := internal.FilterImages(fileList, internal.GetExtensionsMap(ext))
			fmt.Println("Lista di file filtrati:\n", filteredList)
		} else {
			fmt.Println("Nessun file trovato o errore durante la scansione.")
		}
	} else {
		fmt.Println("Nessun path da verificare fornito.")
	}
}
