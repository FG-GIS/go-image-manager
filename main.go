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
	// Definizione delle variabili che conterranno i parametri obbligatori ed opzionali.
	// bool flags => -flag=true/false
	// ext => le estensioni supportate
	var ext stringSlice = internal.Extensions
	// exclude => cartelle da escludere
	var exclude stringSlice
	// verbose => flag per attivare l'output testuale per le operazioni in corso
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	// extension-list => flag per la stampa dei formati supportati
	flist := flag.Bool("extension-list", false, "Prints out supported img extensions")
	path := flag.String("path", ".", "the path to scan for duplicate images")
	flag.Var(&exclude, "exclude", "Comma separated list of folders to exclude from the analisys (ex.: whatsapp,video)")
	flag.Var(&ext, "ext", "Comma separated list of extensions to evaluate (ex.: .jpg,.png)")

	flag.Parse()

	if *flist {
		fmt.Println("Available extensions are:\n", ext)
		os.Exit(0)
	}

	args := flag.Args()
	// in teoria utilizzare arg diversi per funzioni diverse: scan, link, delete
	if len(args) < 1 {
		fmt.Println("Implement guided procedure.")
		os.Exit(0)
	}

	if *verbose {
		fmt.Printf("Path to check: %s\n", *path)
	}

	if len(args) > 0 {
		if *verbose {
			fmt.Printf("Leftover args: %v\n", args)
		}

		fileList := internal.FileScanner(*path, exclude, *verbose)
		if fileList != nil {
			// fmt.Println("Lista di file trovati:\n", fileList)
			filteredList := internal.FilterImages(fileList, internal.GetExtensionsMap(ext))
			fmt.Println("Lista di file filtrati:\n", filteredList)
		} else {
			fmt.Println("No file found, or scan error")
		}
	}
}
