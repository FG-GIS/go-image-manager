package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func FileScanner(root string, exclusions []string) []string {
	var resultList []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %q: %v\n", path, err)
			return err
		}
		name := strings.ToLower(d.Name())
		if d.IsDir() && slices.Contains(exclusions, name) {
			fmt.Printf("Skipping directory %q due to exclusion\n", path)
			return filepath.SkipDir
		}

		if !d.IsDir() {
			resultList = append(resultList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error during FileSearch")
		return nil
	}

	return resultList
}

func FilterImages(imgSlice []string, extMap map[string]bool) []FileData {
	result := []FileData{}
	for _, v := range imgSlice {
		ext := strings.ToLower(filepath.Ext(v))

		if _, exists := extMap[ext]; exists {
			info, err := os.Stat(v)
			var size int64
			if err == nil {
				size = info.Size()
			} else {
				size = 0
			}
			result = append(result, FileData{
				Path: v,
				Size: size,
			})
		}
	}
	return result
}
