package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
)

func FileScanner(root string, exclusions []string) []string {
	var resultList []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %q: %v\n", path, err)
			return nil
		}
		if d.IsDir() && slices.Contains(exclusions, d.Name()) {
			fmt.Printf("Skipping directory %q due to exclusion\n", path)
			return filepath.SkipDir
		}

		if !d.IsDir() {
			resultList = append(resultList, path)
		}

		// qui si può inserire la logica per controllare id metadati dei file
		// bisogna in ogni caso controllare l'estensione e creare una lista
		// si potrà ampliare dopo inserendola nella struct
		// volendo a questo stato si possono raccogliere tutti i file senza filtrare per estensione

		return nil
	})
	if err != nil {
		fmt.Println("Error during FileSearch")
		return nil
	}

	return resultList
}
