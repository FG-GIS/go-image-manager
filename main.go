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
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	flist := flag.Bool("flist", false, "Prints out supported img extensions")

	path := flag.String("path", ".", "the path to scan for duplicate images")
	// exclude := flag.String("exclude", "", "name folders you want to exclude")
	var exclude stringSlice
	var ext stringSlice = internal.Extensions

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

		fileList := internal.FileScanner(*path, exclude, *verbose)
		if fileList != nil {
			// fmt.Println("Lista di file trovati:\n", fileList)
			filteredList := internal.FilterImages(fileList, internal.GetExtensionsMap(ext))
			fmt.Println("Lista di file filtrati:\n", filteredList)
		} else {
			fmt.Println("Nessun file trovato o errore durante la scansione.")
		}
	}
}
